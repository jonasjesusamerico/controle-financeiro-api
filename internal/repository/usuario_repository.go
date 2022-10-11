package repository

import (
	"fmt"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
)

type UsuarioRepository struct {
}

func (UsuarioRepository) Insert(model model.IModel) (id uint64, err error) {
	if err = model.Validate(); err != nil {
		return
	}

	fmt.Println("Fake: Registro inserido com sucesso!")
	return
}

func (UsuarioRepository) Update(model model.IModel) (err error) {
	if err = model.Validate(); err == nil {
		fmt.Println("Fake: Registro atualizado com sucesso!")
	}
	return
}

func (UsuarioRepository) Save(model model.IModel) (id uint64, err error) {
	if err = model.Validate(); err != nil {
		return
	}

	fmt.Println("Fake: Registro salvo com sucesso!")

	return
}

func (UsuarioRepository) SaveAll(models interface{}) (err error) {
	fmt.Println("Fake: Registros salvo com sucesso!")
	return
}

func (UsuarioRepository) FindById(receiver model.IModel, id interface{}) (err error) {
	fmt.Println("Fake: Registro encontrado com sucesso!")
	return
}

func (UsuarioRepository) FindFirst(receiver model.IModel, query string, args ...interface{}) (err error) {
	fmt.Println("Fake: Primeiro registro encontrado com sucesso!")
	return
}

func (UsuarioRepository) FindAll(models interface{}, query string, args ...interface{}) (err error) {
	fmt.Println("Fake: Registros encontrado com sucesso!")
	return
}

func (UsuarioRepository) Delete(model model.IModel, query string, args ...interface{}) (err error) {
	fmt.Println("Fake: Registro deletado com sucesso!")
	return
}
