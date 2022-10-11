package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joninhasamerico/controle-financeiro-api/internal/controllers/resposta"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
	"github.com/joninhasamerico/controle-financeiro-api/internal/repository"
)

type UsuarioController struct {
	Repo repository.IRepository
}

func (uc UsuarioController) NameGroupRoute() string {
	return "/usuarios"
}

func (uc UsuarioController) FindAll(c *gin.Context) {
	var usuarios []model.Usuario

	err := uc.Repo.FindAll(&usuarios, "")
	if err != nil {
		resposta.Erro(c, http.StatusInternalServerError, err)
		return
	}

	if len(usuarios) == 0 {
		resposta.JSON(c, http.StatusNoContent, usuarios)
		return
	}

	resposta.JSON(c, http.StatusOK, usuarios)
}

func (uc UsuarioController) FindById(c *gin.Context) {
	var usuario model.Usuario
	id, err := strconv.ParseUint(c.Params.ByName("id"), 10, 64)
	if err != nil {
		resposta.Erro(c, http.StatusInternalServerError, err)
		return
	}

	err = uc.Repo.FindById(&usuario, id)
	if err != nil {
		resposta.Erro(c, http.StatusInternalServerError, err)
		return
	}

	if usuario.ID == 0 {
		resposta.JSON(c, http.StatusNotFound, errors.New("usuário não encontrado"))
		return
	}

	resposta.JSON(c, http.StatusOK, usuario)
}

func (uc UsuarioController) Create(c *gin.Context) {
	var usuario model.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		resposta.Erro(c, http.StatusBadRequest, err)
		return
	}

	usuarioSalvoNoBanco := model.Usuario{}
	erro := uc.Repo.FindFirst(&usuarioSalvoNoBanco, "email = ?", usuario.Email)
	if erro != nil {
		resposta.Erro(c, http.StatusInternalServerError, erro)
		return
	}

	if usuarioSalvoNoBanco.ID != 0 {
		resposta.Erro(c, http.StatusConflict, errors.New("email informado já está cadastrado"))
		return
	}

	_, err := uc.Repo.Save(&usuario)
	if err != nil {
		resposta.Erro(c, http.StatusInternalServerError, err)
		return
	}

	resposta.JSON(c, http.StatusOK, usuario.GetUsuarioRetorno())
}

func (uc UsuarioController) Update(c *gin.Context) {
	var usuario model.Usuario
	id := c.Params.ByName("id")

	err := uc.Repo.FindById(&usuario, id)
	if err != nil {
		resposta.Erro(c, http.StatusInternalServerError, err)
		return
	}

	if err := c.ShouldBindJSON(&usuario); err != nil {
		resposta.Erro(c, http.StatusInternalServerError, err)
		return
	}

	resposta.JSON(c, http.StatusOK, usuario.GetUsuarioRetorno())
}

func (uc UsuarioController) Delete(c *gin.Context) {
	resposta.BadRequest(c, "Entre em contato com o suporte para solicitar exclusão do usuário desejado")
}

func (uc UsuarioController) RotaCustomizada(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Oi, eu sou uma rota customizada!"})
}
