package main

import (
	"net/http"
	"recipe/controllers/user"
	"recipe/models"

	"github.com/gorilla/mux"
)

var baseURL = "/api/v1"

func main() {
	router := routes()
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

func init() {
	models.CreateTables(models.DB())
}

func routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(baseURL+"/users", user.Get).Methods("GET")
	router.HandleFunc(baseURL+"/users", user.Create).Methods("POST")
	router.HandleFunc(baseURL+"/users/{id:[0-9]+}", user.Update).Methods("PUT")
	router.HandleFunc(baseURL+"/users/{id:[0-9]+}", user.Update).Methods("DELETE")

	return router
}
