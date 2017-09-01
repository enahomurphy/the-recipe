package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//DB connection instance to mysql
func DB() *sql.DB {
	db, err := sql.Open("mysql", "recipe:password@tcp(mysql:3306)/recipe")
	if err != nil {
		fmt.Println(err.Error())
	}

	pingErr := db.Ping()

	if pingErr != nil {
		fmt.Println(pingErr.Error())
	}
	return db
}
