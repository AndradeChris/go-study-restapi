package controllers

import (
	"net/http"
	"main/albumsBD"
	"github.com/gin-gonic/gin"
)

 func PostAlbum(c *gin.Context) {
	var requestBody map[string]interface{}
	
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	id, idExists := requestBody["id"].(string)
	title, titleExists := requestBody["title"].(string)
	artist, artistExists := requestBody["artist"].(string)
	price, priceExists := requestBody["price"].(float64)

	if !idExists || !titleExists || !artistExists || !priceExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required fields"})
		return
	}

	albumsBD.Albums = append(albumsBD.Albums, albumsBD.AlbumType{
		ID:     id,
		Title:  title,
		Artist: artist,
		Price:  price,
	})

	c.JSON(http.StatusCreated, gin.H{"message": "Album created successfully"})
}
