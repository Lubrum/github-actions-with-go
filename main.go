package main

import (
	"log"

	"github.com/Lubrum/github-actions-with-go/database"
	"github.com/Lubrum/github-actions-with-go/routes"
)

func main() {
	if err := database.ConectaComBancoDeDados(); err != nil {
		log.Fatal(err)
	}

	if err := routes.HandleRequest(); err != nil {
		log.Fatal(err)
	}
}
