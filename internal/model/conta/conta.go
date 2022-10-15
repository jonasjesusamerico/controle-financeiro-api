package conta

import (
	"time"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model/enum"
)

type Conta struct {
	ID        int64
	TipoConta enum.TipoConta
	Descricao string
	TenantID  int64     `json:"tenant_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func NewConta() *Conta {
	return &Conta{}
}
