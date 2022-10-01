package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type carta struct {
	id     int
	Numero string
	Nome   string
	Desc   string
}

func addCarta(db *sql.DB, newCarta carta) {
	stmt, _ := db.Prepare("INSERT INTO Cartas (id,Numero, Nome, Desc) VALUES (?,?,?,?)")

	stmt.Exec(nil, newCarta.Numero, newCarta.Nome, newCarta.Desc)

	defer stmt.Close()

	fmt.Printf("Adicionada a carta %v com sucesso!\n", newCarta.Nome)
}
