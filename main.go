package main

import (
	"fmt"

	"github.com/joninhasamerico/controle-financeiro-api/configs"
	"github.com/joninhasamerico/controle-financeiro-api/internal/router"
)

func init() {
	configs.Carregar()
}

func main() {
	fmt.Println("Olá mundo véio!")

	router.CriarRotas()
}
