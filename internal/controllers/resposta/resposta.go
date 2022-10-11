package resposta

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func JSON(c *gin.Context, statusCode int, dados interface{}) {

	if dados == nil {
		return
	}

	c.JSON(statusCode, dados)
}

func Erro(c *gin.Context, statusCode int, erro error) {
	JSON(c, statusCode, struct {
		Erro string `json:"error"`
	}{
		Erro: erro.Error(),
	})

}

func BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{"message": message})
}

func NotFound(c *gin.Context, registro string) {
	c.JSON(http.StatusNotFound, gin.H{"message": registro + " não encontrado"})
}

func OkMessage(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{"message": message})
}
