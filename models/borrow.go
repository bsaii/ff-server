package models

import (
	"gorm.io/gorm"
)

type Borrow struct {
	gorm.Model
	Address    string
	Amount     string
	Collateral string
}
