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

