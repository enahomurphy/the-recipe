package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"recipe/helpers"
)

// Category data to be sent
// When request is made to the server
type Category struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

// GetAllCategory gets all user
func GetAllCategory() []Category {
	db := DB()
	categories := []Category{}

	rows, err := db.Query(`SELECT id, title, description created_at, updated_at FROM categories`)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		category := Category{}

		if err := rows.Scan(&category.ID, &category.Title, &category.Description, &category.CreatedAt, &category.UpdatedAt); err != nil {
			panic(err.Error())
		}
		categories = append(categories, category)
		fmt.Println(categories)
	}
	return categories
}

//GetCategory gets a single user
func GetCategory(id int) (Category, error) {
	db := DB()
	category := Category{}
	db.Close()

	rows := db.QueryRow("SELECT title, description, updated_at, created_ac FROM categories where id = ? ", id)

	err := rows.Scan(&category)

	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No category with that id found", err.Error())
		errMsg := fmt.Errorf("category with (id %d) does not exist", id)
		return category, errMsg
	case err != nil:
		log.Fatal("An error occurred", err.Error())
		errMsg := fmt.Errorf("an unknown error occurred %s", err.Error())
		return category, errMsg
	default:
		return category, nil
	}
}

//CreateCategory creates a new user
func CreateCategory(category *Category) (*Category, error) {
	db := DB()
	defer db.Close()
	var sql, values bytes.Buffer
	sql.WriteString("INSERT INTO categories (")
	values.WriteString("VALUES (")

	if category.Title != "" {
		sql.WriteString("title ")
		values.WriteString("'" + category.Title + "'")
	}
	if category.Description != "" {
		sql.WriteString(", description ")
		values.WriteString(", '" + category.Description + "'")
	}
	sql.WriteString(") ")
	sql.WriteString(values.String() + ")")
	query := sql.String()
	fmt.Println(query)
	_, err := db.Exec(sql.String())
	if err != nil {
		errMsg := fmt.Errorf("Error creating a user: %s?", err.Error())
		return category, errMsg
	}
	return category, nil
}

// DeleteCategory user from database
// NB: this method wipes all user details
// recipe ans ingredients inclusive
func DeleteCategory(id int) (bool, error) {
	db := DB()

	sql := `DELETE * FROM categories WHERE id = ?`
	row, err := db.Exec(sql, id)

	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	fmt.Println(row)

	return true, nil
}

// UpdateCategory updates category details base on the values sent
// takes the user id and user struct containing details to be update
func UpdateCategory(id int, category *Category) (bool, error) {
	db := DB()

	categoryValues := map[string]string{}

	_, err := GetCategory(id)
	if err != nil {
		return false, err
	}

	if category.Title != "" {
		categoryValues["title"] = category.Title
	}
	if category.Description != "" {
		categoryValues["description"] = category.Description
	}

	query := helpers.UpdateBuilder(categoryValues)
	_, UpdateErr := db.Exec(query, id)
	if UpdateErr != nil {
		return false, UpdateErr
	}

	return true, nil
}
