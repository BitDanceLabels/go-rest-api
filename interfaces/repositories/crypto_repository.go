package repositories

import (
	"github.com/BitDanceLabels/go-rest-api/domain/entities"
	"github.com/jinzhu/gorm"
)

type CryptoRepository struct {
	DB *gorm.DB
}

func NewCryptoRepository(db *gorm.DB) *CryptoRepository {
	return &CryptoRepository{
		DB: db,
	}
}

func (r *CryptoRepository) SaveCryptoData(data *entities.CryptoData) error {
	// Save crypto data to the database
	return r.DB.Create(data).Error
}
