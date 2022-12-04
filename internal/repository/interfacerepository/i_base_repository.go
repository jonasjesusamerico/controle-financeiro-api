package interfacerepository

import (
	"context"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model/interfacemodel"
	"gorm.io/gorm"
)

type IBaseRepository interface {
	TenantCtx(ctx context.Context) *gorm.DB
	TenantID(ctx context.Context) int64
	Repo() *gorm.DB
	SetTenant(ctx context.Context, model interfacemodel.IModel) *gorm.DB
}
