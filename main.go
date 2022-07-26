package main

import (
	"fmt"
	"net/http"

	"github.com/thiagonunes-silva/go-desenvolvendo-api-rest-alura/database"
	"github.com/thiagonunes-silva/go-desenvolvendo-api-rest-alura/models"
	"github.com/thiagonunes-silva/go-desenvolvendo-api-rest-alura/routes"
)

func main() {
	fmt.Println("Iniciando servidor rest go")
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "História 1"},
		{Id: 2, Nome: "Nome 2", Historia: "História 2"},
	}
	database.ConectarComBancoDeDados()
	routes.HandleRequest()
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}
