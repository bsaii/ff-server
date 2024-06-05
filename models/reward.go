package models

import (
	"gorm.io/gorm"
)

type Reward struct {
	gorm.Model
	Address string
	Reward  int
}
