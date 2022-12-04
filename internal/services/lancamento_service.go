package services

import (
	"context"
	"errors"
	"time"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model/interfacemodel"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model/lancamento"
	"github.com/joninhasamerico/controle-financeiro-api/internal/repository/interfacerepository"
	"github.com/joninhasamerico/controle-financeiro-api/internal/services/interfaceservice"
)

type LancamentoService struct {
	repository     interfacerepository.ILancamentoRepository
	contextTimeout time.Duration
}

func NewLancamentoService(repository interfacerepository.ILancamentoRepository, timeout time.Duration) interfaceservice.ILancamentoService {
	return &LancamentoService{
		repository:     repository,
		contextTimeout: timeout,
	}
}

func (a *LancamentoService) FindAll(ctx context.Context, models interface{}) (err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	err = a.repository.FindAll(ctx, []string{"Conta", "Setor"}, models)
	if err != nil {
		return err
	}

	return
}

func (a *LancamentoService) GetByID(ctx context.Context, model interfacemodel.IModel, id int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	err = a.repository.GetByID(ctx, model, id)
	return
}

func (a *LancamentoService) Update(ctx context.Context, model interfacemodel.IModel) (err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	return a.repository.Update(ctx, model)
}

func (a *LancamentoService) Save(ctx context.Context, models interfacemodel.IModel) (err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	err = a.repository.Save(ctx, models)
	return
}

func (a *LancamentoService) Delete(ctx context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	if !a.Exists(ctx, id) {
		err = errors.New("there is no record with the given id")
		return
	}

	return a.repository.Delete(ctx, lancamento.NewLancamento(), id)
}

func (a *LancamentoService) Exists(ctx context.Context, id int64) bool {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	return a.repository.Exists(ctx, id)
}
