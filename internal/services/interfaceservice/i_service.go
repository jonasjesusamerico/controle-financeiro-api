package interfaceservice

import (
	"context"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model/interfacemodel"
)

type IService interface {
	FindAll(ctx context.Context, models interface{}) (err error) //pointer
	GetByID(ctx context.Context, model interfacemodel.IModel, id int64) (err error)
	Update(ctx context.Context, model interfacemodel.IModel) (err error) //pointer
	Save(ctx context.Context, model interfacemodel.IModel) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
