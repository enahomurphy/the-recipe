package models

import (
	"database/sql"
	"fmt"
	"log"
)

// User data to be sent
// When request is made to the server
type User struct {
	ID         string `json:"id,omitempty"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	UserName   string `json:"username"`
	Email      string `json:"email"`
	ProfilePic string `json:"profile_pic"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

// GetAll gets all user
func GetAll() []User {
	db := DB()
	users := []User{}

	rows, err := db.Query("SELECT id, first_name, last_name, username, email, profile_pic, created_at, updated_at FROM users")

	defer db.Close()

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		user := User{}

		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.UserName,
			&user.Email, &user.ProfilePic, &user.CreatedAt, &user.UpdatedAt); err != nil {
			panic(err.Error())
		}
		users = append(users, user)
		fmt.Println(user)
	}
	return users
}

// Get gets a single user
func Get(id int) (User, error) {
	db := DB()
	user := User{}
	db.Close()

	err := db.QueryRow("SELECT first_name, last_name, username, email, profile_pic FROM users where id = ? ", id).Scan(&user.FirstName, &user.Email, &user.LastName, &user.UserName, &user.ProfilePic)

	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No rows with that id found", err.Error())
		errMsg := fmt.Errorf("user with (id %d) does not exist", id)
		return user, errMsg
	case err != nil:
		log.Fatal("An error occurred", err.Error())
		errMsg := fmt.Errorf("an unknown error occurred %s", err.Error())
		return user, errMsg
	default:
		return user, nil
	}
}

// CreateUser creates a new user
func (u User) CreateUser(db *sql.DB) {
	sql := `INSERT INTO users(first_name, last_name, email, username,
				profile_pic ) VALUES(?, ?, ?, ?, ?)`
	row, err := db.Exec(sql, u.FirstName, u.LastName, u.Email, u.UserName, u.ProfilePic)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(row)
}

// Delete user from database
func Delete(id int) (bool, error) {
	db := DB()

	sql := `DELETE * FROM users WHERE id = ?`
	row, err := db.Exec(sql, id)

	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	fmt.Println(row)

	return true, nil
}

// Update user details
func Update(id int) (User, error) {
	db := DB()
	user := User{
		FirstName: "Jessy",
		LastName: "enaho",
		Email: "jessyba@gmail.com",
		UserName: "jboys"
	}

	for _, value := range user {
		
	}

	user, err := Get(id)
	if err != nil {
		return user, err
	} else {
		sql := `UPDATE users
				SET column1 = value1, column2 = value2, ...
				WHERE condition;`
	}

	return true, nil
}

func handleUpdate(*User) {

}
