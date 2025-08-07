package types

import (
	"errors"
	"sync"

	"github.com/google/uuid"
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

type MemoryStore struct {
	messages []Message
	mutex    sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		messages: make([]Message, 0),
	}
}

func (m *MemoryStore) GetMessages() ([]Message, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.messages, nil
}

func (m *MemoryStore) GetMessage(id string) (Message, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	for _, message := range m.messages {
		if message.ID.String() == id {
			return message, nil
		}
	}

	return Message{}, ErrMessageNotFound
}

func (m *MemoryStore) CreateMessage(message *Message) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	if err := validateMessage(message); err != nil {
		return err
	}
	message.ID = uuid.New()
	m.messages = append(m.messages, *message)
	return nil
}

func (m *MemoryStore) UpdateMessage(id string, updatedMessage *Message) error {
	if err := validateMessage(updatedMessage); err != nil {
		return err
	}
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for idx, message := range m.messages {
		if message.ID.String() == id {
			// Create new message to be added to the list
			updatedMessage.ID = message.ID // Keep the same ID
			m.messages[idx] = *updatedMessage
			return nil
		}
	}
	return ErrMessageNotFound
}

func (m *MemoryStore) DeleteMessage(id string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	for idx, message := range m.messages {
		if message.ID.String() == id {
			// Re-slice the array to remove the message
			m.messages = append(m.messages[:idx], m.messages[idx+1:]...)
			return nil
		}
	}
	return ErrMessageNotFound
}

func validateMessage(message *Message) error {
	if message.User == "" {
		return ErrUserEmpty
	}
	if message.Text == "" {
		return ErrTextEmpty
	}
	if len(message.Text) > 500 {
		return ErrTextTooLong
	}
	return nil
}
