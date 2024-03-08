// /interfaces/controllers/crypto_controller.go
package controllers

import (
	"github.com/gin-gonic/gin"
)

type CryptoController struct {
	// Add usecase or service dependency
}

func (controller *CryptoController) SearchHandler(c *gin.Context) {
	// Implement search logic using usecase or service
}

func (controller *CryptoController) GetCryptoDetailHandler(c *gin.Context) {
	// Implement get detail logic using usecase or service
}
