package api

import (
	"example-message-api/internal/message"
	"example-message-api/internal/user"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router       *gin.Engine
	userStore    user.UserStore
	messageStore message.MessageStore
}

func NewServer(userStore user.UserStore, messageStore message.MessageStore) *Server {
	server := &Server{
		userStore:    userStore,
		messageStore: messageStore,
	}

	router := gin.Default()

	userAPI := NewUserHandler(server.userStore)
	messageAPI := NewMessageHandler(server.messageStore)

	// User routes
	router.GET("/users", userAPI.GetUsers)
	router.GET("/users/:id", userAPI.GetUser)

	// Message routes
	router.GET("/messages", messageAPI.GetMessages)
	router.GET("/messages/:id", messageAPI.GetMessage)
	router.POST("/messages", messageAPI.CreateMessage)
	router.PATCH("/messages/:id", messageAPI.UpdateMessage)
	router.DELETE("/messages/:id", messageAPI.DeleteMessage)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
