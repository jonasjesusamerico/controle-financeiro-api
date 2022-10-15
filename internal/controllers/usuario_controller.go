package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
	"github.com/joninhasamerico/controle-financeiro-api/internal/repository"
	"github.com/joninhasamerico/controle-financeiro-api/internal/services"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ResponseError struct {
	Message string `json:"message"`
}

// UsuarioController  represent the httpHandler for article
type UsuarioController struct {
	super   BaseController
	service services.IService
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
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, usuarios)
}

func (a *UsuarioController) GetByID(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := a.super.Ctx(c)

	usuario := model.Usuario{}

	err = a.service.GetByID(ctx, &usuario, id)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func isRequestValid(m *model.Usuario) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *UsuarioController) Save(c *gin.Context) {
	var usuario model.Usuario
	err := c.Bind(&usuario)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	var ok bool
	if ok, err = isRequestValid(&usuario); !ok {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	service := a.service.(*services.UsuarioService)
	err = service.Save(nil, &usuario)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, usuario.GetUsuarioRetorno())
}

func (a *UsuarioController) Delete(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
		return
	}

	id := int64(idP)
	ctx := a.super.Ctx(c)

	err = a.service.Delete(ctx, id)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case model.ErrInternalServerError:
		return http.StatusInternalServerError
	case model.ErrNotFound:
		return http.StatusNotFound
	case model.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
