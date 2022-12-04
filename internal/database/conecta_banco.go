package database

import (
	"github.com/joninhasamerico/controle-financeiro-api/configs"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model/comprador"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model/conta"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model/lancamento"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model/setor"
	"github.com/joninhasamerico/controle-financeiro-api/internal/model/usuario"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect é configurado para se conecter ao postgres como padrão
func Conect() (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(configs.StringConexaoBanco), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err != nil {
		return nil, err
	}

	/* auto migrate model */
	db.AutoMigrate(usuario.NewUsuario())
	db.AutoMigrate(conta.NewConta())
	db.AutoMigrate(setor.NewSetor())
	db.AutoMigrate(comprador.NewComprador())
	db.AutoMigrate(lancamento.NewLancamento())

	return
}
