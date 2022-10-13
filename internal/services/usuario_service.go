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

func (a *UsuarioService) FindAll(c context.Context) (res []model.Usuario, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.usuarioRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (a *UsuarioService) GetByID(c context.Context, id int64) (usuario model.Usuario, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	usuario, err = a.usuarioRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (a *UsuarioService) GetByEmail(email string) (usuario model.Usuario, err error) {
	repo := a.usuarioRepo.(*repository.UsuarioRepository)
	usuario, err = repo.GetByEmail(email)
	if err != nil {
		return
	}

	return
}

func (a *UsuarioService) Update(c context.Context, ar *model.Usuario) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	ar.UpdatedAt = time.Now()
	return a.usuarioRepo.Update(ctx, ar)
}

func (a *UsuarioService) Save(c context.Context, m *model.Usuario) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedUsuario, _ := a.GetByID(ctx, int64(m.ID))
	if existedUsuario != (model.Usuario{}) {
		return model.ErrConflict
	}
	m.Validar()
	err = a.usuarioRepo.Save(ctx, m)
	return
}

func (a *UsuarioService) CreateUserLogin(m *model.Usuario) (err error) {
	repo := a.usuarioRepo.(*repository.UsuarioRepository)
	existedUsuario, _ := a.GetByEmail(m.Email)
	if existedUsuario != (model.Usuario{}) {
		return model.ErrConflict
	}
	m.Validar()
	err = repo.CreateUserLogin(m)
	return
}

func (a *UsuarioService) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedUsuario, err := a.usuarioRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedUsuario == (model.Usuario{}) {
		return model.ErrNotFound
	}
	return a.usuarioRepo.Delete(ctx, id)
}

func (a *UsuarioService) Teste() {
	fmt.Println("Funcionou!")
}
