package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"example-message-api/internal/message"
	"example-message-api/internal/user"
	"log"
	"net/http"
)

func mapErrorToResponse(c *gin.Context, err error) {
	switch {
	// Message Route Errors
	case errors.Is(err, message.ErrMessageNotFound):
		errorResponse(c, http.StatusNotFound, err.Error())
	case errors.Is(err, message.ErrTextEmpty):
		errorResponse(c, http.StatusBadRequest, err.Error())
	case errors.Is(err, message.ErrTextTooLong):
		errorResponse(c, http.StatusBadRequest, err.Error())
	case errors.Is(err, message.ErrUserEmpty):
		errorResponse(c, http.StatusBadRequest, err.Error())
	// User Route Errors
	case errors.Is(err, user.ErrUserNotFound):
		errorResponse(c, http.StatusNotFound, err.Error())
	default:
		log.Printf("An unexpected error occured: %v", err.Error())
		errorResponse(c, http.StatusInternalServerError, "An unexpected internal error has occured.")
	}
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"status":  "error",
		"message": message,
	})
}

func successResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"status": "success",
		"data": data,
	})
}
