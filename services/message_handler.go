package services

import (
	"net/http"

	"example-message-api/types"

	"github.com/gin-gonic/gin"
)

type Message = types.Message
type MessageStore = types.MessageStore

type MessageHandler struct {
	Store MessageStore
}

func NewMessageHandler(store MessageStore) *MessageHandler {
	return &MessageHandler{
		Store: store,
	}
}

func (h *MessageHandler) GetMessages(c *gin.Context) {
	messages, err := h.Store.GetMessages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, messages)
}

func (h *MessageHandler) GetMessage(c *gin.Context) {
	id := c.Param("id")
	message, err := h.Store.GetMessage(id)
	if err != nil {
		if err.Error() == "message not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) CreateMessage(c *gin.Context) {
	var message Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Store.CreateMessage(&message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, message)
}

func (h *MessageHandler) UpdateMessage(c *gin.Context) {
	id := c.Param("id")
	var message Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Store.UpdateMessage(id, &message); err != nil {
		if err.Error() == "message not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, message)
}

func (h *MessageHandler) DeleteMessage(c *gin.Context) {
	id := c.Param("id")
	if err := h.Store.DeleteMessage(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
