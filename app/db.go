package app

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type DB struct {
	DB *sql.DB
}

const (
	DB_USER     = "anya"
	DB_PASSWORD = "anya"
	DB_NAME     = "order_api"
)

func NewDB() *DB {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		DB: db,
	}
}
