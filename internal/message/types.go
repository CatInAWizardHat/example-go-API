package message

import(
	"errors"
	"google/uuid"
)

var (
	ErrMessageNotFound = errors.New("message not found")
	ErrUserEmpty       = errors.New("user cannot be empty")
	ErrTextEmpty       = errors.New("text cannot be empty")
	ErrTextTooLong     = errors.New("text cannot exceed 500 characters")
)

// Message struct
type Message struct {
	ID   uuid.UUID `json:"id"`
	User string    `json:"user"`
	Text string    `json:"text"`
}
