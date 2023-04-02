package main

import (
	"github.com/Lubrum/github-actions-with-go/database"
	"github.com/Lubrum/github-actions-with-go/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
