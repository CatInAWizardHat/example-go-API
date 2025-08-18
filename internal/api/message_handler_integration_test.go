package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"example-message-api/internal/message"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpTest() (*MessageHandler, error) {
	store := message.NewMemoryStore()
	handler := NewMessageHandler(store)
	return handler, nil
}

func TestHandler_Integration_GetMessages_Empty(t *testing.T) {
	handler, _ := SetUpTest()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.GetMessages(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, strings.HasPrefix(w.Header().Get("Content-Type"), "application/json"))

	var resp []message.Message
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Empty(t, resp)
}

func TestHandler_Integration_GetMessages_NotEmpty(t *testing.T) {
	handler, _ := SetUpTest()

	msg := &message.Message{
		User: "testuser",
		Text: "This is a test message",
	}
	err := handler.Store.CreateMessage(msg)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.GetMessages(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, strings.HasPrefix(w.Header().Get("Content-Type"), "application/json"))

	var resp []message.Message
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)
	assert.Equal(t, msg.User, resp[0].User)
	assert.Equal(t, msg.Text, resp[0].Text)
	assert.Equal(t, msg.ID, resp[0].ID)
}

func TestHandler_Integration_GetMessage_Valid(t *testing.T) {
	handler, _ := SetUpTest()

	msg := &message.Message{
		User: "testuser",
		Text: "This is a test message",
	}
	err := handler.Store.CreateMessage(msg)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: msg.ID.String()}}
	handler.GetMessage(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, strings.HasPrefix(w.Header().Get("Content-Type"), "application/json"))

	var resp message.Message
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, msg.User, resp.User)
	assert.Equal(t, msg.Text, resp.Text)
	assert.Equal(t, msg.ID, resp.ID)
}

func TestHandler_Integration_GetMessage_NotFound(t *testing.T) {
	handler, _ := SetUpTest()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.GetMessage(c)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.True(t, strings.HasPrefix(w.Header().Get("Content-Type"), "application/json"))

	var resp map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "message not found", resp["error"])
}
