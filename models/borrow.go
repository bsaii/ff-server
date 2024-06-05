package models

import (
	"gorm.io/gorm"
)

type Borrow struct {
	gorm.Model
	Address    string
	Amount     int
	Collateral int
}
