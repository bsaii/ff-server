package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	StakingBalance    int
	LendingBalance    int
	BorrowingBalance  int
	Rewards           int
	LastUpdated       int
	CollateralBalance int
}
