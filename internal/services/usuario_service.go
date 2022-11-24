package services

import (
	"context"
	"fmt"
	"time"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
	"github.com/joninhasamerico/controle-financeiro-api/internal/repository"
)

type UsuarioService struct {
	usuarioRepo    repository.IRepository
	contextTimeout time.Duration
}

func NewUsuarioService(a repository.IRepository, timeout time.Duration) IService {
	return &UsuarioService{
		usuarioRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *UsuarioService) FindAll(ctx context.Context, models interface{}) (err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	err = a.usuarioRepo.FindAll(ctx, models)
	if err != nil {
		return err
	}

	return
}

func (a *UsuarioService) GetByID(ctx context.Context, model model.IModel, id int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	err = a.usuarioRepo.GetByID(ctx, model, id)
	if err != nil {
		return
	}

	return
}

func (a *UsuarioService) Update(ctx context.Context, model model.IModel) (err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()

	return a.usuarioRepo.Update(ctx, model)
}

func (a *UsuarioService) Save(ctx context.Context, models model.IModel) (err error) {
	usuario := models.(*model.Usuario)
	repo := a.usuarioRepo.(*repository.UsuarioRepository)
	existedUsuario, _ := a.GetByEmail(usuario.Email)
	if existedUsuario.GetId() != 0 {
		return model.ErrConflict
	}
	usuario.Validar()
	err = repo.CreateUserLogin(usuario)
	return
}

func (a *UsuarioService) Delete(ctx context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, a.contextTimeout)
	defer cancel()
	existedUsuario := model.Usuario{}
	err = a.usuarioRepo.GetByID(ctx, &existedUsuario, id)
	if err != nil {
		return
	}
	if existedUsuario == (model.Usuario{}) {
		return model.ErrNotFound
	}
	return a.usuarioRepo.Delete(ctx, existedUsuario, id)
}

func (a *UsuarioService) GetByEmail(email string) (usuario model.IModel, err error) {
	repo := a.usuarioRepo.(*repository.UsuarioRepository)
	var u model.Usuario
	err = repo.GetByEmail(email, &u)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (a *UsuarioService) Teste() {
	fmt.Println("Funcionou!")
}
