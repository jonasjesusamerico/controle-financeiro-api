package model

import "time"

type Tenant struct {
	TenantID  int64     `json:"tenant_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}
