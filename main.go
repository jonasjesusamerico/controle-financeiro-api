package main

import (
	"fmt"

	"github.com/joninhasamerico/controle-financeiro-api/configs"
)

func init() {
	configs.Carregar()
}

func main() {
	fmt.Println("Olá mundo véio!")
}
