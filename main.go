package main

import (
	"fmt"
	"go-api-rest/database"
	"go-api-rest/routes"
)

func main() {
	database.ConnectDB()

	fmt.Println("Servidor Rest com Go iniciado")
	routes.HandleResquest()
}
