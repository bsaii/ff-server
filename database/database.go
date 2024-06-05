package database

import (
	"github.com/bsaii/ff-server/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(dbPath string) error {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto-migrate the Event model
	return DB.AutoMigrate(&models.Borrow{}, &models.Lent{}, &models.Repay{}, &models.Reward{}, &models.Stake{}, &models.User{}, &models.Withdraw{})
}
