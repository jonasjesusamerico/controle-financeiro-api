package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/joninhasamerico/controle-financeiro-api/internal/repository"
)

type IHandler interface {
	New(Repo repository.IRepository, Route *gin.RouterGroup) IHandler
	RotasAutenticadas() IHandler
	RotasNaoAutenticadas() IHandler
}

type Handler struct {
	Route *gin.Engine
}

func (h Handler) MakeHandlers() {

	main := h.Route.Group("/")
	api := main.Group("api")
	v1 := api.Group("v1")

	rotasMain := []IHandler{} // Rotas que são da raiz /

	rotasApi := []IHandler{} // Rotas especificas para /main/api

	rotasV1 := []IHandler{
		&UsuarioHandler{Repo: repository.UsuarioRepository{}, Route: v1},
	} // Rotas especificas da versão v1 main/api/v1

	criaRotas(rotasMain)
	criaRotas(rotasApi)
	criaRotas(rotasV1)

}

func criaRotas(rotas []IHandler) {
	if len(rotas) == 0 {
		return
	}

	for _, rota := range rotas {
		rota.RotasAutenticadas().RotasNaoAutenticadas()
	}
}
