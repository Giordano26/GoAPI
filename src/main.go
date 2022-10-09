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
	menu.Option("Listar Cartas", 1, false, nil)
	menu.Option("Encontrar Carta", 2, false, nil)
	menu.Option("Atualizar Carta", 3, false, nil)
	menu.Option("Deletar Carta", 4, false, nil)
	menu.Option("Desligar", 5, false, nil)

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

		carta := listarCartas(db)

		fmt.Printf("Encontrado %v resultados", len(carta))

		for _, buscaCarta := range carta {
			fmt.Printf("\n----\nID: %d \nNumero: %s\nNome: %s\nDescrição: %s\n", buscaCarta.id, buscaCarta.Numero, buscaCarta.Nome, buscaCarta.Desc)
		}

	case 2:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Entre o nome da carta ou numero para pesquisa: ")
		searchString, _ := reader.ReadString('\n')
		searchString = trimLine(searchString)

		carta := searchCarta(db, searchString)

		fmt.Printf("Encontrado %v resultados", len(carta))

		for _, buscaCarta := range carta {
			fmt.Printf("\n----\nID: %d \nNumero: %s\nNome: %s\nDescrição: %s\n", buscaCarta.id, buscaCarta.Numero, buscaCarta.Nome, buscaCarta.Desc)
		}

	case 3:
		fmt.Println("Atualizando carta...")
	case 4:
		fmt.Println("Deletando carta...")
	case 5:
		fmt.Println("Desligando...")

	}

}
