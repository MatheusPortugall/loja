package main

import (
	"loja/database"
	routes "loja/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
