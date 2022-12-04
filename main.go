package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joninhasamerico/controle-financeiro-api/configs"
	"github.com/joninhasamerico/controle-financeiro-api/internal/database"
	"github.com/joninhasamerico/controle-financeiro-api/internal/router"
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
	router.NewSetupRouter(engine, dbCtx)

	log.Fatal(engine.Run(":" + strconv.Itoa(configs.Porta)))
}
