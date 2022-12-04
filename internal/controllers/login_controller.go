package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joninhasamerico/controle-financeiro-api/internal/controllers/rest_err"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model/usuario"
	"github.com/joninhasamerico/controle-financeiro-api/internal/repository"
	"github.com/joninhasamerico/controle-financeiro-api/internal/repository/interfacerepository"
	"github.com/joninhasamerico/controle-financeiro-api/pkg/auth"
	"gorm.io/gorm"
)

type LoginController struct {
	Repo interfacerepository.IUsuarioRepository
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
	user := usuario.NewUsuario()

	if err := c.ShouldBindJSON(&user); err != nil {
		errRest := rest_err.NewBadRequestError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	usuarioSalvoNoBanco := usuario.NewUsuario()
	err := lc.Repo.GetByEmail(user.Email, usuarioSalvoNoBanco)
	if err != nil {
		errRest := rest_err.NewNotFoundError("The email entered is not a valid email")
		c.JSON(errRest.Code, errRest)
		return
	}

	if err = auth.VerificarSenha(usuarioSalvoNoBanco.Senha, user.Senha); err != nil {
		errRest := rest_err.NewForbiddenError(err.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	token, erro := auth.CriarToken(usuarioSalvoNoBanco.ID)
	if erro != nil {
		errRest := rest_err.NewInternalServerError(erro.Error())
		c.JSON(errRest.Code, errRest)
		return
	}

	usuarioID := strconv.FormatUint(usuarioSalvoNoBanco.ID, 10)

	c.JSON(http.StatusOK, model.DadosAutenticacao{ID: usuarioID, Token: token})
}
