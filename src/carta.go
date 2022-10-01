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

func searchCarta(db *sql.DB, searchString string) []carta {

	rows, err := db.Query("SELECT id, Numero, Nome, Desc FROM Cartas WHERE Nome like '%" + searchString + "%' OR Numero like '%" + searchString + "%'")
	checkError(err)

	defer rows.Close()

	err = rows.Err()

	checkError(err)

	cartas := make([]carta, 0)

	for rows.Next() {
		buscaCarta := carta{}

		err = rows.Scan(&buscaCarta.id, &buscaCarta.Numero, &buscaCarta.Nome, &buscaCarta.Desc)

		checkError(err)

		cartas = append(cartas, buscaCarta)
	}

	err = rows.Err()

	checkError(err)

	return cartas

}
