package services

import (
	"context"
	"time"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
	"github.com/joninhasamerico/controle-financeiro-api/internal/repository"
)

type UsuarioService struct {
	articleRepo    repository.IUsuarioRepository
	contextTimeout time.Duration
}

type IUsuarioService interface {
	Fetch(ctx context.Context) ([]model.Usuario, error)
	GetByID(ctx context.Context, id int64) (model.Usuario, error)
	Update(ctx context.Context, ar *model.Usuario) error
	GetByTitle(ctx context.Context, title string) (model.Usuario, error)
	Store(context.Context, *model.Usuario) error
	Delete(ctx context.Context, id int64) error
}

// NewUsuarioService will create new an UsuarioService object representation of models.NewUsuarioService interface
func NewUsuarioService(a repository.IUsuarioRepository, timeout time.Duration) IUsuarioService {
	return &UsuarioService{
		articleRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *UsuarioService) Fetch(c context.Context) (res []model.Usuario, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.articleRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (a *UsuarioService) GetByID(c context.Context, id int64) (res model.Usuario, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.articleRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (a *UsuarioService) Update(c context.Context, ar *model.Usuario) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	ar.UpdatedAt = time.Now()
	return a.articleRepo.Update(ctx, ar)
}

func (a *UsuarioService) GetByTitle(c context.Context, title string) (res model.Usuario, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	res, err = a.articleRepo.GetByTitle(ctx, title)
	if err != nil {
		return
	}

	return
}

func (a *UsuarioService) Store(c context.Context, m *model.Usuario) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedArticle, _ := a.GetByTitle(ctx, m.Title)
	if existedArticle != (model.Usuario{}) {
		return nil //models.ErrConflict
	}

	err = a.articleRepo.Store(ctx, m)
	return
}

func (a *UsuarioService) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedArticle, err := a.articleRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedArticle == (model.Usuario{}) {
		return nil //models.ErrNotFound
	}
	return a.articleRepo.Delete(ctx, id)
}
