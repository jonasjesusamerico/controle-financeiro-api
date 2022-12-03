package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model/enum"
)

type BaseController struct{}

func (b BaseController) Ctx(c *gin.Context) context.Context {
	tenantId := c.MustGet("tenantId")
	return context.WithValue(c.Request.Context(), enum.TENANT_ID, tenantId)
}
