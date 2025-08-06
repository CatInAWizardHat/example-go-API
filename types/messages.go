package types

import (
	"errors"
	"sync"

	"github.com/google/uuid"
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

func (m *MemoryStore) GetMessages() []Message {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.messages
}

func (m *MemoryStore) GetMessage(id string) (Message, error) {
	// Get id from url params
	for _, message := range m.messages {
		if message.ID.String() == id {
			return message, nil
		}
	}

	return Message{}, errors.New("message not found")
}

func (m *MemoryStore) CreateMessage(message Message) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	message.ID = uuid.New()
	m.messages = append(m.messages, message)
	return nil
}

func (m *MemoryStore) UpdateMessage(id string, message *Message) error {
	for idx, message := range m.messages {
		if message.ID.String() == id {
			// Create new message to be added to the list
			if err := validateMessage(&message); err != nil {
				return err
			}
			m.messages[idx] = message
			return nil
		}
	}
	return errors.New("message not found")
}

func (m *MemoryStore) DeleteMessage(id string) error {
	for idx, message := range m.messages {
		if message.ID.String() == id {
			// Code that I found on Stack Overflow
			m.messages = append(m.messages[:idx], m.messages[idx+1:]...)
			return nil
		}
	}
	return errors.New("message not found")
}

func validateMessage(message *Message) error {
	if message.User == "" || message.Text == "" {
		return errors.New("user and text cannot be empty")
	}
	if len(message.Text) > 500 {
		return errors.New("text cannot exceed 500 characters")
	}
	return nil
}
