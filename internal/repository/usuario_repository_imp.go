package repository

import (
	"context"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	base BaseRepository
}

func NewUsuarioRepository(dbCtx *gorm.DB) IUsuarioRepository {
	return &UsuarioRepository{
		base: BaseRepository{
			dbCtx: dbCtx,
		},
	}
}

func (m *UsuarioRepository) FindAll(ctx context.Context, models interface{}) (err error) {

	if err = m.base.tenantCtx(ctx).Select("id", "email").Statement.Find(models).Error; err != nil {
		return
	}

	return
}

func (m *UsuarioRepository) GetByID(ctx context.Context, models model.IModel, id int64) (err error) {
	if err := m.base.tenantCtx(ctx).Where("id = ?", id).First(&models).Error; err != nil {
		return model.ErrNotFound
	}

	return nil
}

func (m *UsuarioRepository) Update(ctx context.Context, model model.IModel) (err error) {
	if err = m.base.tenantCtx(ctx).Save(model).Error; err != nil {
		return err
	}
	return nil
}

func (m *UsuarioRepository) Save(ctx context.Context, model model.IModel) (err error) {
	if err = m.base.tenantCtx(ctx).Create(model).Error; err != nil {
		return err
	}
	return nil
}

func (m *UsuarioRepository) Delete(ctx context.Context, model model.IModel, id int64) (err error) {
	if err = m.base.tenantCtx(ctx).Where("id = ?", id).Delete(&model).Error; err != nil {
		return err
	}
	return nil
}

func (m *UsuarioRepository) GetByEmail(email string, usuario model.IModel) (err error) {

	if err := m.base.dbCtx.Where("email = ?", email).First(&usuario).Error; err != nil {
		return model.ErrNotFound
	}

	return nil
}

func (m *UsuarioRepository) CreateUserLogin(usuario model.IModel) (err error) {
	if err = m.base.dbCtx.Create(usuario).Error; err != nil {
		return err
	}
	return nil
}
