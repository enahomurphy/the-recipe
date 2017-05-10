package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


sql.Open("mysql", )

// Data data to be sent
// When request is made to the server
type Data struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Create creates a
func Create(w http.ResponseWriter, r *http.Request) {

	user := Data{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		UserName:  r.FormValue("username"),
		Email:     r.FormValue("email"),
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}

	resData, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprint(w, string(resData))
}
