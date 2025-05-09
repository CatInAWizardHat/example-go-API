package services

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Message struct
type Message struct {
	ID		string	`json:"id"`
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
	// Get id from url params
	id := c.Param("id")
	for _, message := range messages {
		if message.ID == id {
			// If found, return 200 and the message
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": message})
			return
		}
	}
	// If not found, return 404
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Message not found."})
}

func (m *Message) CreateMessage(c *gin.Context) {
	var newMessage Message
	newMessage.ID = strconv.Itoa(len(messages) + 1)
	if err := c.BindJSON(&newMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}
	messages = append(messages, newMessage)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "data": newMessage})
}

func (m *Message) UpdateMessage(c *gin.Context) {
	id := c.Param("id")
	for idx, message := range messages {
		if message.ID == id {
			var updatedMessage Message
			if err := c.BindJSON(&updatedMessage); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
				return
			}
			updatedMessage.ID = id
			messages[idx] = updatedMessage
			c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Message not found."})
}

func (m *Message) DeleteMessage(c *gin.Context) {
	id := c.Param("id")
	for idx, message := range messages {
		if message.ID == id {
			messages = append(messages[:idx], messages[idx+1:]...)
			c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Message not found."})
}
