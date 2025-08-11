package main

import (
	"github.com/captainmio/go-crud-post/initializers"
	"github.com/captainmio/go-crud-post/models"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
