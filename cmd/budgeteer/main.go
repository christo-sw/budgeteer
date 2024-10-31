package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open connection to SQLite database
	db, err := sqlx.Open("sqlite3", "./budget.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Test the connection
	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("Cannot connect to database: %v\n", err))
	}

	fmt.Println("Successfully connected to DB")
}
