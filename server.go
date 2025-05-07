package main

import (	
	"fmt"
	"os"
	"log"
	"example-message-api/services"

	"github.com/gin-gonic/gin"	
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error, failed to load .env: %s", err)
	}
	host := os.Getenv("HOST_IP")
	port := os.Getenv("HOST_PORT")
	hostname := fmt.Sprintf("%s:%s", host, port)

	messages := &services.Message{}
	  r := gin.Default()
	  r.GET("/messages", messages.GetMessages)
	  r.GET("/messages/:id", messages.GetMessage)

	  err = r.Run(hostname)
	  if err != nil {
	    log.Fatalf("Error, failed to start server: %s", err)
	  }
}
