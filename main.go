package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joninhasamerico/controle-financeiro-api/configs"
	"github.com/joninhasamerico/controle-financeiro-api/internal/controllers"
	"github.com/joninhasamerico/controle-financeiro-api/internal/database"
	"github.com/joninhasamerico/controle-financeiro-api/internal/router/middlewares"
	"gorm.io/gorm"
)

func init() {

}

func main() {
	fmt.Println("Iniciando a aplicação!")
	configs.Carregar()
	dbCtx, err := database.Conect()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	engine := gin.Default()
	setup(engine, dbCtx)

	log.Fatal(engine.Run(":" + strconv.Itoa(configs.Porta)))
}

func setup(engine *gin.Engine, dbCtx *gorm.DB) {
	timeoutCtx := time.Duration(2) * time.Second

	/* Middlewares */
	// engine.Use()

	/* Groups */
	MAIN_GROUP := engine.Group("/")
	API_GROUP := MAIN_GROUP.Group("/api", middlewares.MiddleAuth(), middlewares.FilterTenantMiddleware())
	V1_GROUP := API_GROUP.Group("/v1")

	/* Login */
	controllers.NewLoginController(MAIN_GROUP, dbCtx)

	/* Usuario */
	controllers.NewUsuarioController(MAIN_GROUP, V1_GROUP, dbCtx, timeoutCtx)

}
