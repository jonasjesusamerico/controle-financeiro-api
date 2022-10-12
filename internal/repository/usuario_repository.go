package repository

import (
	"context"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	base BaseRepository
}

func NewUsuarioRepository(dbCtx *gorm.DB) IRepository {
	return &UsuarioRepository{base: BaseRepository{dbCtx: dbCtx}}
}

func (m *UsuarioRepository) FindAll(ctx context.Context) (res []model.Usuario, err error) {
	var usuarios []model.Usuario
	if err = m.base.tenantCtx(ctx).Find(&usuarios).Error; err != nil {
		return usuarios, err
	}

	return usuarios, nil
}
func (m *UsuarioRepository) GetByID(ctx context.Context, id int64) (model.Usuario, error) {
	var usuario model.Usuario
	if err := m.base.tenantCtx(ctx).Where("id = ?", id).First(&usuario).Error; err != nil {
		return usuario, err
	}

	return usuario, nil
}

func (m *UsuarioRepository) Save(ctx context.Context, usuario *model.Usuario) (err error) {
	usuario.TenantID = m.base.TenantID(ctx)

	if err = m.base.tenantCtx(ctx).Create(usuario).Error; err != nil {
		return err
	}
	return nil
}

func (m *UsuarioRepository) Delete(ctx context.Context, id int64) (err error) {
	if err = m.base.tenantCtx(ctx).Where("id = ?", id).Delete(&model.Usuario{}).Error; err != nil {
		return err
	}
	return nil
}

func (m *UsuarioRepository) Update(ctx context.Context, ar *model.Usuario) (err error) {
	if err = m.base.tenantCtx(ctx).Save(ar).Error; err != nil {
		return err
	}
	return nil
}
