package models

import (
	"gorm.io/gorm"
)

type Lent struct {
	gorm.Model
	Address string
	Amount  string
}
