package repository

import (
	"context"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
)

type IRepository interface {
	FindAll(ctx context.Context) (res []model.Usuario, err error)
	GetByID(ctx context.Context, id int64) (model.Usuario, error)
	Update(ctx context.Context, ar *model.Usuario) error
	Save(ctx context.Context, a *model.Usuario) error
	Delete(ctx context.Context, id int64) error
}
