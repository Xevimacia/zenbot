package llm

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// LLMService abstracts an LLM client capable of generating responses given a prompt and conversation history.
type LLMService interface {
	// Generate produces a model response given a prompt and thread history.
	Generate(ctx context.Context, prompt string, history string) (string, error)
}

// OpenAIClient implements LLMService using the OpenAI API.
type OpenAIClient struct {
	client openai.Client
	model  openai.ChatModel
}

// NewOpenAIClient creates a new OpenAIClient with the given model (e.g., openai.ChatModelGPT4_1Mini2025_04_14).
func NewOpenAIClient(model openai.ChatModel, apiKey string) (*OpenAIClient, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key is required")
	}
	client := openai.NewClient(option.WithAPIKey(apiKey))
	return &OpenAIClient{client: client, model: model}, nil
}

// Generate sends a prompt and history to the OpenAI API and returns the model's response.
func (o *OpenAIClient) Generate(ctx context.Context, prompt string, history string) (string, error) {
	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(prompt),
	}
	if history != "" {
		messages = append(messages, openai.UserMessage(history))
	}
	resp, err := o.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: messages,
		Model:    o.model,
	})
	if err != nil {
		return "", fmt.Errorf("OpenAI API error: %w", err)
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from OpenAI API")
	}
	return resp.Choices[0].Message.Content, nil
}

// ArgumentResponse represents the JSON response structure from LLM1 and LLM2
type ArgumentResponse struct {
	Argument string `json:"argument"`
}

// ParseArgumentResponse parses the JSON response from LLM1 and LLM2
func ParseArgumentResponse(response string) (*ArgumentResponse, error) {
	var argResp ArgumentResponse
	if err := json.Unmarshal([]byte(response), &argResp); err != nil {
		return nil, fmt.Errorf("failed to parse argument response: %w", err)
	}
	return &argResp, nil
}
