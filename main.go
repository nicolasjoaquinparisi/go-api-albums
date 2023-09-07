package main

import (
	"github.com/gin-gonic/gin"
	"go-api-albums/handlers"
)

func main() {
	router := gin.Default()

	router.GET("/ping", handlers.Ping)
	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbumByID)
	router.POST("/albums", handlers.PostAlbums)

	if err := router.Run("localhost:8080"); err != nil {
		return
	}
}
