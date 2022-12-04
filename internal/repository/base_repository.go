package repository

import (
	"context"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model/enum"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model/interfacemodel"
	"github.com/joninhasamerico/controle-financeiro-api/internal/repository/interfacerepository"
	"gorm.io/gorm"
)

type BaseRepository struct {
	dbCtx *gorm.DB
}

func NewBaseRepository(dbCtx *gorm.DB) interfacerepository.IBaseRepository {
	return BaseRepository{
		dbCtx: dbCtx,
	}
}

func (b BaseRepository) TenantCtx(ctx context.Context) *gorm.DB {
	return b.dbCtx.Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where("tenant_id = ?", b.TenantID(ctx))
	})
}

func (b BaseRepository) TenantID(ctx context.Context) int64 {
	return ctx.Value(enum.TENANT_ID).(int64)
}

func (b BaseRepository) Repo() *gorm.DB {
	return b.dbCtx
}

func (b BaseRepository) SetTenant(ctx context.Context, model interfacemodel.IModel) *gorm.DB {
	model.SetTenant(b.TenantID(ctx))
	return b.dbCtx
}
