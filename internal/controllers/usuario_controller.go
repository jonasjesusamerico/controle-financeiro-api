package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joninhasamerico/controle-financeiro-api/internal/controllers/rest_err"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
	"github.com/joninhasamerico/controle-financeiro-api/internal/repository"
	"github.com/joninhasamerico/controle-financeiro-api/internal/services"
	"gorm.io/gorm"
)

type ResponseError struct {
	Message string `json:"message"`
}

// UsuarioController  represent the httpHandler for article
type UsuarioController struct {
	super   BaseController
	service services.IUsuarioService
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

// FetchUsuario will fetch the article superd on given params
func (a *UsuarioController) FetchUsuario(c *gin.Context) {
	ctx := a.super.Ctx(c)

	var usuarios []model.Usuario

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
	ctx := a.super.Ctx(c)

	usuario := model.Usuario{}

	err = a.service.GetByID(ctx, &usuario, id)
	if err != nil {
		errRest := rest_err.NewNotFoundError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func (a *UsuarioController) Save(c *gin.Context) {
	var usuario model.Usuario
	err := c.ShouldBindJSON(&usuario)
	if err != nil {
		errRest := rest_err.NewUnprocessableEntityError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	err = a.service.Save(context.TODO(), &usuario)
	if err != nil {
		errRest := rest_err.NewInternalServerError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusCreated, usuario.GetUsuarioRetorno())
}

func (a *UsuarioController) Delete(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errRest := rest_err.NewNotFoundError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	id := int64(idP)
	ctx := a.super.Ctx(c)

	err = a.service.Delete(ctx, id)
	if err != nil {
		errRest := rest_err.NewInternalServerError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
