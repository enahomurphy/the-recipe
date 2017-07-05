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

type usersResponse struct {
	Status   int           `json:"status"`
	MetaData interface{}   `json:"meta_data,omitempty"`
	Data     []models.User `json:"data"`
}

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		UserName:  r.FormValue("username"),
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
	}

	decoder := json.NewDecoder(r.Body)

	decoderErr := decoder.Decode(&user)

	if decoderErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	if user.FirstName == "" || user.LastName == "" ||
		(user.UserName == "" || len(user.UserName) < 3) || !helpers.IsValidEmail(user.Email) ||
		(user.Password == "" || len(user.Password) < 6) {
		errMsg := models.User{
			FirstName: "first name is required",
			LastName:  "Last  name is required",
			UserName:  "Username is required and should be more that 3 characters",
			Email:     "Email is required and should be valid",
			Password:  "Password is required and should be more than 6 characters",
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
	// hash user password and save
	user.Password = helpers.HashPassword(user.Password)
	_, dbErr := models.CreateUser(&user)

	if dbErr != nil {
		helpers.ServerError(w, dbErr)
		return
	}
	helpers.StatusOk(w, user)
}

// GetUser Gets all users and sends the data as response
// to the requesting user
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, parseErr := strconv.Atoi(vars["id"])
	if parseErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	user, err := models.GetUser(id)
	if err != nil {
		helpers.StatusNotFound(w, err)
		return
	}
	helpers.StatusOk(w, user)
}

// GetAllUsers Gets all users and sends the data as response
// to the requesting user
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	query, err := helpers.GetQuery(r)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	users, count, err := models.GetAllUser(query)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	metaData := helpers.MetaData(count, len(users), &query)

	response := usersResponse{
		Status:   http.StatusOK,
		MetaData: metaData,
		Data:     users,
	}
	result, _ := json.Marshal(response)

	helpers.ResponseWriter(w, http.StatusOK, string(result))
}

//UpdateUser updates user's detail
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		UserName:  r.FormValue("username"),
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
	}

	decoder := json.NewDecoder(r.Body)

	decoderErr := decoder.Decode(&user)

	if decoderErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	_, err := models.UpdateUserByID(id, &user)
	if err != nil {
		helpers.BadRequest(w, errors.New(err.Error()))
		return
	}
	helpers.StatusOkMessage(w, "User with "+vars["id"]+" updated")
}

//DeleteUser deletes a user detail
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, parseErr := strconv.Atoi(vars["id"])
	if parseErr != nil {
		helpers.DecoderErrorResponse(w)
		return
	}
	_, err := models.DeleteUser(id)
	if err != nil {
		helpers.BadRequest(w, err)
		return
	}
	helpers.StatusOkMessage(w, "user with id"+vars["id"]+"deleted")
}
