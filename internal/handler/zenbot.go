package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/xevimacia/zenbot/internal/model"
)

// HandleZenbotRequest processes the /zenbot endpoint
func HandleZenbotRequest(w http.ResponseWriter, r *http.Request) {
	// Parse JSON input
	var req model.ZenbotRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendSSEError(w, "Invalid JSON format")
		return
	}

	// Validate input
	if req.Message == "" {
		sendSSEError(w, "Empty message provided")
		return
	}

	// Handle optional conversation_id - generate new one if omitted
	if req.ConversationID == "" {
		req.ConversationID = "conv-" + fmt.Sprintf("%d", time.Now().Unix())
	}

	// Set SSE headers
	setSSEHeaders(w)

	// Create a context that can be cancelled if client disconnects
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	// Check for client disconnect
	go func() {
		<-ctx.Done()
		cancel()
	}()

	// Send initial status
	sendSSEStatus(w, "Processing dilemma...")

	// For now, we'll simulate the LLM responses
	// In the next task, we'll implement the actual LLM calls with concurrency

	// Simulate Build Fast response
	sendSSEStatus(w, "Build Fast argues")
	time.Sleep(1 * time.Second) // Simulate processing time

	// Simulate Stillness response
	sendSSEStatus(w, "Stillness reflects")
	time.Sleep(1 * time.Second) // Simulate processing time

	// Simulate combining results
	sendSSEStatus(w, "Combining results")
	time.Sleep(500 * time.Millisecond)

	// Simulate resolution forming
	sendSSEStatus(w, "Resolution forming")
	time.Sleep(500 * time.Millisecond)

	// Send the final message with progressive content updates
	finalMessage := "A clear path forward balances ambition with wisdom. The journey of innovation requires both speed and reflection."
	streamMessageProgressively(w, finalMessage)
}

// setSSEHeaders sets the required headers for Server-Sent Events
func setSSEHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Cache-Control")
}

// sendSSEStatus sends a status event via SSE
func sendSSEStatus(w http.ResponseWriter, status string) {
	event := fmt.Sprintf("event: status\ndata: %s\n\n", status)
	w.Write([]byte(event))

	// Flush the writer to ensure the event is sent immediately
	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}
}

// sendSSEError sends an error event via SSE
func sendSSEError(w http.ResponseWriter, errorMsg string) {
	event := fmt.Sprintf("event: status\ndata: Error: %s\n\n", errorMsg)
	w.Write([]byte(event))

	// Flush the writer to ensure the event is sent immediately
	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}
}

// streamMessageProgressively streams a message with progressive content updates
func streamMessageProgressively(w http.ResponseWriter, message string) {
	words := strings.Fields(message)
	messageID := "zenbot-" + fmt.Sprintf("%d", time.Now().Unix())

	// Build up the message progressively
	currentContent := ""

	for i, word := range words {
		if i > 0 {
			currentContent += " "
		}
		currentContent += word

		// Send message_id and content directly in SSE format
		event := fmt.Sprintf("event: message\ndata: message_id: %s, content: %s\n\n", messageID, currentContent)
		w.Write([]byte(event))

		// Flush after each update
		if flusher, ok := w.(http.Flusher); ok {
			flusher.Flush()
		}

		// Add a small delay to simulate progressive typing
		time.Sleep(150 * time.Millisecond)
	}
}
