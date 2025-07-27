package model

// ZenbotRequest represents the incoming request to the /zenbot endpoint
type ZenbotRequest struct {
	ConversationID string `json:"conversation_id"`
	Message        string `json:"message"`
}

// LLMResponse represents a response from any of the LLMs
type LLMResponse struct {
	Argument string `json:"argument"`
}

// SSEEvent represents a Server-Sent Event
type SSEEvent struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}
