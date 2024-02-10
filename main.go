package main

import (
    "github.com/gin-gonic/gin"
	"main/routes"
)

func main() {
	router := gin.Default()
	routes.Route(router)
	router.Run("localhost:3004")
} 