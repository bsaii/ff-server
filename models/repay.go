package models

import (
	"gorm.io/gorm"
)

type Repay struct {
	gorm.Model
	Address string
	Amount  int
}
