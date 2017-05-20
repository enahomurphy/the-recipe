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
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

// GetAllCategory gets all category
func GetAllCategory() ([]Category, error) {
	db := DB()

	defer db.Close()

	categories := []Category{}

	rows, err := db.Query(`SELECT id, title, description, created_at, updated_at FROM categories`)

	if err != nil {
		errMsg := fmt.Errorf("an unknown error occurred %s", err.Error())
		return nil, errMsg

	}

	for rows.Next() {
		category := Category{}

		if err := rows.Scan(&category.ID, &category.Title, &category.Description, &category.CreatedAt, &category.UpdatedAt); err != nil {
			errMsg := fmt.Errorf("an unknown error occurred %s", err.Error())
			return nil, errMsg
		}
		categories = append(categories, category)
		fmt.Println(categories)
	}
	return categories, nil
}

//GetCategory gets a single category
func GetCategory(id int) (Category, error) {
	db := DB()
	defer db.Close()

	category := Category{}

	row := db.QueryRow("SELECT id, title, description, updated_at, created_at FROM categories where id = ? ", id)

	err := row.Scan(&category.ID, &category.Title, &category.Description, &category.UpdatedAt, &category.CreatedAt)

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

//CreateCategory creates a new category
func CreateCategory(category *Category) (*Category, error) {
	db := DB()
	defer db.Close()
	sql := "INSERT INTO categories (title, description) VALUES (?, ?)"
	_, err := db.Exec(sql, &category.Title, &category.Description)
	if err != nil {
		errMsg := fmt.Errorf("Error creating a category: %s?", err.Error())
		return category, errMsg
	}
	return category, nil
}

// DeleteCategory category from database
// NB: this method wipes all category details
// recipe ans ingredients inclusive
func DeleteCategory(id int) (bool, error) {
	db := DB()

	sql := `DELETE FROM categories WHERE id = ?`
	_, err := db.Exec(sql, id)

	defer db.Close()

	if err != nil {
		errMsg := fmt.Errorf("Error Deletin a category: %s?", err.Error())
		return false, errMsg
	}
	return true, nil
}

// UpdateCategory updates category details base on the values sent
// takes the category id and category struct containing details to be update
func UpdateCategoryById(id int, category *Category) (bool, error) {
	db := DB()

	defer db.Close()

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

	query := helpers.UpdateBuilder(categoryValues, "CATEGORIES")
	fmt.Println(query)
	_, UpdateErr := db.Exec(query, id)
	if UpdateErr != nil {
		return false, UpdateErr
	}

	return true, nil
}
