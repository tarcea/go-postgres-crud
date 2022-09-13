package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tarcea/go-crud-01/initializers"
	"github.com/tarcea/go-crud-01/models"
)

func CreatePost(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.Db.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	result := initializers.Db.Find(&posts)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPostById(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	result := initializers.Db.First(&post, id)
	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var post models.Post
	result := initializers.Db.First(&post, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	initializers.Db.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	// id := c.Param("id")
	// var post models.Post
	// result := initializers.Db.First(&post, id)
	// if result.Error != nil {
	// 	c.Status(400)
	// 	return
	// }
	// initializers.Db.Delete(post)
	// c.JSON(200, gin.H{
	// 	"post": post,
	// })
	id := c.Param("id")
	initializers.Db.Delete(&models.Post{}, id)
	c.Status(200)
}
