package types

import (
	"testing"

	"github.com/google/uuid"
	assert "github.com/stretchr/testify/assert"
)

func TestMemoryStore_NewMemoryStore(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store)
	assert.Empty(t, &store.messages)
	assert.NotNil(t, &store.mutex)
}

func TestMemoryStore_GetMessages_Empty(t *testing.T) {
	store := NewMemoryStore()
	messages, err := store.GetMessages()
	assert.NoError(t, err)
	assert.NotNil(t, messages)
	assert.Empty(t, messages)
}

func TestMemoryStore_GetMessage_NotFound(t *testing.T) {
	store := NewMemoryStore()
	message, err := store.GetMessage("non-existent-id")
	assert.Error(t, err)
	assert.Equal(t, ErrMessageNotFound, err)
	assert.Equal(t, Message{}, message)
}

func TestMemoryStore_CreateMessage_Valid(t *testing.T) {
	store := NewMemoryStore()
	message := &Message{
		User: "testuser",
		Text: "This is a test message",
	}
	err := store.CreateMessage(message)
	assert.NoError(t, err)
	assert.NotEmpty(t, message.ID)
	messages, err := store.GetMessages()
	assert.NoError(t, err)
	assert.Len(t, messages, 1)
	assert.Equal(t, message.User, messages[0].User)
	assert.Equal(t, message.Text, messages[0].Text)
}

func TestMemoryStore_CreateMessage_Invalid_NoUser(t *testing.T) {
	store := NewMemoryStore()
	message := &Message{
		User: "", // Invalid user
		Text: "This is a test message",
	}
	err := store.CreateMessage(message)
	assert.Error(t, err)
	assert.Equal(t, ErrUserEmpty, err)
	messages, err := store.GetMessages()
	assert.NoError(t, err)
	assert.Empty(t, messages)
}

func TestMemoryStore_CreateMessage_Invalid_NoText(t *testing.T) {
	store := NewMemoryStore()
	message := &Message{
		User: "testuser",
		Text: "", // Invalid text
	}
	err := store.CreateMessage(message)
	assert.Error(t, err)
	assert.Equal(t, ErrTextEmpty, err)
	messages, err := store.GetMessages()
	assert.NoError(t, err)
	assert.Empty(t, messages)
}

func TestMemoryStore_CreateMessage_Invalid_TextTooLong(t *testing.T) {
	store := NewMemoryStore()
	longText := "a" // Create a long text exceeding 500 characters
	for range 501 {
		longText += "a"
	}
	message := &Message{
		User: "testuser",
		Text: longText,
	}
	err := store.CreateMessage(message)
	assert.Error(t, err)
	assert.Equal(t, ErrTextTooLong, err)
	messages, err := store.GetMessages()
	assert.NoError(t, err)
	assert.Empty(t, messages)
}

func TestMemoryStore_UpdateMessage_Valid(t *testing.T) {
	store := NewMemoryStore()
	message := &Message{
		User: "testuser",
		Text: "test test test test",
	}
	err := store.CreateMessage(message)
	assert.NoError(t, err)
	originalID := message.ID
	updatedMessage := &Message{
		ID:   originalID,
		User: "testuser",
		Text: "updated text",
	}
	err = store.UpdateMessage(originalID.String(), updatedMessage)
	assert.NoError(t, err)
	messages, err := store.GetMessages()
	assert.NoError(t, err)
	assert.Len(t, messages, 1)
	assert.Equal(t, updatedMessage.User, messages[0].User)
	assert.Equal(t, updatedMessage.Text, messages[0].Text)
}

func TestMemoryStore_UpdateMessage_NotFound(t *testing.T) {
	store := NewMemoryStore()
	updatedMessage := &Message{
		ID:   uuid.New(), // Non-existent ID
		User: "testuser",
		Text: "updated text",
	}
	err := store.UpdateMessage(updatedMessage.ID.String(), updatedMessage)
	assert.Error(t, err)
	assert.Equal(t, ErrMessageNotFound, err)
	messages, err := store.GetMessages()
	assert.NoError(t, err)
	assert.Empty(t, messages)
}

func TestMemoryStore_UpdateMessage_UserEmpty(t *testing.T) {
	store := NewMemoryStore()
	message := &Message{
		User: "testuser",
		Text: "test test test test",
	}
	err := store.CreateMessage(message)
	assert.NoError(t, err)
	originalID := message.ID
	updatedMessage := &Message{
		ID:   originalID,
		User: "",
		Text: "updated text",
	}
	err = store.UpdateMessage(originalID.String(), updatedMessage)
	assert.Error(t, err)
	assert.Equal(t, ErrUserEmpty, err)
	messages, err := store.GetMessages()
	assert.NoError(t, err)
	assert.Len(t, messages, 1)
	assert.Equal(t, message.User, messages[0].User)
	assert.Equal(t, message.Text, messages[0].Text)
}

func TestMemoryStore_UpdateMessage_Invalid_TextTooLong(t *testing.T) {
	store := NewMemoryStore()
	longText := "a" // Create a long text exceeding 500 characters
	for range 501 {
		longText += "a"
	}
	message := &Message{
		User: "testuser",
		Text: longText,
	}
	err := store.CreateMessage(message)
	assert.Error(t, err)
	assert.Equal(t, ErrTextTooLong, err)
	messages, err := store.GetMessages()
	assert.NoError(t, err)
	assert.Empty(t, messages)
}

func TestMemoryStore_DeleteMessage_Valid(t *testing.T) {
	store := NewMemoryStore()
	message := &Message{
		User: "testuser",
		Text: "This is a test message",
	}
	err := store.CreateMessage(message)
	assert.NoError(t, err)
	id := message.ID.String()
	err = store.DeleteMessage(id)
	assert.NoError(t, err)
	messages, err := store.GetMessages()
	assert.NoError(t, err)
	assert.Empty(t, messages)
}

func TestMemoryStore_DeleteMessage_NotFound(t *testing.T) {
	store := NewMemoryStore()
	err := store.DeleteMessage("non-existent-id")
	assert.Error(t, err)
	assert.Equal(t, ErrMessageNotFound, err)
	messages, err := store.GetMessages()
	assert.NoError(t, err)
	assert.Empty(t, messages)
}
