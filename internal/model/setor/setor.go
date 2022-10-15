package setor

import "time"

type setor struct {
	ID        int64
	Descricao string
	TenantID  int64     `json:"tenant_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func NewSetor() *setor {
	return &setor{}
}
