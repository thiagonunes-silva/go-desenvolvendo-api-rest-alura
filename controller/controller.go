package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/thiagonunes-silva/go-desenvolvendo-api-rest-alura/database"
	"github.com/thiagonunes-silva/go-desenvolvendo-api-rest-alura/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ListarPersonalidades(w http.ResponseWriter, r *http.Request) {
	var listaPersonalidades []models.Personalidade
	database.DB.Find(&listaPersonalidades)
	json.NewEncoder(w).Encode(listaPersonalidades)
}

func DetalharPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	personalidade := consultarPersonalidadeById(id)
	json.NewEncoder(w).Encode(personalidade)
}

func consultarPersonalidadeById(id string) models.Personalidade {
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	return personalidade
}

func CriarPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	database.DB.Create(&novaPersonalidade)
	json.NewEncoder(w).Encode(novaPersonalidade)
}

func DeletarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidadeDeletada models.Personalidade
	database.DB.Delete(&personalidadeDeletada, id)
}

func AtualizarPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	personalidadeParaAtualizar := consultarPersonalidadeById(id)
	json.NewDecoder(r.Body).Decode(&personalidadeParaAtualizar)
	database.DB.Save(&personalidadeParaAtualizar)
	json.NewEncoder(w).Encode(personalidadeParaAtualizar)
}
