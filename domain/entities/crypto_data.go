package entities

import "github.com/jinzhu/gorm"

type CryptoData struct {
	gorm.Model
	Symbol string
	Open   float64
	Close  float64
	High   float64
	Low    float64
	Time   string
	Change float64
}
