package main

import (
	"fmt"
	"log"

	"github.com/christo-sw/budgeteer/internal/database"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	categories = []string{
		"Tithe",
		"Tax",
		"Medical",
		"Pension",
		"Internet",
		"Clothing",
		"Cellphone",
		"Car Insurance",
		"Car Licence",
		"Rent",
		"Obsidian",
		"1Password",
		"uSchool",
		"Petrol/Parking",
		"Eating Out",
		"Offering",
		"Gym",
		"Banking Fees",
		"Travel",
		"Christmas Gifts",
		"Birthday Gifts",
		"Vacation",
		"Office Equipment",
		"XTO Bespoke",
		"Sundry",
	}
)

func main() {
	db, err := gorm.Open(sqlite.Open("budget.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&database.Transaction{})
	db.AutoMigrate(&database.Balance{})

	var balances []database.Balance
	db.Find(&balances)
	fmt.Printf("Balance: %+v\n", balances)
}
