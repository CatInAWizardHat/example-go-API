package user

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type User struct {
	ID            uuid.UUID
	Name          string
	Email         string
	PasswordHash  []byte
	CreatedAt     time.Time
	UpdatedAt     time.Time
	EmailVerified bool
}
