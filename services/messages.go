package services

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


type Message struct {
	ID	string	`json:"id"`
	User	string	`json:"user"`
	Text	string	`json:"text"`
}

var messages = []Message{
	{ID: "1", User: "Alex", Text: "Ping"},
	{ID: "2", User: "API", Text: "Pong"},
}

func (m *Message) GetMessages(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, messages)
}

func (m *Message) GetMessage(c *gin.Context) {
	id := c.Param("id")
	for _, message := range messages {
		if message.ID == id {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": message})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Message not found."})
}
