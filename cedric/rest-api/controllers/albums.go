package controllers

import (
	"net/http"

	"example.com/hello/models/dto"
	"github.com/gin-gonic/gin"
)

func RegisterAlbumRoutes(rg *gin.RouterGroup) {
	rg.GET("/", getAlbums)
	rg.GET("/:id", getAlbumByID)
	rg.POST("/", postAlbums)
}

func getAlbums(c *gin.Context) {
	// Get all albums
	c.IndentedJSON(http.StatusOK, dto.Albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum dto.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	dto.Albums = append(dto.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range dto.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
