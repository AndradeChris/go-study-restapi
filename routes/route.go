package routes 

import (
	"main/controllers"
	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine) {
	router.GET("/albums", controllers.GetAlbums)
	router.POST("/albums", controllers.PostAlbum)
}