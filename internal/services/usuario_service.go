package services

import (
	"context"
	"fmt"
	"time"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
	"github.com/joninhasamerico/controle-financeiro-api/internal/repository"
)

type UsuarioService struct {
	repository     repository.IUsuarioRepository
	contextTimeout time.Duration
}

func NewUsuarioService(repository repository.IUsuarioRepository, timeout time.Duration) IUsuarioService {
	return &UsuarioService{
		repository:     repository,
		contextTimeout: timeout,
	}
}

func (a *UsuarioService) FindAll(ctx context.Context, models interface{}) (err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	err = a.repository.FindAll(ctx, models)
	if err != nil {
		return err
	}

	return
}

func (a *UsuarioService) GetByID(ctx context.Context, model model.IModel, id int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	err = a.repository.GetByID(ctx, model, id)
	if err != nil {
		return
	}

	return
}

func (a *UsuarioService) Update(ctx context.Context, model model.IModel) (err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	return a.repository.Update(ctx, model)
}

func (a *UsuarioService) Save(ctx context.Context, models model.IModel) (err error) {
	usuario := models.(*model.Usuario)
	existedUsuario, err := a.GetByEmail(usuario.Email)
	if existedUsuario != nil {
		return err
	}
	usuario.Validar()
	err = a.repository.CreateUserLogin(usuario)
	return
}

func (a *UsuarioService) Delete(ctx context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()
	existedUsuario := model.Usuario{}
	err = a.repository.GetByID(ctx, &existedUsuario, id)
	if err != nil {
		return
	}
	if existedUsuario == (model.Usuario{}) {
		return model.ErrNotFound
	}
	return a.repository.Delete(ctx, existedUsuario, id)
}

func (a *UsuarioService) GetByEmail(email string) (usuario model.IModel, err error) {
	var u model.Usuario
	err = a.repository.GetByEmail(email, &u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (a *UsuarioService) Teste() {
	fmt.Println("Funcionou!")
}
