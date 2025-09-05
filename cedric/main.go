package main

import (
	"example.com/hello/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/albums", controllers.GetAlbums)
	router.GET("/albums", controllers.GetAlbumByID)
	router.POST("/albums", controllers.PostAlbums)
	router.Run("localhost:8080")
}
