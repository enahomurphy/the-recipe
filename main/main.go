package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"recipe/controllers/user"
	"recipe/models"
	"strconv"

	"github.com/gorilla/mux"
)

var baseURL = "/api/v1"

func main() {

	// users, err := models.Get(1)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	router := routes()
	http.Handle("/", router)
	http.ListenAndServe(":8081", nil)
}

func init() {
	// models.CreateTables(models.DB())
}

func routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(baseURL+"/users", user.Get).Methods("GET")
	router.HandleFunc(baseURL+"/users", test).Methods("POST")
	router.HandleFunc(baseURL+"/users/{id:[0-9]+}", test).Methods("PUT")
	router.HandleFunc(baseURL+"/users/{id:[0-9]+}", user.Update).Methods("DELETE")

	router.HandleFunc(baseURL+"/recipes", user.Get).Methods("GET")
	router.HandleFunc(baseURL+"/recipes", user.Create).Methods("POST")
	router.HandleFunc(baseURL+"/recipes/{id:[0-9]+}", user.Update).Methods("PUT")
	router.HandleFunc(baseURL+"/recipes/{id:[0-9]+}", user.Update).Methods("DELETE")

	return router
}

func test(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		panic(err)
	}

	var userData models.User

	decoder := json.NewDecoder(r.Body)

	decoderErr := decoder.Decode(&userData)

	if decoderErr != nil {
		panic(decoderErr)
	}

	fmt.Println(userData, 1, id)

	if err != nil {
		panic(err.Error())
	} else {
		models.UpdateUser(id, &userData)
	}
}
