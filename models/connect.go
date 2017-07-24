package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func DB() *sql.DB {
	db, err := sql.Open("mysql", "recipe:password@tcp(mysql)/recipe?tls=skip-verify&autocommit=true")
	if err != nil {
		fmt.Println(err.Error())
	}

	pingErr := db.Ping()

	if pingErr != nil {
		fmt.Println(pingErr.Error())
	}
	return db
}
