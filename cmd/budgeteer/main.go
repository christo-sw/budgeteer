package main

import (
	"fmt"
	"os"

	"github.com/christo-sw/budgeteer/internal/database"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var (
	categories = []database.Category{
		{Category: "Tithe"},
		{Category: "Tax"},
		{Category: "Medical"},
		{Category: "Pension"},
		{Category: "Internet"},
		{Category: "Clothing"},
		{Category: "Cellphone"},
		{Category: "Car Insurance"},
		{Category: "Car Licence"},
		{Category: "Rent"},
		{Category: "Obsidian"},
		{Category: "1Password"},
		{Category: "uSchool"},
		{Category: "Petrol/Parking"},
		{Category: "Eating Out"},
		{Category: "Offering"},
		{Category: "Gym"},
		{Category: "Banking Fees"},
		{Category: "Travel"},
		{Category: "Christmas Gifts"},
		{Category: "Birthday Gifts"},
		{Category: "Vacation"},
		{Category: "Office Equipment"},
		{Category: "XTO Bespoke"},
		{Category: "Sundry"},
	}
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

	schema, err := os.ReadFile("internal/database/schema.sql")
	if err != nil {
		panic(err)
	}

	db.MustExec(string(schema))

	fmt.Println("Successfully connected to DB")

	store := database.Store{DB: db}

	balance, err := store.CreateBalance(database.Category{Category: "Tithe"}, 300000)
	if err != nil {
		panic(err)
	}

	balance, err = store.GetLatestBalance(database.Category{Category: "Tithe"})
	if err != nil {
		panic(err)
	}

	fmt.Println("Latest balances:")
	fmt.Printf("%+v\n", balance)
}
