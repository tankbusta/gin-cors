package cors_test

import (
	"time"

	"github.com/itsjamie/gin-cors"
	"gopkg.in/gin-gonic/gin.v1"
)

func ExampleMiddleware() {
	// Initialize the gin-gonic router
	router := gin.New()

	// Set up CORS middleware options
	config := cors.Config{
		Origins:         "*",
		RequestHeaders:  "Authorization",
		Methods:         "GET, POST, PUT",
		Credentials:     true,
		ValidateHeaders: false,
		MaxAge:          1 * time.Minute,
	}

	// Apply the middleware to the router (works on groups too)
	router.Use(cors.Middleware(config))
}
