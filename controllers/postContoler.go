package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//Get data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	//create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {

		c.Status(400)
		return
	}

	//Return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {

	// get the post
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond
	c.JSON(200, gin.H{
		"posts": posts,
	})
}
