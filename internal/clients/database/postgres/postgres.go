package postgres

import (
	"database/sql"
	"fmt"
	"loaddados/domains"
)

type Database struct {
	DB *sql.DB
}

func NewConnection(host, port, user, password, dbname string) *Database {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &Database{
		DB: db,
	}
}

func (d *Database) Insert(dado domains.Dado) error {
	query := "INSERT INTO desafio.dados (CPF, IS_PRIVATE, INCOMPLETO, DATA_COMPRA, TICKET_MEDIO, TICKET_ULTIMA, LOJA_FREQUENCIA, ULTIMA_LOJA) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)"

	_, err := d.DB.Exec(query, dado.CPF, dado.IS_PRIVATE, dado.INCOMPLETO, dado.DATA_COMPRA, dado.TICKET_MEDIO, dado.TICKET_ULTIMA, dado.LOJA_FREQUENCIA, dado.ULTIMA_LOJA)
	if err != nil {
		panic(err)
	}

	return err
}

func (d *Database) Close() {
	if err := d.DB.Close(); err != nil {
		return
	}
}
