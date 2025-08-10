package message

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	assert "github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MockMessageStore struct {
	GetMessageResult Message
	GetMessageError  error

	GetMessagesResult []Message
	GetMessagesError  error

	CreateMessageError error
	UpdateMessageError error
	DeleteMessageError error
}

func (m *MockMessageStore) GetMessage(id string) (Message, error) {
	return m.GetMessageResult, m.GetMessageError
}

func (m *MockMessageStore) GetMessages() ([]Message, error) {
	return m.GetMessagesResult, m.GetMessagesError
}

func (m *MockMessageStore) CreateMessage(message *Message) error {
	return m.CreateMessageError
}

func (m *MockMessageStore) UpdateMessage(id string, message *Message) error {
	return m.UpdateMessageError
}

func (m *MockMessageStore) DeleteMessage(id string) error {
	return m.DeleteMessageError
}

func TestHandler_Unit_GetMessage_Found(t *testing.T) {

}

func TestHandler_Unit_GetMessage_NotFound(t *testing.T) {
	mock := &MockMessageStore{
		GetMessageError:  ErrMessageNotFound,
	}

	handler := NewMessageHandler(mock)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "some-non-existent-id"}}
	handler.GetMessage(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.True(t, strings.HasPrefix(w.Header().Get("Content-Type"), "application/json"))

	var resp map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "message not found", resp["error"])
}
