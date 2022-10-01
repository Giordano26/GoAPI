package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/dixonwille/wmenu"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func trimLine(text string) string {
	if text != "\n" {
		text = strings.TrimSuffix(text, "\n")

	}

	return text
}

func main() {

	//conectando no banco de dados
	db, err := sql.Open("sqlite3", "../db/cartas_db.db")
	checkError(err)

	menu := wmenu.NewMenu("O que gostaria de fazer?")

	menu.Action(func(opts []wmenu.Opt) error { handleFunc(db, opts); return nil })

	menu.Option("Criar Carta", 0, true, nil)
	menu.Option("Encontrar Carta", 1, false, nil)
	menu.Option("Atualizar Carta", 2, false, nil)
	menu.Option("Deletar Carta", 3, false, nil)
	menu.Option("Desligar", 4, false, nil)

	menuErr := menu.Run()
	checkError(menuErr)

	//defer close
	defer db.Close()

}

func handleFunc(db *sql.DB, opts []wmenu.Opt) {

	switch opts[0].Value {
	case 0:
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Insira o número da carta: ")
		numCarta, _ := reader.ReadString('\n')
		numCarta = trimLine(numCarta)

		fmt.Print("Insira o nome da carta: ")
		nomeCarta, _ := reader.ReadString('\n')
		nomeCarta = trimLine(nomeCarta)

		fmt.Print("Insira uma descrição da carta: ")
		descCarta, _ := reader.ReadString('\n')
		descCarta = trimLine(descCarta)

		newCarta := carta{
			Nome:   nomeCarta,
			Numero: numCarta,
			Desc:   descCarta,
		}

		addCarta(db, newCarta)

	case 1:
		fmt.Println("Encontrando carta...")
	case 2:
		fmt.Println("Atualizando carta...")
	case 3:
		fmt.Println("Deletando carta...")
	case 4:
		fmt.Println("Desligando...")

	}

}