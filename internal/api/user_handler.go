package api

import (
	"example-message-api/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type UserHandler struct {
	Store user.UserStore
}

type RegisterUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding"required,min=8"`
}

type UserProfileResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewUserHandler(store user.UserStore) *UserHandler {
	return &UserHandler{
		Store: store,
	}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.Store.GetUsers()
	if err != nil {
		mapErrorToResponse(c, err)
		return
	}
	successResponse(c, http.StatusOK, users)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usr, err := h.Store.GetUser(id)
	if err != nil {
		mapErrorToResponse(c, err)
	}
	successResponse(c, http.StatusOK, usr)
}
