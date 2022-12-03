package interfacerepository

import "github.com/joninhasamerico/controle-financeiro-api/internal/model/interfacemodel"

type IUsuarioRepository interface {
	IRepository

	GetByEmail(email string, usuario interfacemodel.IModel) (err error)
	CreateUserLogin(usuario interfacemodel.IModel) (err error)
}
