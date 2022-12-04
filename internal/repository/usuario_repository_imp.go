package repository

import (
	"context"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model/interfacemodel"
	"github.com/joninhasamerico/controle-financeiro-api/internal/repository/interfacerepository"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	base interfacerepository.IBaseRepository
}

func NewUsuarioRepository(dbCtx *gorm.DB) interfacerepository.IUsuarioRepository {
	return &UsuarioRepository{
		base: NewBaseRepository(dbCtx),
	}
}

func (m *UsuarioRepository) FindAll(ctx context.Context, models interface{}) (err error) {

	if err = m.base.TenantCtx(ctx).Select("id", "email").Statement.Find(models).Error; err != nil {
		return
	}

	return
}

func (m *UsuarioRepository) GetByID(ctx context.Context, models interfacemodel.IModel, id int64) (err error) {
	if err := m.base.TenantCtx(ctx).Where("id = ?", id).First(&models).Error; err != nil {
		return model.ErrNotFound
	}

	return nil
}

func (m *UsuarioRepository) Update(ctx context.Context, model interfacemodel.IModel) (err error) {
	if err = m.base.TenantCtx(ctx).Save(model).Error; err != nil {
		return err
	}
	return nil
}

func (m *UsuarioRepository) Save(ctx context.Context, model interfacemodel.IModel) (err error) {
	if err = m.base.TenantCtx(ctx).Create(model).Error; err != nil {
		return err
	}
	return nil
}

func (m *UsuarioRepository) Delete(ctx context.Context, model interfacemodel.IModel, id int64) (err error) {
	if err = m.base.TenantCtx(ctx).Where("id = ?", id).Delete(&model).Error; err != nil {
		return err
	}
	return nil
}

func (m *UsuarioRepository) GetByEmail(email string, usuario interfacemodel.IModel) (err error) {

	if err := m.base.Repo().Where("email = ?", email).First(usuario).Error; err != nil {
		return model.ErrNotFound
	}

	return nil
}

func (m *UsuarioRepository) CreateUserLogin(usuario interfacemodel.IUsuario) (err error) {
	if err = m.base.Repo().Create(usuario).Error; err != nil {
		return err
	}
	return nil
}
