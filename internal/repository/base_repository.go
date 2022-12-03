package repository

import (
	"context"

	"gorm.io/gorm"
)

type BaseRepository struct {
	dbCtx *gorm.DB
}

func (b BaseRepository) tenantCtx(ctx context.Context) *gorm.DB {
	return b.dbCtx.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("tenant_id = ?", b.TenantID(ctx))
	})
}

func (b BaseRepository) TenantID(ctx context.Context) int64 {
	return ctx.Value("tenantId").(int64)
}
