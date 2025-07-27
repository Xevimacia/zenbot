package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/xevimacia/zenbot/internal/llm"
	"github.com/xevimacia/zenbot/internal/model"
)

// argumentResponse represents the JSON response from LLM agents
type argumentResponse struct {
	Argument string `json:"argument"`
}

// agentConfig defines the configuration for the two LLM agents
type agentConfig struct {
	name      string
	prompt    string
	statusMsg string
}

// agents defines the static configurations for the two LLM agents
var agents = []agentConfig{
	{
		name:      "BuildFast",
		prompt:    llm.BUILD_FAST_PROMPT,
		statusMsg: "Build Fast argues",
	},
	{
		name:      "Stillness",
		prompt:    llm.STILLNESS_PROMPT,
		statusMsg: "Stillness reflects",
	},
}

// sendSSEEvent sends an event via SSE with the specified event type
func sendSSEEvent(w http.ResponseWriter, eventType string, data string) {
	event := fmt.Sprintf("event: %s\ndata: %s\n\n", eventType, data)
	w.Write([]byte(event))

	// Flush the writer to ensure the event is sent immediately
	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}
}

// sendSSEStatus sends a status event via SSE
func sendSSEStatus(w http.ResponseWriter, status string) {
	sendSSEEvent(w, "status", status)
}

// sendSSEError sends an error event via SSE
func sendSSEError(w http.ResponseWriter, errorMsg string) {
	sendSSEEvent(w, "error", errorMsg)
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

// callLLMAgent calls a specific LLM agent and returns the response via channel
func callLLMAgent(ctx context.Context, w http.ResponseWriter, llmService *llm.OpenAIClient, config agentConfig, message string, conversationHistory string, responseChan chan<- *argumentResponse, errorChan chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()
	sendSSEStatus(w, config.statusMsg)

	// Format the prompt with the user's dilemma and conversation history
	prompt := fmt.Sprintf(config.prompt, message, conversationHistory)

	// Call the LLM agent
	response, err := llmService.GenerateForAgent(ctx, config.name, prompt)
	if err != nil {
		errorChan <- fmt.Errorf("%s error: %w", config.name, err)
		return
	}

	// Parse the JSON response
	var argResp argumentResponse
	if err := json.Unmarshal([]byte(response), &argResp); err != nil {
		errorChan <- fmt.Errorf("failed to parse %s response: %w", config.name, err)
		return
	}

	responseChan <- &argResp
	sendSSEStatus(w, fmt.Sprintf("Got response from %s", config.name))
}

// validateRequest validates the incoming request
func validateRequest(req model.ZenbotRequest) error {
	if req.Message == "" {
		return fmt.Errorf("empty message provided")
	}
	return nil
}

// setupConversationID ensures the conversation ID is set
func setupConversationID(req *model.ZenbotRequest) {
	if req.ConversationID == "" {
		req.ConversationID = "conv-" + fmt.Sprintf("%d", time.Now().Unix())
	}
}

// orchestrateLLMResponse coordinates the LLM calls and returns the final response
func orchestrateLLMResponse(ctx context.Context, w http.ResponseWriter, req model.ZenbotRequest, conversationHistory string) (string, error) {
	// Initialize LLM service
	llmService, err := llm.NewOpenAIClient()
	if err != nil {
		return "", fmt.Errorf("failed to initialize LLM service: %v", err)
	}

	// Create channels for LLM responses
	responseChans := make([]chan *argumentResponse, len(agents))
	for i := range responseChans {
		responseChans[i] = make(chan *argumentResponse, 1)
	}
	errorChan := make(chan error, len(agents))

	// Use WaitGroup to synchronize the concurrent LLM calls
	var wg sync.WaitGroup
	wg.Add(len(agents))

	// Call all agents concurrently
	for i, agent := range agents {
		go callLLMAgent(ctx, w, llmService, agent, req.Message, conversationHistory, responseChans[i], errorChan, &wg)
	}

	// Wait for all LLMs to complete
	wg.Wait()

	// Check for errors from any LLM
	select {
	case err := <-errorChan:
		return "", err
	default:
		// No errors, continue
	}

	// Get responses from channels
	buildFastResp := <-responseChans[0]
	stillnessResp := <-responseChans[1]

	// Send status for combining results
	sendSSEStatus(w, "Combining results")

	// Call LLM3 (Zen Judge) to synthesize the responses
	sendSSEStatus(w, "Resolution forming")

	// Format the prompt for Zen Judge with both arguments (4 arguments: dilemma, buildFast, stillness, history)
	prompt := fmt.Sprintf(llm.ZEN_JUDGE_PROMPT, req.Message, buildFastResp.Argument, stillnessResp.Argument, conversationHistory)

	// Use GPT-4o for Zen Judge
	zenJudgeResponse, err := llmService.GenerateForAgent(ctx, "ZenJudge", prompt)
	if err != nil {
		return "", fmt.Errorf("zen judge error: %v", err)
	}

	return zenJudgeResponse, nil
}

// HandleZenbotRequest processes the /zenbot endpoint
func HandleZenbotRequest(w http.ResponseWriter, r *http.Request) {
	// Parse JSON input
	var req model.ZenbotRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendSSEError(w, "Invalid JSON format")
		return
	}

	// Validate input
	if err := validateRequest(req); err != nil {
		sendSSEError(w, err.Error())
		return
	}

	// Setup conversation ID
	setupConversationID(&req)

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

	// Prepare conversation history (for now, empty - will be enhanced in future tasks)
	conversationHistory := "[]"

	// Orchestrate LLM response
	zenJudgeResponse, err := orchestrateLLMResponse(ctx, w, req, conversationHistory)
	if err != nil {
		sendSSEError(w, err.Error())
		return
	}

	// Stream the final response word by word
	streamMessageProgressively(w, zenJudgeResponse)
}
