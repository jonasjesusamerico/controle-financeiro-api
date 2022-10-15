package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joninhasamerico/controle-financeiro-api/internal/controllers/resposta"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
	"github.com/joninhasamerico/controle-financeiro-api/internal/repository"
	"github.com/joninhasamerico/controle-financeiro-api/pkg/auth"
	"gorm.io/gorm"
)

type LoginController struct {
	Repo repository.IRepository
}

func NewLoginController(rotaMain *gin.RouterGroup, dbCtx *gorm.DB) {

	usuarioRepository := repository.NewUsuarioRepository(dbCtx)

	handler := &LoginController{
		Repo: usuarioRepository,
	}

	{
		rotaMain.POST("/login", handler.Login)
	}
}

func (lc LoginController) NameGroupRoute() string {
	return "/login"
}

func (lc LoginController) Login(c *gin.Context) {
	var usuario model.Usuario

	if err := c.ShouldBindJSON(&usuario); err != nil {
		resposta.Erro(c, http.StatusBadRequest, err)
		return
	}

	usuarioRepo := lc.Repo.(*repository.UsuarioRepository)
	usuarioSalvoNoBanco := model.Usuario{}
	err := usuarioRepo.GetByEmail(usuario.Email, &usuarioSalvoNoBanco)
	if err != nil {
		resposta.Erro(c, http.StatusInternalServerError, err)
		return
	}

	if err = auth.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); err != nil {
		resposta.Erro(c, http.StatusUnauthorized, err)
		return
	}

	token, erro := auth.CriarToken(usuarioSalvoNoBanco.ID)
	if erro != nil {
		resposta.Erro(c, http.StatusInternalServerError, erro)
		return
	}

	usuarioID := strconv.FormatUint(usuarioSalvoNoBanco.ID, 10)

	c.JSON(http.StatusOK, model.DadosAutenticacao{ID: usuarioID, Token: token})
}
