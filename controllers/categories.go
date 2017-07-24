package controllers

import (
	"encoding/json"
	"recipe/helpers"
	"recipe/models"
	"net/http"
	"strconv"

	"errors"

	"github.com/gorilla/mux"
)

// CategoryResponse is
type CategoryResponse struct {
	Status   int               `json:"status"`
	MetaData interface{}       `json:"meta_data,omitempty"`
	Data     []models.Category `json:"data"`
}

// CreateCategory creates a new category
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	category := models.Category{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}
	decoder := json.NewDecoder(r.Body)
	decoderErr := decoder.Decode(&category)

	if decoderErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	if category.Title == "" {
		errMsg := models.Category{
			Title: "Title  is required",
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
	_, dbErr := models.CreateCategory(&category)

	if dbErr != nil {
		helpers.ServerError(w, errors.New(dbErr.Error()))
		return
	}
	helpers.StatusOk(w, category)
}

// GetCategory Gets all category and sends the data as response
// to the requesting category
func GetCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, parseErr := strconv.Atoi(vars["id"])
	if parseErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	category, err := models.GetCategory(id)
	if err != nil {
		helpers.StatusNotFound(w, err)
		return
	}
	helpers.StatusOk(w, category)
}

// GetAllcategory Gets all category and sends the data as response
// to the requesting category
func GetAllcategory(w http.ResponseWriter, r *http.Request) {
	query, err := helpers.GetQuery(r)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	category, count, err := models.GetAllCategory(query)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	metaData := helpers.MetaData(count, len(category), &query)

	response := CategoryResponse{
		Status:   http.StatusOK,
		Data:     category,
		MetaData: metaData,
	}
	result, _ := json.Marshal(response)

	helpers.ResponseWriter(w, http.StatusOK, string(result))
}

// UpdateCategory updates category's detail
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	category := models.Category{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}

	decoder := json.NewDecoder(r.Body)
	decoderErr := decoder.Decode(&category)

	if decoderErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	_, err := models.UpdateCategory(id, &category)
	if err != nil {
		helpers.BadRequest(w, errors.New(err.Error()))
		return
	}
	helpers.StatusOkMessage(w, "category updated")
}

// DeleteCategory deletes a category detail
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, parseErr := strconv.Atoi(vars["id"])
	if parseErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	_, err := models.DeleteCategory(id)
	if err != nil {
		helpers.BadRequest(w, err)
		return
	}
	helpers.StatusOkMessage(w, "category deleted")
}
