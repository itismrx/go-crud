package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/itismrx/go-crud/initalizers"
	"github.com/itismrx/go-crud/models"
)

func CreatePost(c *gin.Context) {
	// Get data off req body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)
	// Create a post

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initalizers.DB.Create(&post)

	if result.Error != nil {
		log.Println(result.Error)
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post":          post,
		"affected_rows": result.RowsAffected,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initalizers.DB.Find(&posts)

	// Respond with then

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	// Get the posts
	id := c.Param("id")
	var post models.Post
	initalizers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostUpdate(c *gin.Context) {

	// Get the id off the url
	id := c.Param("id")
	// Get the data off req body
	var body struct {
		Title string
		Body  string
	}
	log.Println(body)
	c.Bind(&body)
	// find the post were updateing
	var post models.Post
	initalizers.DB.First(&post, id)
	// update it
	initalizers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// Respond with it

	c.JSON(200, gin.H{
		"status": true,
		"post":   post,
	})
}

func PostsDelete(c *gin.Context) {
	// Get the id off the url
	id := c.Param("id")

	initalizers.DB.Delete(&models.Post{}, id)

	c.Status(200)
}
