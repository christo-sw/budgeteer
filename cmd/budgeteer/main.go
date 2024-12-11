package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/christo-sw/budgeteer/internal/database"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	db, err := gorm.Open(sqlite.Open("budget.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&database.Transaction{})
	db.AutoMigrate(&database.Balance{})

	// Run through all categories, asking user to input balance for category if it does not exist
	reader := bufio.NewReader(os.Stdin)
	for _, category := range categories {
		var openingBalance database.Balance
		result := db.Where("category = ?", category).Last(&openingBalance)
		if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Fatalf("Failed to get opening balance: %v", result.Error)
		}

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			inputValid := false
			for !inputValid {
				fmt.Printf("Please enter an opening balance for '%s': ", category)
				text, err := reader.ReadString('\n')
				if err != nil {
					log.Fatalf("Failed to read input: %v", err)
				}

				amount, err := strconv.ParseFloat(strings.TrimSpace(text), 64)
				if err != nil {
					fmt.Printf("'%v' is not a valid amount, please try again\n", text)
					continue
				}

				openingBalance = database.Balance{
					Date:        time.Now(),
					Category:    category,
					AmountCents: int64(math.Floor(amount * 100)),
				}
				db.Create(&openingBalance)
				inputValid = true
			}
		}

		fmt.Printf("Opening balance for '%s' is %s\n", category, openingBalance.Amount())
	}
}
