package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	//Get data off request body
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

func PostsShow(c *gin.Context) {

	//Get id off url

	id := c.Param("id")

	// get the post
	var post []models.Post
	initializers.DB.First(&post, id)
	//Respond
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsUpdate(c *gin.Context) {

	// get the id of the url
	id := c.Param("id")
	// get the data of request body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	// find the post we are updating
	var post []models.Post
	initializers.DB.First(&post, id)

	// update it
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// respond with it
	c.JSON(200, gin.H{
		"posts": post,
	})

}

func PostsDelete(c *gin.Context) {

	// get rid of the url
	id := c.Param("id")

	//delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	//respond

	c.Status((200))
}
