package controllers

import (
	"github.com/captainmio/go-crud-post/db"
	"github.com/captainmio/go-crud-post/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	// bind request body to struct
	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	post := models.Post{Title: body.Title, Content: body.Content}

	// create post in database
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

	// retrieve all posts
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

func GetPostByID(c *gin.Context) {
	// get id param
	id := c.Param("id")
	var post models.Post

	// find post by id
	result := db.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Post not found",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    post,
	})
}

func UpdatePost(c *gin.Context) {
	// get id param
	id := c.Param("id")
	var post models.Post

	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	// find post by id
	result := db.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"success": false,
			"message": "Post not found",
		})
		return
	}

	// update post
	post.Title = body.Title
	post.Content = body.Content
	result = db.DB.Save(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"success": false,
			"message": "Failed to update post",
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    post,
	})
}
