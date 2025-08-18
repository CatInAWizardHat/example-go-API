package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"example-message-api/internal/message"
)

type MessageHandler struct {
	Store message.MessageStore
}

func NewMessageHandler(store message.MessageStore) *MessageHandler {
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
	msg, err := h.Store.GetMessage(id)
	if err != nil {
		if errors.Is(err, message.ErrMessageNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, msg)
}

func (h *MessageHandler) CreateMessage(c *gin.Context) {
	var msg message.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Store.CreateMessage(&msg); err != nil {
		if errors.Is(err, message.ErrUserEmpty) || errors.Is(err, message.ErrTextEmpty) || errors.Is(err, message.ErrTextTooLong) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Handle other errors, such as database errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, msg)
}

func (h *MessageHandler) UpdateMessage(c *gin.Context) {
	id := c.Param("id")
	var msg message.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Store.UpdateMessage(id, &msg); err != nil {
		if errors.Is(err, message.ErrMessageNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, msg)
}

func (h *MessageHandler) DeleteMessage(c *gin.Context) {
	id := c.Param("id")
	if err := h.Store.DeleteMessage(id); err != nil {
		if errors.Is(err, message.ErrMessageNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.Status(http.StatusNoContent)
}
