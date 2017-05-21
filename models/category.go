package models

import (
	"database/sql"
	"fmt"
	"recipe/helpers"
	"strconv"
)

// Category data to be sent
// When request is made to the server
type Category struct {
	MetaData    interface{} `json:"meta_data,omitempty"`
	ID          int         `json:"id,omitempty"`
	Title       string      `json:"title,omitempty"`
	Description string      `json:"description"`
	CreatedAt   string      `json:"created_at,omitempty"`
	UpdatedAt   string      `json:"updated_at,omitempty"`
}

// GetAllCategory gets all category
func GetAllCategory(query helpers.Query) ([]Category, int, error) {
	db := DB()

	defer db.Close()

	limit := strconv.Itoa(query.Limit)
	offset := strconv.Itoa(query.Offset)
	var dbQuery string
	categories := []Category{}
	if query.Q == "" {
		dbQuery = `SELECT id, title, description, created_at, updated_at FROM categories
			LIMIT ` + limit + ` OFFSET ` + offset
	} else {
		dbQuery = `SELECT id, title, description, created_at, updated_at FROM categories 
			WHERE title ILIKE %` + query.Q + `% ` +
			` LIMIT = ` + limit + ` OFFSET = ` + offset
	}
	rows, err := db.Query(dbQuery)
	if err != nil {
		errMsg := fmt.Errorf("an unknown error occurred %s", err.Error())
		return nil, 0, errMsg
	}

	for rows.Next() {
		category := Category{}

		if err := rows.Scan(&category.ID, &category.Title, &category.Description, &category.CreatedAt, &category.UpdatedAt); err != nil {
			errMsg := fmt.Errorf("an unknown error occurred %s", err.Error())
			return nil, 0, errMsg
		}
		categories = append(categories, category)
	}
	count, countErr := helpers.GetCount(db, "categories")
	if countErr != nil {
		errMsg := fmt.Errorf("an unknown error occurred %s", countErr.Error())
		return nil, 0, errMsg
	}
	return categories, count, nil
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
		errMsg := fmt.Errorf("category with (id %d) does not exist", id)
		return category, errMsg
	case err != nil:
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
func UpdateCategory(id int, category *Category) (bool, error) {
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
	_, UpdateErr := db.Exec(query, id)
	if UpdateErr != nil {
		return false, UpdateErr
	}
	return true, nil
}
