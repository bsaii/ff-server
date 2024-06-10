package models

import "time"

type Transfer struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`

	From  string
	To    string
	Value string
}
