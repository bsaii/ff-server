package models

import "gorm.io/gorm"

type Transfer struct {
	gorm.Model
	From  string
	To    string
	Value string
}
