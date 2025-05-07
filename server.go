package main

import (	
	"github.com/gin-gonic/gin"	
	"example-message-api/services"
)

func main() {
	messages := &services.Message{}
	  r := gin.Default()
	  r.GET("/messages", messages.GetMessages)
	  r.Run("localhost:8080")
}
