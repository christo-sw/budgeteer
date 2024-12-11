package database

import (
	"fmt"
	"strconv"
	"strings"
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

func (b Balance) Amount() string {
	if b.AmountCents < 0 {
		return fmt.Sprintf("N$ (%s)",
			strings.ReplaceAll(strconv.FormatFloat(float64(b.AmountCents)/100, 'f', 2, 64),
				"-", ""))
	} else {
		return fmt.Sprintf("N$ %s",
			strconv.FormatFloat(float64(b.AmountCents)/100, 'f', 2, 64))
	}
}
