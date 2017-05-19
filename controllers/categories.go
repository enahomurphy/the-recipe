package controllers

import (
	"encoding/json"
	"net/http"
	"recipe/helpers"
	"recipe/models"
	"strconv"

	"fmt"

	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// CategoryResponse is
type CategoryResponse struct {
	Status int               `json:"status"`
	Data   []models.Category `json:"data"`
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
		helpers.ServerError(w, dbErr)
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
	category, err := models.GetAllCategory()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	response := CategoryResponse{
		Status: http.StatusOK,
		Data:   category,
	}
	result, _ := json.Marshal(response)

	helpers.ResponseWriter(w, http.StatusOK, string(result))
}

//UpdateCategory updates category's detail
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

	fmt.Println(id)

	_, err := models.UpdateCategoryById(id, &category)
	// fmt.Println(err.Error())
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
