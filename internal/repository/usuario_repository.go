package repository

import (
	"context"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
	"gorm.io/gorm"
)

// UsuarioRepository represent the Usuario's repositories contract
type IUsuarioRepository interface {
	Fetch(ctx context.Context) (res []model.Usuario, err error)
	GetByID(ctx context.Context, id int64) (model.Usuario, error)
	GetByTitle(ctx context.Context, title string) (model.Usuario, error)
	Update(ctx context.Context, ar *model.Usuario) error
	Store(ctx context.Context, a *model.Usuario) error
	Delete(ctx context.Context, id int64) error
}

type mysqlUsuarioRepository struct {
	base BaseRepository
}

// NewMysqlArticleRepository will create an object that represent the usuario.Repository interface
func NewUsuarioRepository(dbCtx *gorm.DB) IUsuarioRepository {
	return &mysqlUsuarioRepository{base: BaseRepository{dbCtx: dbCtx}}
}

func (m *mysqlUsuarioRepository) Fetch(ctx context.Context) (res []model.Usuario, err error) {
	var usuarios []model.Usuario
	if err = m.base.tenantCtx(ctx).Find(&usuarios).Error; err != nil {
		return usuarios, err
	}

	return usuarios, nil
}
func (m *mysqlUsuarioRepository) GetByID(ctx context.Context, id int64) (model.Usuario, error) {
	var usuario model.Usuario
	if err := m.base.tenantCtx(ctx).Where("id = ?", id).First(&usuario).Error; err != nil {
		return usuario, err
	}

	return usuario, nil
}

func (m *mysqlUsuarioRepository) GetByTitle(ctx context.Context, title string) (model.Usuario, error) {
	var usuario model.Usuario
	if err := m.base.tenantCtx(ctx).Where("title = ?", title).First(&usuario).Error; err != nil {
		return usuario, err
	}

	return usuario, nil
}

func (m *mysqlUsuarioRepository) Store(ctx context.Context, usuario *model.Usuario) (err error) {
	usuario.TenantID = m.base.TenantID(ctx)

	if err = m.base.tenantCtx(ctx).Create(usuario).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUsuarioRepository) Delete(ctx context.Context, id int64) (err error) {
	if err = m.base.tenantCtx(ctx).Where("id = ?", id).Delete(&model.Usuario{}).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlUsuarioRepository) Update(ctx context.Context, ar *model.Usuario) (err error) {
	if err = m.base.tenantCtx(ctx).Save(ar).Error; err != nil {
		return err
	}
	return nil
}
