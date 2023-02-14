package controllers

import (
	"github.com/dondonudon/json-rest-api/initializers"
	"github.com/dondonudon/json-rest-api/models"
	"github.com/gin-gonic/gin"
)

var body struct {
	Title   string
	Content string
}

func PostsCreate(c *gin.Context) {
	//declare a struct to hold the request body

	//bind the request body to the struct
	c.Bind(&body)

	//create a post
	post := models.Post{Title: body.Title, Content: body.Content}

	//save the post to the database
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return

	}
	c.JSON(200, gin.H{
		"data": post,
	})

}

func PostsGetAll(c *gin.Context) {
	//get all posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	//return the posts
	c.JSON(200, gin.H{
		"data": posts,
	})
}

func PostsGet(c *gin.Context) {
	//get the post if deleted_at is null
	var post models.Post
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(400, gin.H{"error": "Record not found!"})
		return
	}

	//return the post
	c.JSON(200, gin.H{
		"data": post,
	})
}

func PostsUpdate(c *gin.Context) {
	//get the post
	var post models.Post
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(400, gin.H{"error": "Record not found!"})
		return
	}

	//bind the request body to the struct
	c.Bind(&body)

	//update the post
	post.Title = body.Title
	post.Content = body.Content

	//save the post to the database
	result := initializers.DB.Save(&post)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return

	}
	c.JSON(200, gin.H{
		"data": post,
	})
}

func PostsDelete(c *gin.Context) {

	//get the post
	var post models.Post
	if err := initializers.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(400, gin.H{"error": "Record not found!"})
		return
	}

	//delete the post
	initializers.DB.Delete(&post)

	//return the post
	c.JSON(200, gin.H{
		"data": post,
	})
}
