package authentication

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"recipe/helpers"
	"recipe/jwt"
	"recipe/models"
	"strings"
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

// AuthMiddleware routes
func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-Type", "application/json")
		auth := r.Header.Get("Authorization")
		token := strings.Split(auth, " ")
		if len(token) != 2 {
			helpers.Unauthorized(w, "Authentication failed: Invalid token")
			return
		}
		payload, err := jwt.Decode(token[1], "secret")
		fmt.Println(payload)
		if err != nil {
			helpers.Unauthorized(w, err.Error())
			return
		}
		ctx := context.WithValue(r.Context(), "user", payload)
		h.ServeHTTP(w, r.WithContext(ctx))
		// h.ServeHTTP(w, r)
	})
}
