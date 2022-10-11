package repository

import (
	"fmt"

	"github.com/joninhasamerico/controle-financeiro-api/internal/model"
)

type Basic struct {
}

func (Basic) Insert(model model.IModel) (id uint64, err error) {
	if err = model.Validate(); err != nil {
		return
	}

	fmt.Println("Fake: Registro inserido com sucesso!")
	return
}

func (Basic) Update(model model.IModel) (err error) {
	if err = model.Validate(); err == nil {
		fmt.Println("Fake: Registro atualizado com sucesso!")
	}
	return
}

func (Basic) Save(model model.IModel) (id uint64, err error) {
	if err = model.Validate(); err != nil {
		return
	}

	fmt.Println("Fake: Registro salvo com sucesso!")

	return
}

func (Basic) SaveAll(models interface{}) (err error) {
	fmt.Println("Fake: Registros salvo com sucesso!")
	return
}

func (Basic) FindById(receiver model.IModel, id interface{}) (err error) {
	fmt.Println("Fake: Registro encontrado com sucesso!")
	return
}

func (Basic) FindFirst(receiver model.IModel, query string, args ...interface{}) (err error) {
	fmt.Println("Fake: Primeiro registro encontrado com sucesso!")
	return
}

func (Basic) FindAll(models interface{}, query string, args ...interface{}) (err error) {
	fmt.Println("Fake: Registros encontrado com sucesso!")
	return
}

func (Basic) Delete(model model.IModel, query string, args ...interface{}) (err error) {
	fmt.Println("Fake: Registro deletado com sucesso!")
	return
}
