package models

import "gorm.io/gorm"

type Approval struct {
	gorm.Model
	Owner   string
	Spender string
	Value   string
}
