package setor

import (
	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
)

type Setor struct {
	ID        int64
	Descricao string
	model.Tenant
}

func NewSetor() *Setor {
	return &Setor{}
}
