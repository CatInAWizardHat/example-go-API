package types

import (
	"testing"

	"strings"

	assert "github.com/stretchr/testify/assert"
)

func TestMemoryStore_NewMemoryStore(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store)
	assert.Empty(t, store.messages)
	assert.NotNil(t, &store.mutex)
}

func TestMemoryStore_GetMessage_AfterCreation(t *testing.T) {
	store := NewMemoryStore()
	message := &Message{
		User: "testuser",
		Text: "This is a test message",
	}
	err := store.CreateMessage(message)
	assert.NoError(t, err)
	retrievedMessage, err := store.GetMessage(message.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, message.User, retrievedMessage.User)
	assert.Equal(t, message.Text, retrievedMessage.Text)
	assert.Equal(t, message.ID, retrievedMessage.ID)
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

func TestMemoryStore_CreateMessage_Validation(t *testing.T) {
	testCases := []struct {
		name          string
		inputMessage  *Message
		expectedError error
	}{
		{
			name:          "Invalid_NoUser",
			inputMessage:  &Message{User: "", Text: "This is a test message"},
			expectedError: ErrUserEmpty,
		},
		{
			name:          "Invalid_NoText",
			inputMessage:  &Message{User: "testuser", Text: ""},
			expectedError: ErrTextEmpty,
		},
		{
			name:          "Invalid_TextTooLong",
			inputMessage:  &Message{User: "testuser", Text: strings.Repeat("a", 501)},
			expectedError: ErrTextTooLong,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			store := NewMemoryStore()
			err := store.CreateMessage(tc.inputMessage)
			assert.Error(t, err)
			assert.Equal(t, tc.expectedError, err)
		})
	}
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

func TestMemoryStore_UpdateMessage_Validation(t *testing.T) {
	store := NewMemoryStore()
	message := &Message{
		User: "testuser",
		Text: "test test test test",
	}
	err := store.CreateMessage(message)
	assert.NoError(t, err)
	originalID := message.ID

	testCases := []struct {
		name          string
		id            string
		inputMessage  *Message
		expectedError error
	}{
		{
			name:          "Invalid_NotFound",
			id:            "non-existent-id",
			inputMessage:  &Message{User: "testuser", Text: "updated text"},
			expectedError: ErrMessageNotFound,
		},
		{
			name:          "Invalid_NoUser",
			id:            originalID.String(),
			inputMessage:  &Message{User: "", Text: "updated text"},
			expectedError: ErrUserEmpty,
		},
		{
			name:          "Invalid_NoText",
			id:            originalID.String(),
			inputMessage:  &Message{User: "testuser", Text: ""},
			expectedError: ErrTextEmpty,
		},
		{
			name:          "Invalid_TextTooLong",
			id:            originalID.String(),
			inputMessage:  &Message{User: "testuser", Text: strings.Repeat("a", 501)},
			expectedError: ErrTextTooLong,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := store.UpdateMessage(tc.id, tc.inputMessage)
			assert.Error(t, err)
			assert.Equal(t, tc.expectedError, err)
			messages, err := store.GetMessages()
			assert.NoError(t, err)
			assert.Equal(t, len(messages), 1)
			assert.Equal(t, message.ID, messages[0].ID)
			assert.Equal(t, message.User, messages[0].User)
			assert.Equal(t, message.Text, messages[0].Text)
		})
	}
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
