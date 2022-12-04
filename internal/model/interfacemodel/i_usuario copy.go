package interfacemodel

type ILancamento interface {
	GetEmail() string
	Validar() (erro error)
}
