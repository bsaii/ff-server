package models

import "time"

type Stake struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`

	Address string
	Amount  string
}
