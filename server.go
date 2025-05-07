package main

import (	
	"log"
	"github.com/gin-gonic/gin"	
	"example-message-api/services"
)

func main() {
	messages := &services.Message{}
	  r := gin.Default()
	  r.GET("/messages", messages.GetMessages)
	  r.GET("/messages/:id", messages.GetMessage)
	  err := r.Run("localhost:8080")
	  if err != nil {
	    log.Fatalf("Error, failed to start server: %s", err)
	  }
}
