package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joninhasamerico/controle-financeiro-api/internal/controllers/rest_err"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model/usuario"
	"github.com/joninhasamerico/controle-financeiro-api/internal/repository"
	"github.com/joninhasamerico/controle-financeiro-api/internal/services"
	"github.com/joninhasamerico/controle-financeiro-api/internal/services/interfaceservice"
	"gorm.io/gorm"
)

type UsuarioController struct {
	BaseController
	service interfaceservice.IUsuarioService
}

func NewUsuarioController(rotaMain *gin.RouterGroup, rotaV1 *gin.RouterGroup, dbCtx *gorm.DB, timeoutCtx time.Duration) {

	usuarioRepository := repository.NewUsuarioRepository(dbCtx)
	usuarioService := services.NewUsuarioService(usuarioRepository, timeoutCtx)

	handler := &UsuarioController{
		service: usuarioService,
	}

	{
		rotaMain.POST("/usuarios", handler.Save)
		rotaV1.GET("/usuarios", handler.FetchUsuario)
		rotaV1.GET("/usuarios/:id", handler.GetByID)
		rotaV1.DELETE("/usuarios/:id", handler.Delete)
	}
}

func (a *UsuarioController) FetchUsuario(c *gin.Context) {
	ctx := a.Ctx(c)

	usuarios := usuario.NewSliceUsuario()

	err := a.service.FindAll(ctx, &usuarios)
	if err != nil {
		errRest := rest_err.NewNotFoundError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, usuarios)
}

func (a *UsuarioController) GetByID(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errRest := rest_err.NewInternalServerError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	id := int64(idP)
	ctx := a.Ctx(c)

	usuario := *usuario.NewUsuario()

	err = a.service.GetByID(ctx, &usuario, id)
	if err != nil {
		errRest := rest_err.NewNotFoundError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func (a *UsuarioController) Save(c *gin.Context) {
	user := *usuario.NewUsuario()
	err := c.ShouldBindJSON(&user)
	if err != nil {
		errRest := rest_err.NewUnprocessableEntityError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	err = a.service.Save(context.TODO(), &user)
	if err != nil {
		errRest := rest_err.NewInternalServerError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusCreated, user.GetUsuarioRetorno())
}

func (a *UsuarioController) Delete(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errRest := rest_err.NewNotFoundError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	id := int64(idP)
	ctx := a.Ctx(c)

	err = a.service.Delete(ctx, id)
	if err != nil {
		errRest := rest_err.NewInternalServerError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
