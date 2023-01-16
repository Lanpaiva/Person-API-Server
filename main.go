package main

import (
	"fmt"

	"github.com/lanpaiva/go-rest-api/models"
	"github.com/lanpaiva/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "historia 1"},
		{Nome: "Nome 2", Historia: "historia 2"},
	}
	fmt.Println("Ola, Mundo")

	routes.HandleRequest()
}
