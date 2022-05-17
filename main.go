package main

import (
	"fmt"
	"github.com/go-personal-info/database"
	"github.com/go-personal-info/routes"
)

func main() {
	database.DBConnection()
	fmt.Println("Iniciando")
	routes.HandleRequest()
}
