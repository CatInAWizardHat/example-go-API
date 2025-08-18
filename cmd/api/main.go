package main

import (
	"example-message-api/internal/message"
	"example-message-api/internal/user"
	"example-message-api/internal/api"
	"fmt"
	"log"
	"os"
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

	messageDB := message.NewMemoryStore()
	userDB := user.NewMemoryStore()

	server := api.NewServer(userDB, messageDB)

	if err := server.Start(hostname); err != nil {
		log.Fatalf("Error, failed to start server: %s", err)
	}
}
