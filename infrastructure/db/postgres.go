package db

import (
	"github.com/BitDanceLabels/go-rest-api/domain/entities"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=BitDance dbname=BitDance_Servers sslmode=disable")
	if err != nil {
		return nil, err
	}

	// Auto-migrate tables
	db.AutoMigrate(&entities.CryptoData{})

	return db, nil
}
