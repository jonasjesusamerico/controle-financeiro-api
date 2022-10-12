package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joninhasamerico/controle-financeiro-api/configs"
	"github.com/joninhasamerico/controle-financeiro-api/internal/controllers"
	"github.com/joninhasamerico/controle-financeiro-api/internal/database"
	"github.com/joninhasamerico/controle-financeiro-api/internal/repository"
	"github.com/joninhasamerico/controle-financeiro-api/internal/router/middlewares"
	"github.com/joninhasamerico/controle-financeiro-api/internal/services"
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

	log.Fatal(engine.Run())
}

func setup(engine *gin.Engine, dbCtx *gorm.DB) {
	timeoutCtx := time.Duration(200) * time.Second

	/* Middlewares */
	engine.Use(middlewares.FilterTenantMiddleware())

	/* Article */
	usuarioRepo := repository.NewUsuarioRepository(dbCtx)
	usuarioService := services.NewUsuarioService(usuarioRepo, timeoutCtx)
	controllers.NewUsuarioController(engine, usuarioService)
}
