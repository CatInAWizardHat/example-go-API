package main

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
		log.Print("Successful GET Request.")
	})

	r.Run("localhost:8080")
}
