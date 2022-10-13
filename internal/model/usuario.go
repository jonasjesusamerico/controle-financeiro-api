package model

import (
	"strings"
	"time"

	"github.com/joninhasamerico/controle-financeiro-api/pkg/auth"
	"gorm.io/gorm"
)

type Usuario struct {
	ID        uint64    `json:"id,omitempty"`
	Email     string    `json:"email,omitempty" validate:"required"`
	Senha     string    `json:"senha,omitempty" validate:"required"`
	TenantID  int64     `json:"tenant_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func (a *Usuario) TableName() string {
	return "usuarios"
}

func (u *Usuario) AfterCreate(tx *gorm.DB) (err error) {
	if u.TenantID == 0 {
		tx.Model(u).Update("tenant_id", u.ID)
	}
	return
}

func (u Usuario) GetId() uint64 {
	return u.ID
}

func (u *Usuario) Validar() (erro error) {
	u.Email = strings.TrimSpace(u.Email)

	senhaComHash, erro := auth.Hash(u.Senha)
	if erro != nil {
		return
	}

	u.Senha = string(senhaComHash)

	return
}

func (u Usuario) GetUsuarioRetorno() Usuario {
	return Usuario{
		ID:    u.ID,
		Email: u.Email,
		Senha: "",
	}
}
