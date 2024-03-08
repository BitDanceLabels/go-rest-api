package main

import (
	"os"

	"github.com/BitDanceLabels/go-rest-api/infrastructure/api"
	"github.com/BitDanceLabels/go-rest-api/infrastructure/cache"
	"github.com/BitDanceLabels/go-rest-api/infrastructure/db"
	"github.com/BitDanceLabels/go-rest-api/interfaces/controllers"
	"github.com/BitDanceLabels/go-rest-api/interfaces/repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	// Lấy giá trị Coingecko API key từ biến môi trường
	apiKey := os.Getenv("COINGECKO_API_KEY")

	// Kiểm tra xem có giá trị không, nếu không, đặt một giá trị mặc định
	if apiKey == "" {
		apiKey = "CG-psD9U7zRqPYWHqKH3p4r6cCG"
	}

	// Set giá trị Coingecko API key thành một biến môi trường
	os.Setenv("COINGECKO_API_KEY", apiKey)

	// Initialize components
	cache.InitCache()
	db.InitDB()

	// Use Coingecko API key from environment variable
	api.InitAPIClient(apiKey)

	// Inject dependencies
	cryptoRepository := repositories.NewCryptoRepository(db.DB)
	cryptoController := controllers.CryptoController{
		// Inject other dependencies like usecases or services
	}

	// Register routes
	router := gin.Default()
	router.GET("/search", cryptoController.SearchHandler)
	router.GET("/crypto-detail/:id", cryptoController.GetCryptoDetailHandler)
	router.Static("/static", "./presentation")

	// ... (other routes)

	router.Run(":8080")
}
