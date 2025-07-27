package llm

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	// Load .env file for tests
	godotenv.Load()
}

// MockLLMService is a minimal mock for testing
type MockLLMService struct{}

func (m *MockLLMService) GenerateForAgent(ctx context.Context, agent string, prompt string) (string, error) {
	return `{"argument": "Mock response"}`, nil
}

// TestNewOpenAIClient tests basic client creation
func TestNewOpenAIClient(t *testing.T) {
	// Check if API key is available
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		t.Skip("OPENAI_API_KEY not set, skipping real API tests")
	}

	// Test successful creation
	client, err := NewOpenAIClient()
	if err != nil {
		t.Errorf("NewOpenAIClient() error: %v", err)
	}
	if client == nil {
		t.Error("NewOpenAIClient() returned nil")
	}
}

// TestGenerateForAgent tests the agent-based generation
func TestGenerateForAgent(t *testing.T) {
	// Check if API key is available
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		t.Skip("OPENAI_API_KEY not set, skipping real API tests")
	}

	client, err := NewOpenAIClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Test with each agent type
	agents := []string{"BuildFast", "Stillness", "ZenJudge"}

	for _, agent := range agents {
		response, err := client.GenerateForAgent(context.Background(), agent, "test prompt")
		if err != nil {
			t.Errorf("GenerateForAgent(%s) error: %v", agent, err)
		}
		if response == "" {
			t.Errorf("GenerateForAgent(%s) returned empty response", agent)
		}
	}
}

// TestEnvironmentVariable tests if the API key is loaded
func TestEnvironmentVariable(t *testing.T) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		t.Skip("OPENAI_API_KEY not set, skipping real API tests")
	}
	t.Logf("API key loaded successfully (length: %d)", len(apiKey))
}
