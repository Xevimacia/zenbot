package llm

import (
	"context"
	"testing"

	"github.com/openai/openai-go"
)

// MockLLMService is a minimal mock for testing
type MockLLMService struct{}

func (m *MockLLMService) Generate(ctx context.Context, prompt string, history string) (string, error) {
	return `{"argument": "Mock response"}`, nil
}

// TestNewOpenAIClient tests basic client creation
func TestNewOpenAIClient(t *testing.T) {
	// Test successful creation
	client, err := NewOpenAIClient(openai.ChatModelGPT4_1Mini2025_04_14, "test-key")
	if err != nil {
		t.Errorf("NewOpenAIClient() error: %v", err)
	}
	if client == nil {
		t.Error("NewOpenAIClient() returned nil")
	}

	// Test failure case - empty API key
	_, err = NewOpenAIClient(openai.ChatModelGPT4_1Mini2025_04_14, "")
	if err == nil {
		t.Error("NewOpenAIClient() expected error for empty API key")
	}
}
