package controllers

import (
	"fmt"

	"github.com/captainmio/go-crud-post/db"
	"github.com/captainmio/go-crud-post/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	fmt.Println("Creating a new post...")

	post := models.Post{Title: "Test 111", Content: "This is a test post."}

	result := db.DB.Create(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Failed to create post",
		})
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    post,
	})
}
