package controllers

import (
	"encoding/json"
	"errors"
	"recipe/helpers"
	"recipe/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// IngredientResponse is
type IngredientResponse struct {
	Status   int                 `json:"status"`
	MetaData interface{}         `json:"meta_data,omitempty"`
	Data     []models.Ingredient `json:"data"`
}

// CreateIngredient creates a new Ingredient
func CreateIngredient(w http.ResponseWriter, r *http.Request) {
	Ingredient := models.Ingredient{
		Name:     r.FormValue("name"),
		Quantity: r.FormValue("quantity"),
		RecipeID: r.FormValue("recipeID"),
		Unit:     r.FormValue("unit"),
	}
	decoder := json.NewDecoder(r.Body)
	decoderErr := decoder.Decode(&Ingredient)

	if decoderErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	if Ingredient.Name == "" || Ingredient.Quantity == "" || Ingredient.RecipeID == "" {
		errMsg := models.Ingredient{
			Name:     "Name  is required",
			Quantity: "quantity id is required",
			RecipeID: "recipeID id is required",
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
	_, dbErr := models.CreateIngredient(&Ingredient)

	if dbErr != nil {
		helpers.ServerError(w, dbErr)
		return
	}
	helpers.StatusOkMessage(w, "Ingredient created")
}

// GetIngredient Gets all Ingredient and sends the data as response
// to the requesting Ingredient
func GetIngredient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, parseErr := strconv.Atoi(vars["id"])
	if parseErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	Ingredient, err := models.GetIngredient(id)
	if err != nil {
		helpers.StatusNotFound(w, err)
		return
	}
	helpers.StatusOk(w, Ingredient)
}

// GetAllIngredient Gets all Ingredient and sends the data as response
// to the requesting Ingredient
func GetAllIngredient(w http.ResponseWriter, r *http.Request) {
	query, err := helpers.GetQuery(r)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	ingredient, count, err := models.GetAllIngredient(&query)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	metaData := helpers.MetaData(count, len(ingredient), &query)

	response := IngredientResponse{
		Status:   http.StatusOK,
		MetaData: metaData,
		Data:     ingredient,
	}
	result, _ := json.Marshal(response)

	helpers.ResponseWriter(w, http.StatusOK, string(result))
}

//UpdateIngredient updates Ingredient's details
func UpdateIngredient(w http.ResponseWriter, r *http.Request) {
	Ingredient := models.Ingredient{
		Name:     r.FormValue("name"),
		RecipeID: r.FormValue("RecipeID"),
		Quantity: r.FormValue("quantity"),
		Unit:     r.FormValue("unit"),
	}
	decoder := json.NewDecoder(r.Body)
	decoderErr := decoder.Decode(&Ingredient)

	if decoderErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	_, err := models.UpdateIngredient(id, &Ingredient)
	if err != nil {
		helpers.BadRequest(w, errors.New(err.Error()))
		return
	}

	helpers.StatusOkMessage(w, "Ingredient updated")
}

// DeleteIngredient deletes an Ingredient
func DeleteIngredient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, parseErr := strconv.Atoi(vars["id"])
	if parseErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	_, err := models.DeleteIngredient(id)
	if err != nil {
		helpers.BadRequest(w, err)
		return
	}
	helpers.StatusOkMessage(w, "Ingredient deleted")
}
