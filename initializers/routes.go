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
	router.GET("/posts", controllers.GetPosts)
	router.GET("/posts/:id", controllers.GetPostByID)
	router.PUT("/posts/:id", controllers.UpdatePost)
	router.DELETE("/posts/:id", controllers.DeletePost)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
