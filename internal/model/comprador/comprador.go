package comprador

import "time"

type comprador struct {
	ID        int64
	Nome      string
	TenantID  int64     `json:"tenant_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func NewComprador() *comprador {
	return &comprador{}
}
