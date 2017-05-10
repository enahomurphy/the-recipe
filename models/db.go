package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DB() *sql.DB {
	db, err := sql.Open("mysql", "root@/go_recipes")
	if err != nil {
		panic(err.Error())
	}

	pingErr := db.Ping()

	if pingErr != nil {
		panic(pingErr.Error())
	}

	return db
}
