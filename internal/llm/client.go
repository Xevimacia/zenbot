package llm

import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// OpenAIClient handles OpenAI API calls with agent-specific model selection
type OpenAIClient struct {
	client openai.Client
}

// NewOpenAIClient creates a new OpenAIClient
func NewOpenAIClient() (*OpenAIClient, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable not set")
	}

	return &OpenAIClient{
		client: openai.NewClient(option.WithAPIKey(apiKey)),
	}, nil
}

// GenerateForAgent generates a response for a specific agent using the appropriate model
func (o *OpenAIClient) GenerateForAgent(ctx context.Context, agent string, prompt string) (string, error) {
	// Agent to model mapping
	var model openai.ChatModel
	switch agent {
	case "BuildFast", "Stillness":
		model = openai.ChatModelGPT4_1Mini2025_04_14
	case "ZenJudge":
		model = openai.ChatModelGPT4o
	default:
		return "", fmt.Errorf("unknown agent: %s", agent)
	}

	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(prompt),
	}

	resp, err := o.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: messages,
		Model:    model,
	})
	if err != nil {
		return "", fmt.Errorf("OpenAI API error: %w", err)
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no choices returned from OpenAI API")
	}
	return resp.Choices[0].Message.Content, nil
}
