package repository

import (
	"context"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model/interfacemodel"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model/lancamento"
	"github.com/joninhasamerico/controle-financeiro-api/internal/repository/interfacerepository"
	"gorm.io/gorm"
)

type LancamentoRepository struct {
	base interfacerepository.IBaseRepository
}

func NewLancamentoRepository(dbCtx *gorm.DB) interfacerepository.ILancamentoRepository {
	return &LancamentoRepository{
		base: NewBaseRepository(dbCtx),
	}
}

func (m *LancamentoRepository) FindAll(ctx context.Context, preloads []string, models interface{}) (err error) {
	stat := m.base.TenantCtx(ctx)

	teste := []string{"Conta", "Setor"}

	for _, value := range teste {
		stat = stat.Preload(value)
	}

	if err = stat.Statement.Find(models).Error; err != nil {
		return
	}

	return
}

func (m *LancamentoRepository) GetByID(ctx context.Context, models interfacemodel.IModel, id int64) (err error) {
	if err := m.base.TenantCtx(ctx).Where("id = ?", id).First(&models).Error; err != nil {
		return model.ErrNotFound
	}

	return nil
}

func (m *LancamentoRepository) Update(ctx context.Context, model interfacemodel.IModel) (err error) {
	if err = m.base.TenantCtx(ctx).Save(model).Error; err != nil {
		return err
	}
	return nil
}

func (m *LancamentoRepository) Save(ctx context.Context, model interfacemodel.IModel) (err error) {
	if err = m.base.SetTenant(ctx, model).Create(model).Error; err != nil {
		return err
	}
	return nil
}

func (m *LancamentoRepository) Delete(ctx context.Context, model interfacemodel.IModel, id int64) (err error) {
	if err = m.base.TenantCtx(ctx).Where("id = ?", id).Delete(&model).Error; err != nil {
		return err
	}
	return nil
}

func (m *LancamentoRepository) Exists(ctx context.Context, id int64) bool {
	lancamento := lancamento.NewLancamento()

	var exists bool
	m.base.TenantCtx(ctx).Model(&lancamento).Select("count(*) > 0").Where("id = ?", id).Find(&exists).Get("ID")
	return exists
}
