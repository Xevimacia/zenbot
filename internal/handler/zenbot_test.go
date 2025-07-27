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

func TestHandleZenbotRequest_ValidRequest(t *testing.T) {
	// Create a test request
	reqBody := model.ZenbotRequest{
		ConversationID: "test-conversation",
		Message:        "Should we launch the new feature now?",
	}
	jsonBody, _ := json.Marshal(reqBody)

	// Create HTTP request
	req := httptest.NewRequest("POST", "/zenbot", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	// Call the handler
	HandleZenbotRequest(rr, req)

	// Check status code and SSE headers
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}
	if rr.Header().Get("Content-Type") != "text/event-stream" {
		t.Error("Expected Content-Type text/event-stream")
	}

	// Check that we got SSE events
	body := rr.Body.String()
	if !strings.Contains(body, "event: status") {
		t.Error("Expected SSE status events")
	}
	if !strings.Contains(body, "event: message") {
		t.Error("Expected SSE message events")
	}
}

func TestHandleZenbotRequest_InvalidMethod(t *testing.T) {
	req := httptest.NewRequest("GET", "/zenbot", nil)
	rr := httptest.NewRecorder()

	// Test with middleware applied
	MethodMiddleware(http.MethodPost)(HandleZenbotRequest)(rr, req)

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
