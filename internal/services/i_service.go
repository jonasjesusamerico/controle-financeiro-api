package services

import (
	"context"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
)

type IService interface {
	FindAll(ctx context.Context) ([]model.Usuario, error)
	GetByID(ctx context.Context, id int64) (model.Usuario, error)
	Update(ctx context.Context, ar *model.Usuario) error
	Save(context.Context, *model.Usuario) error
	Delete(ctx context.Context, id int64) error
}
