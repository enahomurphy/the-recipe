package main

import (
	"net/http"
	"recipe/controllers/user"
	"recipe/models"

	"fmt"

	"log"

	"github.com/gorilla/mux"
)

var baseURL = "/api/v1"

func main() {

	users, err := models.Get(1)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)

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

	router.HandleFunc(baseURL+"/recipes", user.Get).Methods("GET")
	router.HandleFunc(baseURL+"/recipes", user.Create).Methods("POST")
	router.HandleFunc(baseURL+"/recipes/{id:[0-9]+}", user.Update).Methods("PUT")
	router.HandleFunc(baseURL+"/recipes/{id:[0-9]+}", user.Update).Methods("DELETE")

	return router
}
