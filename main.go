package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tarcea/go-crud-01/controllers"
	"github.com/tarcea/go-crud-01/initializers"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.POST("/posts", controllers.CreatePost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPostById)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run() // listen and serve on 0.0.0.0:8080 if no PORT in .env
}
