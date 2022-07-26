package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/thiagonunes-silva/go-desenvolvendo-api-rest-alura/controller"
	"github.com/thiagonunes-silva/go-desenvolvendo-api-rest-alura/middleware"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controller.Home)
	r.HandleFunc("/api/personalidades", controller.ListarPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controller.DetalharPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades", controller.CriarPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades/{id}", controller.DeletarPersonalidade).Methods("DELETE")
	r.HandleFunc("/api/personalidades/{id}", controller.AtualizarPersonalidade).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
