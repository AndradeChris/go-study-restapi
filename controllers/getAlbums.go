package controllers

import (
	"net/http"
    "github.com/gin-gonic/gin"
	"main/albumsBD"
)

func GetAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albumsBD.Albums)
}
