package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"recipe/models"

	"strconv"

	"recipe/helpers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Error message
type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Data data to be sent
// When request is made to the server
type Data struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	UserName  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

type usersResponse struct {
	Status int           `json:"status"`
	Data   []models.User `json:"data"`
}

// Create creates a
func Create(w http.ResponseWriter, r *http.Request) {
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
		response := helpers.RespondMessage("Error parsing body set content-type to application/json", http.StatusInternalServerError)
		fmt.Println(response)
		w.Header().Set("Content-type", "Application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, response)
	} else {
		fmt.Println(helpers.IsValidEmail(user.Email))
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
			w.Header().Set("Content-type", "Application/json")
			response := helpers.RespondMessages(errMsg, http.StatusBadRequest)
			fmt.Fprint(w, response)
		} else {
			_, dbErr := models.CreateUser(&user)

			if dbErr != nil {
				response := helpers.RespondMessage(dbErr.Error(), http.StatusBadRequest)
				w.Header().Set("Content-type", "Application/json")
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, response)
			} else {
				w.Header().Set("Content-type", "Application/json")
				w.WriteHeader(http.StatusOK)
				response := helpers.RespondWithData("user created", http.StatusOK, user)
				fmt.Println(response)
				fmt.Fprint(w, response)
			}
		}
	}
}

// GetUser Gets all users and sends the data as response
// to the requesting user
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, parseErr := strconv.Atoi(vars["id"])
	if parseErr != nil {
		w.Header().Set("Content-type", "Application/json")
		errMsg := helpers.RespondMessage(parseErr.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, errMsg)
	} else {
		user, err := models.GetUser(id)
		if err != nil {
			w.Header().Set("Content-type", "Application/json")
			errMsg := helpers.RespondMessage(err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, errMsg)
		} else {
			w.Header().Set("Content-type", "Application/json")
			w.WriteHeader(http.StatusOK)
			response := helpers.RespondWithData("user deleted", http.StatusOK, user)
			fmt.Fprintf(w, response)
		}
	}

	user, err := models.GetUser(id)

	if err != nil {
		w.Header().Set("Content-type", "Application/json")
		errMsg := helpers.RespondMessage(err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, errMsg)
	} else {
		response := helpers.RespondWithData("", http.StatusOK, user)
		fmt.Println(w, response)
	}

}

// GetAllUsers Gets all users and sends the data as response
// to the requesting user
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := models.GetAllUser()
	fmt.Println(users)
	response := usersResponse{
		Status: http.StatusOK,
		Data:   users,
	}
	result, _ := json.Marshal(response)

	w.Header().Set("Content-type", "Application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(result))
}

//Updates a  user detail in the database
func Update(w http.ResponseWriter, r *http.Request) {
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
		response := helpers.RespondMessage("Error parsing body set content-type to application/json", http.StatusInternalServerError)
		fmt.Println(response)
		w.Header().Set("Content-type", "Application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, response)
	} else {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		_, err := models.UpdateUser(id, &user)

		if err != nil {
			w.Header().Set("Content-type", "Application/json")
			errMsg := helpers.RespondMessage(err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, errMsg)
		} else {
			w.Header().Set("Content-type", "Application/json")
			w.WriteHeader(http.StatusOK)
			response := helpers.RespondWithData("user updated", http.StatusOK, user)
			fmt.Fprintf(w, response)
		}
	}
}

//Deletes a  user detail in the database
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, parseErr := strconv.Atoi(vars["id"])
	if parseErr != nil {
		w.Header().Set("Content-type", "Application/json")
		errMsg := helpers.RespondMessage(parseErr.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, errMsg)
	} else {
		_, err := models.DeleteUser(id)
		if err != nil {
			w.Header().Set("Content-type", "Application/json")
			errMsg := helpers.RespondMessage(err.Error(), http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, errMsg)
		} else {
			w.Header().Set("Content-type", "Application/json")
			w.WriteHeader(http.StatusOK)
			response := helpers.RespondMessages("user deleted", http.StatusOK)
			fmt.Fprintf(w, response)
		}
	}
}
