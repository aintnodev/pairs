package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	initDB()
	router := gin.Default()

	// router.StaticFile("/", "public/index.html")
	// router.StaticFile("/results", "public/result.html")
	// router.Static("/style", "public/style")

	router.GET("/api", welcomeMsg)
	router.GET("/api/get", getPair)
	router.POST("/api/add", addPair)
	router.DELETE("/api/delete/:id", deletePair)
	router.Run("localhost:3000")
}
