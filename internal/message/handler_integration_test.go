package services

import (
	"encoding/json"
	"example-message-api/types"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpTest() (*MessageHandler, error) {
	store := types.NewMemoryStore()
	handler := NewMessageHandler(store)
	return handler, nil
}

func TestHandler_GetMessages_Empty(t *testing.T) {
	handler, _ := SetUpTest()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.GetMessages(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, strings.HasPrefix(w.Header().Get("Content-Type"), "application/json"))

	var resp []types.Message
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Empty(t, resp)
}

func TestHandler_GetMessages_NotEmpty(t *testing.T) {
	handler, _ := SetUpTest()

	message := &types.Message{
		User: "testuser",
		Text: "This is a test message",
	}
	err := handler.Store.CreateMessage(message)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.GetMessages(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, strings.HasPrefix(w.Header().Get("Content-Type"), "application/json"))

	var resp []types.Message
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Len(t, resp, 1)
	assert.Equal(t, message.User, resp[0].User)
	assert.Equal(t, message.Text, resp[0].Text)
	assert.Equal(t, message.ID, resp[0].ID)
}

func TestHandler_GetMessage_Valid(t *testing.T) {
	handler, _ := SetUpTest()

	message := &types.Message{
		User: "testuser",
		Text: "This is a test message",
	}
	err := handler.Store.CreateMessage(message)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: message.ID.String()}}
	handler.GetMessage(c)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, strings.HasPrefix(w.Header().Get("Content-Type"), "application/json"))

	var resp types.Message
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, message.User, resp.User)
	assert.Equal(t, message.Text, resp.Text)
	assert.Equal(t, message.ID, resp.ID)
}

func TestHandler_GetMessage_NotFound(t *testing.T) {
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
