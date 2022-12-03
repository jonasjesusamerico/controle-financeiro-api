package interfacerepository

import (
	"context"

	"gorm.io/gorm"
)

type IBaseRepository interface {
	TenantCtx(ctx context.Context) *gorm.DB
	TenantID(ctx context.Context) int64
	Repo() *gorm.DB
}
