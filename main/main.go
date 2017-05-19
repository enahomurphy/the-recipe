package main

import (
	"net/http"
	"recipe/controllers"
	"recipe/models"

	"github.com/gorilla/mux"
)

var baseURL = "/api/v1"

func main() {
	router := routes()
	http.Handle("/", router)
	http.ListenAndServe(":8085", nil)
}

func init() {
	models.CreateTables(models.DB())
}

func routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(baseURL+"/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc(baseURL+"/users", controllers.CreateUser).Methods("POST")
	router.HandleFunc(baseURL+"/users/{id:[0-9]+}", controllers.GetUser).Methods("GET")
	router.HandleFunc(baseURL+"/users/{id:[0-9]+}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc(baseURL+"/users/{id:[0-9]+}", controllers.DeleteUser).Methods("DELETE")

	router.HandleFunc(baseURL+"/categories", controllers.GetAllcategory).Methods("GET")
	router.HandleFunc(baseURL+"/categories", controllers.CreateCategory).Methods("POST")
	router.HandleFunc(baseURL+"/categories/{id:[0-9]+}", controllers.GetCategory).Methods("GET")
	router.HandleFunc(baseURL+"/categories/{id:[0-9]+}", controllers.UpdateCategory).Methods("PUT")
	router.HandleFunc(baseURL+"/categories/{id:[0-9]+}", controllers.DeleteCategory).Methods("DELETE")

	// router.HandleFunc(baseURL+"/recipes", controllers.GetUser).Methods("GET")
	// router.HandleFunc(baseURL+"/recipes", controllers.Create).Methods("POST")
	// router.HandleFunc(baseURL+"/recipes/{id:[0-9]+}", controllers.UpdateCategory).Methods("PUT")
	// router.HandleFunc(baseURL+"/recipes/{id:[0-9]+}", controllers.Update).Methods("DELETE")

	return router
}
