package controllers

import (
	"github.com/captainmio/go-crud-post/db"
	"github.com/captainmio/go-crud-post/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	post := models.Post{Title: body.Title, Content: body.Content}

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

func GetPosts(c *gin.Context) {
	var posts []models.Post

	result := db.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Failed to retrieve posts",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    posts,
	})
}
