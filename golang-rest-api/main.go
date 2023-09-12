package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", RootHandler)

	router.GET("/hello", HelloRootHandler)

	router.Run()
}

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Mukhammad Fahlevi Ali",
		"bio":  "A software Engineer",
	})
}

func HelloRootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content": "Belajar golang",
		"bio":     "belajar golang bareng levi gaskeun!!",
	})
}
