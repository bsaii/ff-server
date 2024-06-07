package models

import (
	"gorm.io/gorm"
)

type Stake struct {
	gorm.Model
	Address string
	Amount  string
}
