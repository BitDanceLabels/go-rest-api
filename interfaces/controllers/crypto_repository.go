// /interfaces/repositories/crypto_repository.go
package repositories

import (
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

// Implement repository methods as needed
