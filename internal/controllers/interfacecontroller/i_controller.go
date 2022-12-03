package interfacecontroller

import (
	"github.com/gin-gonic/gin"
)

type IController interface {
	NameGroupRoute() string

	FindAll(c *gin.Context)

	FindById(c *gin.Context)

	Create(c *gin.Context)

	Update(c *gin.Context)

	Delete(c *gin.Context)
}
