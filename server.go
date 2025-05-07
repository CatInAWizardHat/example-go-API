package main

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r, err := gin.Default()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
