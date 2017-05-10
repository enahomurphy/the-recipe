package main

import (
	"net/http"
	"recipe/controllers/user"

	"github.com/gorilla/mux"
)

var baseURL = "/api/v1"

func main() {
	router := mux.NewRouter()
	router.HandleFunc(baseURL+"/users", user.Create).Methods("GET")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
