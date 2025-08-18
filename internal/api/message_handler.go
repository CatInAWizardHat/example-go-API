package api

import (
	"example-message-api/internal/message"
	"github.com/gin-gonic/gin"
	"net/http"
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
		mapErrorToResponse(c, err)
		return
	}
	successResponse(c, http.StatusOK, messages)
}

func (h *MessageHandler) GetMessage(c *gin.Context) {
	id := c.Param("id")
	msg, err := h.Store.GetMessage(id)
	if err != nil {
		mapErrorToResponse(c, err)
		return
	}
	successResponse(c, http.StatusOK, msg)
}

func (h *MessageHandler) CreateMessage(c *gin.Context) {
	var msg message.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Store.CreateMessage(&msg); err != nil {
		mapErrorToResponse(c, err)
		return
	}
	successResponse(c, http.StatusOK, msg)
}

func (h *MessageHandler) UpdateMessage(c *gin.Context) {
	id := c.Param("id")
	var msg message.Message
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Store.UpdateMessage(id, &msg); err != nil {
		mapErrorToResponse(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *MessageHandler) DeleteMessage(c *gin.Context) {
	id := c.Param("id")
	if err := h.Store.DeleteMessage(id); err != nil {
		mapErrorToResponse(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
