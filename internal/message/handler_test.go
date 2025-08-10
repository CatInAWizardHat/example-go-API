package message

import (
	"github.com/gin-gonic/gin"
	assert "github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockMessageStore struct {
	GetMessageResult  Message
	GetMessagesResult []Message
	GetMessageError   error
}

func (m *MockMessageStore) GetMessage(id string) (Message, error) {
	return m.GetMessageResult, m.GetMessageError
}

func (m *MockMessageStore) GetMessages() ([]Message, error) {
	return m.GetMessagesResult, m.GetMessageError
}

func (m *MockMessageStore) CreateMessage(message *Message) error {
	return m.GetMessageError
}

func (m *MockMessageStore) UpdateMessage(id string, message *Message) error {
	return m.GetMessageError
}

func (m *MockMessageStore) DeleteMessage(id string) error {
	return m.GetMessageError
}

func TestHandler_GetMessage(t *testing.T) {
	mock := &MockMessageStore{
		GetMessageResult: Message{},
		GetMessageError:  nil,
	}

	handler := NewMessageHandler(mock)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.GetMessages(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, strings.HasPrefix(w.Header().Get("Content-Type"), "application/json"))
}
