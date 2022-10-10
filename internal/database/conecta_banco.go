package database

import (
	"database/sql"

	"github.com/joninhasamerico/controle-financeiro-api/configs"
	_ "github.com/lib/pq"
)

// Connect é configurado para se conecter ao postgres como padrão
func Conect() (db *sql.DB, err error) {
	db, err = sql.Open("postgres", configs.StringConexaoBanco)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	return
}
