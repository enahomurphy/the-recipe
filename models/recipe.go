package models

import (
	"database/sql"
	"fmt"
	"log"
	"recipe/helpers"
	"strings"
)

// Recipe data to be sent
// When request is made to the server
type Recipe struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	UserID      string `json:"userID,omitempty"`
	CategoryID  string `json:"catergoryID,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

// GetAllRecipe gets all recipe
func GetAllRecipe() []Recipe {
	db := DB()
	recipes := []Recipe{}

	rows, err := db.Query(`SELECT id, name, userID, categoryID, description created_at, updated_at FROM recipes`)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		recipe := Recipe{}

		if err := rows.Scan(&recipe.ID, &recipe.Name, &recipe.UserID, &recipe.CategoryID,
			&recipe.Description, &recipe.CreatedAt, &recipe.UpdatedAt); err != nil {
			panic(err.Error())
		}
		recipes = append(recipes, recipe)
		fmt.Println(recipe)
	}
	return recipes
}

// GetRecipe gets a single recipe
func GetRecipe(id int) (Recipe, error) {
	db := DB()
	recipe := Recipe{}
	db.Close()

	err := db.QueryRow(`SELECT name, userID, categoryID, description, created_at, 
		updated_at FROM recipes where id = ? `, id).Scan(&recipe)

	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No rows with that id found", err.Error())
		errMsg := fmt.Errorf("recipe with (id %d) does not exist", id)
		return recipe, errMsg
	case err != nil:
		log.Fatal("An error occurred", err.Error())
		errMsg := fmt.Errorf("an unknown error occurred %s", err.Error())
		return recipe, errMsg
	default:
		return recipe, nil
	}
}

// CreateRecipe creates a new recipe
func CreateRecipe(recipe *Recipe) (*Recipe, error) {
	db := DB()
	defer db.Close()
	fmt.Println(recipe)
	sql := `INSERT INTO recipes (name, userID, categoryID, description) VALUES(?, ?, ?, ?)`
	_, err := db.Exec(sql, &recipe.Name, &recipe.UserID, &recipe.CategoryID, &recipe.Description)

	switch {
	case err != nil && strings.Contains(err.Error(), "categoryID"):
		errMsg := fmt.Errorf("Error creating a recipe: %s", "category does not exist")
		return recipe, errMsg
	case err != nil && strings.Contains(err.Error(), "userID"):
		errMsg := fmt.Errorf("Error creating a recipe: %s", "user does not exist")
		return recipe, errMsg
	case err != nil:
		errMsg := fmt.Errorf("Error creating a recipe: %s", err.Error())
		return recipe, errMsg
	default:
		return recipe, nil
	}
}

// DeleteRecipe recipe from database
// NB: this method wipes all recipe details
// recipe ans ingredients inclusive
func DeleteRecipe(id int) (bool, error) {
	db := DB()

	sql := `DELETE * FROM recipes WHERE id = ?`
	row, err := db.Exec(sql, id)

	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	fmt.Println(row)

	return true, nil
}

// UpdateRecipe updates recipe details base on the values sent
// takes the recipe id and recipe struct containing
// details to be update
func UpdateRecipe(id int, recipe *Recipe) (bool, error) {
	db := DB()

	recipeValues := map[string]string{}

	if recipe.Name != "" {
		recipeValues["name"] = recipe.Name
	}
	if recipe.UserID != "" {
		recipeValues["userID"] = recipe.UserID
	}
	if recipe.CategoryID != "" {
		recipeValues["categoryID"] = recipe.CategoryID
	}
	if recipe.Description != "" {
		recipeValues["description"] = recipe.Description
	}

	query := helpers.UpdateBuilder(recipeValues, "RECIPES")

	println(query)
	rows, err := db.Exec(query, id)
	if err != nil {
		panic(err.Error())
	} else {
		println(rows)
	}

	return true, nil
}
