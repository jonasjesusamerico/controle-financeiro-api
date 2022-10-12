package model

import (
	"errors"
	"time"

	"github.com/badoux/checkmail"
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

func (a *Usuario) BeforeCreate(tx *gorm.DB) (err error) {
	if a.TenantID == 0 {
		err = errors.New("missing tenantID in usuario model")
	}

	return
}

func (u Usuario) GetId() uint64 {
	return u.ID
}

func (u *Usuario) Validate() (err error) {
	u.Validar()

	if erro := checkmail.ValidateFormat(u.Email); erro != nil {
		return errors.New("o e-mail inserido é inválido")
	}

	if u.Email == "" {
		err = errors.New("o email é obrigatório e não pode estar em branco")
	}

	if u.Senha == "" {
		err = errors.New("a senha é obrigatório e não pode estar em branco")
	}

	return
}

func (u *Usuario) Validar() (erro error) {

	return
}

func (u Usuario) GetUsuarioRetorno() Usuario {
	return Usuario{
		ID:    u.ID,
		Email: u.Email,
		Senha: "",
	}
}
