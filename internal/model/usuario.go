package model

import (
	"errors"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID    uint64 `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Senha string `json:"senha,omitempty"`
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
