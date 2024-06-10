package models

import "time"

type Approval struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`

	Owner   string
	Spender string
	Value   string
}
