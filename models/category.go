package models

import (
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
func (u User) CreateCategory(category Category) {
	db := DB()
	sql := `INSERT INTO categories (title, description) VALUES(?, ?)`
	row, err := db.Exec(sql, category.ID, category.Title, category.Description)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(row)
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

	if category.Title != "" {
		categoryValues["title"] = category.Title
	}
	if category.Description != "" {
		categoryValues["description"] = category.Description
	}

	query := helpers.UpdateBuilder(categoryValues)

	println(query)
	rows, err := db.Exec(query, id)
	if err != nil {
		panic(err.Error())
	} else {
		println(rows)
	}

	return true, nil
}
