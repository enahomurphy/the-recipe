package models

import (
	"database/sql"
	"recipe/helpers"
)

// CreateTables needed to store
// recipes and user details
func CreateTables(*sql.DB) {
	_, userErr := DB().Exec(
		`CREATE TABLE users (
			id int(11) AUTO_INCREMENT NOT NULL,
			first_name varchar(255),
			last_name varchar(255),
			email varchar(255),
			username varchar(255),
			profile_pic text,
			PRIMARY KEY (id)
		)`)

	helpers.PrintErr(userErr)

	_, categoryErr := DB().Exec(
		`CREATE TABLE categories (
			id int(11) AUTO_INCREMENT NOT NULL,
			title varchar(255),
			description text(300),
			PRIMARY KEY (id)
		)`)

	helpers.PrintErr(categoryErr)

	_, ingredientErr := DB().Exec(
		`CREATE TABLE ingredients (
			id int(11) NOT NULL AUTO_INCREMENT,
			name varchar(255),
			quantity varchar(50),
			unit varchar(50),
			PRIMARY KEY (id)
		)`)

	helpers.PrintErr(ingredientErr)

	_, recipeError := DB().Exec(
		`CREATE TABLE recipes (
			id int NOT NUll AUTO_INCREMENT,
			name varchar(255),
			description text(300),
			PRIMARY KEY (id)
		)`)

	helpers.PrintErr(recipeError)

	_, categoryForeign := DB().Exec(
		`
		ALTER TABLE recipes
			ADD FOREIGN KEY (categoryID) REFERENCES categories(id);
		`)
	helpers.PrintErr(categoryForeign)

	_, ingredientForeign := DB().Exec(
		`
		ALTER TABLE recipes
			ADD FOREIGN KEY (ingredientID) REFERENCES ingredients(id);
		`)
	helpers.PrintErr(ingredientForeign)
}
