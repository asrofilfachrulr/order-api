package app

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "anya:anya@tcp(localhost:3306)/order_api")

	if err != nil {
		log.Fatal(err)
	}

	return db

}
