package models

import (
	"gorm.io/gorm"
)

type Withdraw struct {
	gorm.Model
	Address string
	Amount  int
}
