package repository

import "github.com/joninhasamerico/controle-financeiro-api/internal/model"

type IUsuarioRepository interface {
	IRepository

	GetByEmail(email string, usuario model.IModel) (err error)
	CreateUserLogin(usuario model.IModel) (err error)
}
