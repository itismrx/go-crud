package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itismrx/go-crud/controllers"
	"github.com/itismrx/go-crud/initalizers"
)

func init() {
	initalizers.LoadEnvVariables()
	initalizers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.Run()
}
