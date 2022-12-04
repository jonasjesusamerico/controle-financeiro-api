package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model/enum"
	"github.com/joninhasamerico/controle-financeiro-api/pkg/auth"
)

func FilterTenantMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		/* Check tenantId is valid here */
		tenantId, err := auth.ExtrairTenantID(c)
		if err != nil {
			c.AbortWithError(http.StatusNotFound, errors.New("TenantId invalid ! "))
			return
		}

		c.Set(string(enum.TENANT_ID), int64(tenantId))
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func MiddleAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if erro := auth.ValidarToken(ctx); erro != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
