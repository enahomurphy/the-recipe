package controllers

import (
	"encoding/json"
	"net/http"
	"recipe/helpers"
	"recipe/models"
	"strconv"

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

// // GetAllcategorys Gets all categorys and sends the data as response
// // to the requesting category
// func GetAllcategorys(w http.ResponseWriter, r *http.Request) {
// 	categorys, err := models.GetAllCategory()
// 	if err != nil {
// 		helpers.ServerError(w, err)
// 		return
// 	}
// 	response := categorysResponse{
// 		Status: http.StatusOK,
// 		Data:   categorys,
// 	}
// 	result, _ := json.Marshal(response)

// 	helpers.ResponseWriter(w, http.StatusOK, string(result))
// }

// //Updatecategory updates category's detail
// func Updatecategory(w http.ResponseWriter, r *http.Request) {
// 	category := models.category{
// 		FirstName:    r.FormValue("first_name"),
// 		LastName:     r.FormValue("last_name"),
// 		categoryName: r.FormValue("categoryname"),
// 		Email:        r.FormValue("email"),
// 		Password:     r.FormValue("password"),
// 	}

// 	decoder := json.NewDecoder(r.Body)

// 	decoderErr := decoder.Decode(&category)

// 	if decoderErr != nil {
// 		helpers.DecoderErrorResponse(w)
// 		return
// 	}
// 	vars := mux.Vars(r)
// 	id, _ := strconv.Atoi(vars["id"])

// 	_, err := models.Updatecategory(id, &category)
// 	fmt.Println(err.Error())
// 	if err != nil {
// 		helpers.BadRequest(w, err)
// 		return
// 	}
// 	helpers.StatusOk(w, category)
// }

// //Deletecategory deletes a category detail
// func Deletecategory(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)

// 	id, parseErr := strconv.Atoi(vars["id"])
// 	if parseErr != nil {
// 		helpers.DecoderErrorResponse(w)
// 		return
// 	}
// 	_, err := models.Deletecategory(id)
// 	if err != nil {
// 		helpers.BadRequest(w, err)
// 		return
// 	}
// 	helpers.ResponseWriter(w, http.StatusOK, "category deleted")
// }
