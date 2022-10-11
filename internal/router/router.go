package router

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joninhasamerico/controle-financeiro-api/configs"
	"github.com/joninhasamerico/controle-financeiro-api/internal/router/handlers"
)

func CriarRotas() {

	r := gin.Default()
	r.SetTrustedProxies(nil)

	handlers.Handler{Route: r}.MakeHandlers()

	port := configs.Porta

	r.Run(":" + strconv.Itoa(port))
}
