package api

import (
	// ... (các imports khác)
	"github.com/BitDanceLabels/go-rest-api/interfaces/repositories"
)

// InitAPIClient khởi tạo client API và thiết lập giá trị COINGECKO_API_KEY
func InitAPIClient() {
	// Thiết lập giá trị COINGECKO_API_KEY
	repositories.SetAPIKey("CG-psD9U7zRqPYWHqKH3p4r6cCG")

	// Khởi tạo và cấu hình client API ở đây
	// ...

	// Ví dụ:
	// client := yourAPIClientInitializationFunction()
	// repositories.SetAPIClient(client)
}
