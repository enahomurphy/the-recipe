package authentication

import (
	"encoding/json"
	"errors"
	"net/http"
	"recipe/helpers"
	"recipe/jwt"
	"recipe/models"
	"time"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Token to be sent to user for further authentication
type Token struct {
	Token string `json:"token"`
}

// Login controller
func Login(w http.ResponseWriter, r *http.Request) {
	auth := Auth{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	decoder := json.NewDecoder(r.Body)

	if decodeErr := decoder.Decode(&auth); decodeErr != nil {
		helpers.BadRequest(w, errors.New("Valid username and password required"))
		return
	}

	user, err := models.GetUserByUsername(auth.Username)
	if err != nil {
		helpers.BadRequest(w, err)
		return
	}

	payload := jwt.Payload{
		Sub:    user.UserName,
		Exp:    time.Now().Unix() + 2*24*60*60,
		Public: user,
	}

	secret := "secret"
	token := jwt.Encode(payload, secret)

	helpers.StatusOk(w, Token{Token: token})
}
