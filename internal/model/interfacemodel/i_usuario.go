package interfacemodel

type IUsuario interface {
	GetEmail() string
	Validar() (erro error)
}
