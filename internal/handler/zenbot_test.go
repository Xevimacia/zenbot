package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/xevimacia/zenbot/internal/model"
)

func TestHandleZenbotRequest_EmptyMessage(t *testing.T) {
	// Create a test request with empty message
	reqBody := model.ZenbotRequest{
		ConversationID: "test-conversation",
		Message:        "",
	}
	jsonBody, _ := json.Marshal(reqBody)

	// Create HTTP request
	req := httptest.NewRequest("POST", "/v1/zenbot", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Call the handler with SSE middleware
	SSEMiddleware(HandleZenbotRequest)(rr, req)

	// Check that we got an error
	body := rr.Body.String()
	if !strings.Contains(body, "empty message provided") {
		t.Error("Expected error for empty message")
	}
}

func TestHandleZenbotRequest_InvalidMethod(t *testing.T) {
	req := httptest.NewRequest("GET", "/v1/zenbot", nil)
	rr := httptest.NewRecorder()

	// Test with middleware applied
	handlerFunc := func(w http.ResponseWriter, r *http.Request) {
		HandleZenbotRequest(w, r)
	}

	MethodMiddleware(http.MethodPost)(handlerFunc)(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status 405, got %d", rr.Code)
	}
}

func TestStreamMessageProgressively(t *testing.T) {
	rr := httptest.NewRecorder()
	testMessage := "Hello world"
	streamMessageProgressively(rr, testMessage)

	body := rr.Body.String()

	// Check progressive content
	if !strings.Contains(body, "content: Hello") {
		t.Error("Expected progressive content")
	}
	if !strings.Contains(body, "content: Hello world") {
		t.Error("Expected complete progressive content")
	}
	if !strings.Contains(body, "message_id:") {
		t.Error("Expected message_id in response")
	}
}
