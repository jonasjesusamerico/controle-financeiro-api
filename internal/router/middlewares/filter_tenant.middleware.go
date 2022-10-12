package middlewares

import (
	"github.com/gin-gonic/gin"
)

func FilterTenantMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		/* Check tenantId is valid here */
		// tenantId, err := strconv.ParseInt(c.Request.Header.Get("X-Tenant-ID"), 10, 64)
		// if err != nil {
		// 	c.AbortWithError(http.StatusNotFound, errors.New("TenantId invalid ! "))
		// 	return
		// }

		c.Set("tenantId", 1)
		c.Next()
	}
}
