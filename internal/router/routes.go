package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joninhasamerico/controle-financeiro-api/internal/controllers"
	"github.com/joninhasamerico/controle-financeiro-api/internal/router/middlewares"
	"gorm.io/gorm"
)

func NewSetupRouter(engine *gin.Engine, dbCtx *gorm.DB) {
	timeoutCtx := time.Duration(2) * time.Second

	/* Middlewares */
	engine.Use(middlewares.CORSMiddleware())

	/* Groups */
	MAIN_GROUP := engine.Group("/")
	API_GROUP := MAIN_GROUP.Group("/api", middlewares.MiddleAuth(), middlewares.FilterTenantMiddleware())
	V1_GROUP := API_GROUP.Group("/v1")

	/* Login */
	controllers.NewLoginController(MAIN_GROUP, dbCtx)

	/* Usuario */
	controllers.NewUsuarioController(MAIN_GROUP, V1_GROUP, dbCtx, timeoutCtx)

	/* Lancamento */
	controllers.NewLancamentoController(V1_GROUP, dbCtx, timeoutCtx)
}
