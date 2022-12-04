package usuario

import (
	"strings"
	"time"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model/enum"
	"github.com/joninhasamerico/controle-financeiro-api/pkg/auth"
	"gorm.io/gorm"
)

type usuario struct {
	ID        uint64    `json:"id,omitempty"`
	Email     string    `json:"email,omitempty" validate:"required"`
	Senha     string    `json:"senha,omitempty" validate:"required"`
	TenantID  int64     `json:"tenant_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUsuario() *usuario {
	return &usuario{}
}

func NewSliceUsuario() []usuario {
	return []usuario{}
}

func (a *usuario) TableName() string {
	return "usuarios"
}

func (u *usuario) AfterCreate(tx *gorm.DB) (err error) {
	if u.TenantID == 0 {
		tx.Model(u).Update(string(enum.TENANT_ID), u.ID)
	}
	return
}

func (u usuario) GetId() uint64 {
	return u.ID
}

func (u usuario) GetEmail() string {
	return u.Email
}

func (u *usuario) Validar() (erro error) {
	u.Email = strings.TrimSpace(u.Email)

	senhaComHash, erro := auth.Hash(u.Senha)
	if erro != nil {
		return
	}

	u.Senha = string(senhaComHash)

	return
}

func (l *usuario) SetTenant(tenantId int64) {
	l.TenantID = tenantId
}

func (u usuario) GetUsuarioRetorno() usuario {
	return usuario{
		ID:    u.ID,
		Email: u.Email,
		Senha: "",
	}
}
