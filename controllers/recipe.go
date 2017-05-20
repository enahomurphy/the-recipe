package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"recipe/helpers"
	"recipe/models"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// RecipeResponse is
type RecipeResponse struct {
	Status int             `json:"status"`
	Data   []models.Recipe `json:"data"`
}

// CreateRecipe creates a new Recipe
func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	Recipe := models.Recipe{
		Name:        r.FormValue("name"),
		UserID:      r.FormValue("userID"),
		CategoryID:  r.FormValue("categoryID"),
		Image:       r.FormValue("image_url"),
		Description: r.FormValue("description"),
	}
	decoder := json.NewDecoder(r.Body)
	decoderErr := decoder.Decode(&Recipe)

	if decoderErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	if Recipe.Name == "" || Recipe.UserID == "" || Recipe.CategoryID == "" {
		errMsg := models.Recipe{
			Name:       "Name  is required",
			UserID:     "User id is required",
			CategoryID: "Category id is required",
		}
		type Error struct {
			Status  int
			Message interface{}
		}
		newError := Error{Status: http.StatusBadRequest, Message: errMsg}
		response, _ := json.Marshal(newError)
		helpers.ResponseWriter(w, http.StatusBadRequest, string(response))
		return
	}
	_, dbErr := models.CreateRecipe(&Recipe)

	if dbErr != nil {
		helpers.ServerError(w, dbErr)
		return
	}
	helpers.StatusOkMessage(w, "Recipe created")
}

// GetRecipe Gets all Recipe and sends the data as response
// to the requesting Recipe
func GetRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, parseErr := strconv.Atoi(vars["id"])
	if parseErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	Recipe, err := models.GetRecipe(id)
	if err != nil {
		helpers.StatusNotFound(w, err)
		return
	}
	helpers.StatusOk(w, Recipe)
}

// GetAllRecipe Gets all Recipe and sends the data as response
// to the requesting Recipe
func GetAllRecipe(w http.ResponseWriter, r *http.Request) {
	Recipe, err := models.GetAllRecipe()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	response := RecipeResponse{
		Status: http.StatusOK,
		Data:   Recipe,
	}
	result, _ := json.Marshal(response)

	helpers.ResponseWriter(w, http.StatusOK, string(result))
}

//UpdateRecipe updates Recipe's detail
func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	Recipe := models.Recipe{
		Name:        r.FormValue("name"),
		UserID:      r.FormValue("userID"),
		CategoryID:  r.FormValue("categoryID"),
		Image:       r.FormValue("image_url"),
		Description: r.FormValue("description"),
	}
	decoder := json.NewDecoder(r.Body)
	decoderErr := decoder.Decode(&Recipe)

	if decoderErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	fmt.Println(id)

	_, err := models.UpdateRecipe(id, &Recipe)
	if err != nil {
		helpers.BadRequest(w, errors.New(err.Error()))
		return
	}
	helpers.StatusOkMessage(w, "Recipe updated")
}

// DeleteRecipe deletes a Recipe detail
func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, parseErr := strconv.Atoi(vars["id"])
	if parseErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	_, err := models.DeleteRecipe(id)
	if err != nil {
		helpers.BadRequest(w, err)
		return
	}
	helpers.StatusOkMessage(w, "Recipe deleted")
}
