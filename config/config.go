package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	dsn := "user=postgres database=postgres password=new_password sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Error opening database", err)
		panic(err)
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Error pinging database", err)
		panic(err)
	}
	fmt.Println("Successfully connected to database")
	DB = db
}
