package initializers

import (
	"fmt"
	"os"

	"github.com/captainmio/go-crud-post/controllers"
	"github.com/gin-gonic/gin"
)

func Routes() {
	fmt.Println("Setting up routes...")
	router := gin.Default()

	// Posts routes
	router.POST("/posts", controllers.CreatePost)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
