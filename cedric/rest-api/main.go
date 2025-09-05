package main

import (
	"example.com/hello/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	albumsGroup := router.Group("/albums")
	controllers.RegisterAlbumRoutes(albumsGroup)
	router.Run("localhost:8080")
}
