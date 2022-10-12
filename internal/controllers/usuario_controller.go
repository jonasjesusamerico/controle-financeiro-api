package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
	"github.com/joninhasamerico/controle-financeiro-api/internal/services"
	"github.com/sirupsen/logrus"
)

type ResponseError struct {
	Message string `json:"message"`
}

// UsuarioController  represent the httpHandler for article
type UsuarioController struct {
	super   BaseController
	service services.IService
}

// NewUsuarioController will initialize the usuarios/ resources endpoint
func NewUsuarioController(e *gin.Engine, us services.IService) {
	handler := &UsuarioController{
		service: us.(*services.UsuarioService),
	}

	grp1 := e.Group("/v1")
	{
		grp1.GET("/usuarios", handler.FetchUsuario)
		grp1.POST("/usuarios", handler.Save)
		grp1.GET("/usuarios/:id", handler.GetByID)
		grp1.DELETE("/usuarios/:id", handler.Delete)
	}
}

// FetchUsuario will fetch the article superd on given params
func (a *UsuarioController) FetchUsuario(c *gin.Context) {
	ctx := a.super.Ctx(c)

	listAr, err := a.service.FindAll(ctx)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, listAr)
}

// GetByID will get article by given id
func (a *UsuarioController) GetByID(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, model.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := a.super.Ctx(c)

	art, err := a.service.GetByID(ctx, id)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, art)
}

func isRequestValid(m *model.Usuario) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Store will store the article by given request body
func (a *UsuarioController) Save(c *gin.Context) {
	var article model.Usuario
	err := c.Bind(&article)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	var ok bool
	if ok, err = isRequestValid(&article); !ok {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx := a.super.Ctx(c)
	err = a.service.Save(ctx, &article)
	if err != nil {
		c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, article)
}

// Delete will delete article by given param
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
