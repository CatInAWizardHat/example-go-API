package main

import (
	"example-message-api/services"
	"example-message-api/types"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file to make it accessible
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error, failed to load .env: %s", err)
	}

	// Get environment variables
	host := os.Getenv("HOST_IP")
	port := os.Getenv("HOST_PORT")
	// Generate hostname for Gin router
	hostname := fmt.Sprintf("%s:%s", host, port)

	messageDB := types.NewMemoryStore()
	// Create operator for the endpoints
	messages := services.NewMessageHandler(messageDB)
	// Create router
	r := gin.Default()
	r.GET("/messages", messages.GetMessages)
	r.GET("/messages/:id", messages.GetMessage)
	r.POST("/messages", messages.CreateMessage)
	r.PATCH("/messages/:id", messages.UpdateMessage)
	r.DELETE("/messages/:id", messages.DeleteMessage)

	if err := r.Run(hostname); err != nil {
		log.Fatalf("Error, failed to start server: %s", err)
	}
}
