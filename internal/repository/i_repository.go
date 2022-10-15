package repository

import (
	"context"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
)

type IRepository interface {
	FindAll(ctx context.Context, models interface{}) (err error)
	GetByID(ctx context.Context, model model.IModel, id int64) (err error)
	Update(ctx context.Context, model model.IModel) (err error)
	Save(ctx context.Context, model model.IModel) (err error)
	Delete(ctx context.Context, model model.IModel, id int64) (err error)
}
