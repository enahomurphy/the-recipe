package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DB() *sql.DB {
	db, err := sql.Open("mysql", "root@/go_recipes")
	if err != nil {
		fmt.Println(err.Error())
	}

	pingErr := db.Ping()

	if pingErr != nil {
		fmt.Println(pingErr.Error())
	}
	return db
}
