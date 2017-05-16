package models

import (
	"database/sql"
	"fmt"
	"log"
	"recipe/helpers"
)

// Ingredient data to be sent
// When request is made to the server
type Ingredient struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Quantity  string `json:"quantity,omitempty"`
	RecipeID  string `json:"recipeID,omitempty"`
	Unit      string `json:"unit,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// GetAllIngredient gets all ingredient
func GetAllIngredient() []Ingredient {
	db := DB()
	ingredients := []Ingredient{}

	rows, err := db.Query(`SELECT id, name, quantity, recipeID, unit created_at, updated_at FROM recipes`)

	defer db.Close()

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		ingredient := Ingredient{}

		if err := rows.Scan(&ingredient); err != nil {
			panic(err.Error())
		}
		ingredients = append(ingredients, ingredient)
		fmt.Println(ingredient)
	}
	return ingredients
}

// GetIngredient gets a single ingredient
func GetIngredient(id int) (Ingredient, error) {
	db := DB()
	ingredient := Ingredient{}
	db.Close()

	err := db.QueryRow(`SELECT name, quantity, recipeID, unit, created_at, 
		updated_at FROM ingredients where id = ? `, id).Scan(&ingredient)

	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No rows with that id found", err.Error())
		errMsg := fmt.Errorf("ingredient with (id %d) does not exist", id)
		return ingredient, errMsg
	case err != nil:
		log.Fatal("An error occurred", err.Error())
		errMsg := fmt.Errorf("an unknown error occurred %s", err.Error())
		return ingredient, errMsg
	default:
		return ingredient, nil
	}
}

// CreateIngredient creates a new recipe
func CreateIngredient() {
	db := DB()
	ingredient := Ingredient{}
	sql := `INSERT INTO recipes (name, quantity, recipeID, unit) VALUES(?, ?, ?, ?)`
	row, err := db.Exec(sql, &ingredient)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(row)
}

// DeleteIngredient ingredient from database
// NB: this method wipes all recipe details
// recipe ans ingredients inclusive
func DeleteIngredient(id int) (bool, error) {
	db := DB()

	sql := `DELETE * FROM ingredients WHERE id = ?`
	row, err := db.Exec(sql, id)

	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	fmt.Println(row)

	return true, nil
}

// UpdateIngredient updates recipe details base on the values sent
// takes the recipe id and recipe struct containing
// details to be update
func UpdateIngredient(id int, ingredient *Ingredient) (bool, error) {
	db := DB()

	ingredientValues := map[string]string{}

	if ingredient.Name != "" {
		ingredientValues["name"] = ingredient.Name
	}
	if ingredient.RecipeID != "" {
		ingredientValues["recipeID"] = ingredient.RecipeID
	}
	if ingredient.Unit != "" {
		ingredientValues["unit"] = ingredient.Unit
	}
	if ingredient.Quantity != "" {
		ingredientValues["quantity"] = ingredient.Quantity
	}

	query := helpers.UpdateBuilder(ingredientValues)

	println(query)
	rows, err := db.Exec(query, id)
	if err != nil {
		panic(err.Error())
	} else {
		println(rows)
	}

	return true, nil
}
