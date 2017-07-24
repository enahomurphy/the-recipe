package models

import (
	"database/sql"
	"fmt"
	"recipe/helpers"
	"strconv"
	"strings"
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
func GetAllIngredient(query *helpers.Query) ([]Ingredient, int, error) {
	db := DB()

	defer db.Close()
	ingredients := []Ingredient{}

	limit := strconv.Itoa(query.Limit)
	offset := strconv.Itoa(query.Offset)
	var dbQuery string
	if query.Q == "" {
		dbQuery = `SELECT id, name, quantity, recipeID, unit created_at, updated_at FROM ingredients
			LIMIT ` + limit + ` OFFSET ` + offset
	} else {
		dbQuery = `SELECT id, name, quantity, recipeID, unit created_at, updated_at FROM ingredients
			WHERE title ILIKE %` + query.Q + `% ` +
			` LIMIT = ` + limit + ` OFFSET = ` + offset
	}
	rows, err := db.Query(dbQuery)
	if err != nil {
		errMsg := fmt.Errorf("an unknown error occurred %s", err.Error())
		return nil, 0, errMsg
	}

	for rows.Next() {
		ingredient := Ingredient{}
		if rows.Scan(&ingredient.ID, &ingredient.Name, &ingredient.Quantity,
			&ingredient.RecipeID, &ingredient.Unit, &ingredient.CreatedAt, &ingredient.UpdatedAt); err != nil {
			errMsg := fmt.Errorf("an unknown error occurred %s", err.Error())
			return nil, 0, errMsg
		}
		ingredients = append(ingredients, ingredient)
	}
	count, countErr := helpers.GetCount(db, "ingredients")
	if countErr != nil {
		errMsg := fmt.Errorf("an unknown error occurred %s", countErr.Error())
		return nil, 0, errMsg
	}
	return ingredients, count, nil
}

// GetIngredient gets a single ingredient
func GetIngredient(id int) (Ingredient, error) {
	db := DB()
	defer db.Close()

	ingredient := Ingredient{}
	err := db.QueryRow(`SELECT id, name, quantity, recipeID, unit, created_at, 
		updated_at FROM ingredients where id = ? `, id).
		Scan(&ingredient.ID, &ingredient.Name, &ingredient.Quantity, &ingredient.RecipeID, &ingredient.Unit,
			&ingredient.CreatedAt, &ingredient.UpdatedAt)

	switch {
	case err == sql.ErrNoRows:
		errMsg := fmt.Errorf("ingredient with (id %d) does not exist", id)
		return ingredient, errMsg
	case err != nil:
		errMsg := fmt.Errorf("an unknown error occurred %s", err.Error())
		return ingredient, errMsg
	default:
		return ingredient, nil
	}
}

// CreateIngredient creates a new recipe
func CreateIngredient(ingredient *Ingredient) (*Ingredient, error) {
	db := DB()
	sql := `INSERT INTO ingredients (name, quantity, recipeID, unit) VALUES(?, ?, ?, ?)`
	_, err := db.Exec(sql, &ingredient.Name, &ingredient.Quantity, &ingredient.RecipeID, &ingredient.Unit)
	switch {
	case err != nil && strings.Contains(err.Error(), "recipeID"):
		errMsg := fmt.Errorf("Error creating a ingredients: %s", "recipeID does not exist")
		return nil, errMsg
	case err != nil:
		errMsg := fmt.Errorf("Error creating a ingredients: %s", err.Error())
		return nil, errMsg
	default:
		return ingredient, nil
	}
}

// DeleteIngredient ingredient from database
// NB: this method wipes all recipe details
// recipe ans ingredients inclusive
func DeleteIngredient(id int) (bool, error) {
	db := DB()

	defer db.Close()

	query := `DELETE FROM ingredients WHERE id = ?`
	row, err := db.Exec(query, id)
	RowsAffected, _ := row.RowsAffected()
	switch {
	case RowsAffected == 0:
		errMsg := fmt.Errorf("ingredient with (id %d) does not exist", id)
		return false, errMsg
	case err != nil:
		errMsg := fmt.Errorf("an unknown error occurred %s", err.Error())
		return false, errMsg
	default:
		return true, nil
	}
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

	query := helpers.UpdateBuilder(ingredientValues, "INGREDIENTS")
	_, UpdateErr := db.Exec(query, id)
	if UpdateErr != nil {
		return false, UpdateErr
	}

	return true, nil
}
