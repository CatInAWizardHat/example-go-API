package main

import (
	"example-message-api/internal/message"
	"example-message-api/internal/user"
	"example-message-api/internal/api"
	"fmt"
	"log"
	"os"
)

func main() {
	// Get environment variables
	host := os.Getenv("HOST_IP")
	port := os.Getenv("HOST_PORT")
	if host == "" || port == "" {
		log.Fatalf("Failed to load HOST or PORT from env")
	}

	// Generate hostname for Gin router
	hostname := fmt.Sprintf("%s:%s", host, port)

	messageDB := message.NewMemoryStore()
	userDB := user.NewMemoryStore()

	server := api.NewServer(userDB, messageDB)

	if err := server.Start(hostname); err != nil {
		log.Fatalf("Error, failed to start server: %s", err)
	}
}
