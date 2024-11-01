package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	DB *sqlx.DB
}

type Category struct {
	Category string `db:"category"`
}

type Balance struct {
	Date        time.Time `db:"date"`
	Category    Category  `db:"category"`
	AmountCents int64     `db:"amount_cents"`
}

type Transaction struct {
	ID          int64     `db:"id"`
	Date        time.Time `db:"date"`
	AmountCents int64     `db:"amount_cents"`
	Description string    `db:"description"`
}

func (s Store) GetLatestBalance(category Category) (*Balance, error) {
	balance := Balance{}

	err := s.DB.Get(&balance, "SELECT * FROM balances WHERE category=$1 ORDER BY date DESC LIMIT 1",
		category.Category)
	if err != nil {
		return nil, err
	}

	return &balance, nil
}

func (s Store) CreateBalance(category Category, amount int64) (*Balance, error) {
	tx, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec("INSERT INTO balances (date, category, amount_cents) VALUES ($1, $2, $3)",
		time.Now(), category.Category, amount)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	balance, err := s.GetLatestBalance(category)
	if err != nil {
		return nil, err
	}

	return balance, nil
}

func (s Store) CreateCategory() {}
