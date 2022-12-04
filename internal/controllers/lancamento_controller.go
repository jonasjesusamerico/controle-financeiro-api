package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joninhasamerico/controle-financeiro-api/internal/controllers/rest_err"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model/lancamento"
	"github.com/joninhasamerico/controle-financeiro-api/internal/repository"
	"github.com/joninhasamerico/controle-financeiro-api/internal/services"
	"github.com/joninhasamerico/controle-financeiro-api/internal/services/interfaceservice"
	"gorm.io/gorm"
)

type ResponseError struct {
	Message string `json:"message"`
}

type LancamentoController struct {
	BaseController
	service interfaceservice.ILancamentoService
}

func NewLancamentoController(rotaV1 *gin.RouterGroup, dbCtx *gorm.DB, timeoutCtx time.Duration) {

	lancamentoRepository := repository.NewLancamentoRepository(dbCtx)
	lancamentoService := services.NewLancamentoService(lancamentoRepository, timeoutCtx)

	handler := &LancamentoController{
		service: lancamentoService,
	}

	{
		rotaV1.POST("/lancamentos", handler.Save)
		rotaV1.GET("/lancamentos", handler.FetchLancamento)
		rotaV1.GET("/lancamentos/:id", handler.GetByID)
		rotaV1.DELETE("/lancamentos/:id", handler.Delete)
	}
}

func (a *LancamentoController) FetchLancamento(c *gin.Context) {
	ctx := a.Ctx(c)

	lancamentos := lancamento.NewSliceLancamento()

	err := a.service.FindAll(ctx, &lancamentos)
	if err != nil {
		errRest := rest_err.NewNotFoundError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, lancamentos)
}

func (a *LancamentoController) GetByID(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errRest := rest_err.NewInternalServerError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	l := *lancamento.NewLancamento()

	err = a.service.GetByID(a.Ctx(c), &l, int64(idP))
	if err != nil {
		errRest := rest_err.NewNotFoundError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, l)
}

func (a *LancamentoController) Save(c *gin.Context) {
	lancamentoDto := lancamento.NewLancamentoDto()

	err := c.ShouldBindJSON(&lancamentoDto)
	if err != nil {
		errRest := rest_err.NewUnprocessableEntityError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	l := *lancamento.LancamentoConverter(*lancamentoDto)

	err = a.service.Save(a.Ctx(c), &l)
	if err != nil {
		errRest := rest_err.NewInternalServerError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusCreated, l)
}

func (a *LancamentoController) Delete(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errRest := rest_err.NewNotFoundError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	id := int64(idP)
	err = a.service.Delete(a.Ctx(c), id)
	if err != nil {
		errRest := rest_err.NewNotFoundError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
