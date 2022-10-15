package services

import (
	"context"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
)

type IService interface {
	FindAll(ctx context.Context, models interface{}) (err error) //pointer
	GetByID(ctx context.Context, model model.IModel, id int64) (err error)
	Update(ctx context.Context, model model.IModel) (err error) //pointer
	Save(ctx context.Context, model model.IModel) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
