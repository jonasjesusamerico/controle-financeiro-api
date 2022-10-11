package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/joninhasamerico/controle-financeiro-api/internal/controllers"
	"github.com/joninhasamerico/controle-financeiro-api/internal/repository"
)

type UsuarioHandler struct {
	Repo  repository.IRepository
	Route *gin.RouterGroup
}

func (u UsuarioHandler) New(repo repository.IRepository, rota *gin.RouterGroup) IHandler {
	u.Repo = repo
	u.Route = rota
	return &u
}

func (u UsuarioHandler) RotasAutenticadas() IHandler {

	controller := controllers.UsuarioController{Repo: u.Repo}

	route := u.Route.Group(controller.NameGroupRoute())
	{
		route.GET("/", controller.FindAll)
		route.GET("/:id", controller.FindById)
		route.PATCH("/:id", controller.Update)
		route.DELETE("/:id", controller.Delete)
	}

	return &u
}

func (u UsuarioHandler) RotasNaoAutenticadas() IHandler {
	controller := controllers.UsuarioController{Repo: u.Repo}

	route := u.Route.Group(controller.NameGroupRoute())
	{
		route.POST("/", controller.Create)
	}

	return &u
}
