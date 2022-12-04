package conta

import (
	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model/enum"
)

type Conta struct {
	ID        int64
	TipoConta enum.TipoConta
	Descricao string
	model.Tenant
}

func NewConta() *Conta {
	return &Conta{}
}
