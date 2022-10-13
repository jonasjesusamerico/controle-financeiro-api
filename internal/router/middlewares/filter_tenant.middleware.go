package middlewares

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
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

		c.Set("tenantId", int64(tenantId))
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
