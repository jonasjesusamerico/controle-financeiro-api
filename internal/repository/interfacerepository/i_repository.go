package interfacerepository

import (
	"context"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model/interfacemodel"
)

type IRepository interface {
	FindAll(ctx context.Context, preloads []string, models interface{}) (err error)
	GetByID(ctx context.Context, model interfacemodel.IModel, id int64) (err error)
	Update(ctx context.Context, model interfacemodel.IModel) (err error)
	Save(ctx context.Context, model interfacemodel.IModel) (err error)
	Delete(ctx context.Context, model interfacemodel.IModel, id int64) (err error)
	Exists(ctx context.Context, id int64) bool
}
