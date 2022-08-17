package main

import (
	"fmt"
	"go-api-rest/database"
	"go-api-rest/routes"
)

func main() {

	database.ConectaComBancoDeDados()

	fmt.Println("Servidor Rest com Go iniciado")
	routes.HandleResquest()
}
