package database

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Category string
}

type Balance struct {
	gorm.Model
	Date        time.Time
	Category    string
	AmountCents int64
}

type Transaction struct {
	gorm.Model
	ID          int64
	Date        time.Time
	AmountCents int64
	Description string
}
