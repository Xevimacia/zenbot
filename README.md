# ZenBot API

A mindfulness-based decision-making API that resolves team dilemmas through a unique dialogue between three AI perspectives: **Build Fast** (entrepreneurial drive), **Stillness** (ego-less reflection), and **Zen Judge** (wise synthesis). Built in Go with OpenAI LLMs, SQLite, and Server-Sent Events (SSE) streaming.

## 🎯 Project Overview

ZenBot embodies the dual nature of entrepreneurial builders and advocates of stillness. When teams face dilemmas, the API orchestrates a conversation between contrasting perspectives, ultimately synthesizing a balanced, professional resolution that aligns with values of curiosity, generosity, and stillness.

## 🚀 Quick Start

### Prerequisites
- Go 1.20+
- OpenAI API key

### Installation
```bash
# Clone the repository
git clone https://github.com/xevimacia/zenbot.git
cd zenbot

# Install dependencies
go mod tidy

# Copy the example environment file and fill in your values
cp .env.example .env
# Edit .env and add your OpenAI API key

# Run the server
go run cmd/zenbot/main.go
```

The server will start on `http://localhost:8080`

## 📁 Project Structure

```
zenbot/
├── cmd/zenbot/           # Main application entrypoint
├── internal/
│   ├── db/              # SQLite database operations
│   ├── handler/         # HTTP handlers and SSE streaming
│   ├── llm/            # OpenAI client and prompt management
│   └── model/          # Shared data structures
├── db/                 # SQLite database files
└── README.md
```

## 🔌 API Usage

### Endpoint: `POST /zenbot`

**Request Body:**
```json
{
  "user_id": "user123",
  "thread_id": "thread_001", 
  "message": "Should we launch the new AI feature now or refine it further?"
}
```

*If `thread_id` is omitted, a new thread will be created and its ID returned as an SSE event.*

**Response:** Server-Sent Events (SSE) stream with real-time updates:

```
event: status
data: Build Fast argues

event: status  
data: Stillness reflects

event: status
data: Resolution forming

event: message
data: **Resolution:** A clear path forward balances ambition with wisdom. Launch the core AI feature this week to seize opportunities, while scheduling a refinement phase next week to ensure quality and team alignment, fostering curiosity and generosity in our approach.
```

### Example with curl:
```bash
curl -X POST http://localhost:8080/zenbot \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "user123",
    "thread_id": "thread_001", 
    "message": "Should we launch the new AI feature now or refine it further?"
  }'
```

## 🧠 How It Works

1. **User submits a dilemma** via POST to `/zenbot`
2. **Build Fast (LLM1)** generates an action-oriented argument
3. **Stillness (LLM2)** provides a reflective, balanced perspective  
4. **Zen Judge (LLM3)** synthesizes both views into a professional resolution
5. **SSE streaming** delivers real-time status updates and final response
6. **Thread history** is stored in SQLite for conversation continuity
7. **Only the last 5 messages are kept in the thread history** to ensure efficient context for LLMs

## 🧪 Testing

Run the test suite:
```bash
# Run all tests
go test ./...

# Run specific package tests
go test ./internal/db -v
go test ./internal/llm -v
go test ./internal/handler -v

# Run with race detection
go test -race ./...
```

**Test Coverage:**
- ✅ Database operations (thread storage/retrieval)
- ✅ LLM service interface and error handling
- ✅ Handler integration tests (SSE, error cases, thread continuity)
- ✅ Race detection (thread-safe SSE streaming)
- ✅ Environment variable validation

All tests pass and the codebase is safe for concurrent use. The project is ready for review and production use.

## 🏗️ Technical Design

### Concurrency
- Uses Go goroutines and channels for parallel LLM calls
- `sync.WaitGroup` synchronizes responses from Build Fast and Stillness
- Context cancellation handles client disconnections

### Database
- **SQLite** for lightweight, serverless storage
- Thread history stored as a JSON array of messages, each with a `role` (e.g., "user", "zen_judge") and `content`
- Only the last 5 messages are kept in the thread history for context and token efficiency
- Simple CRUD operations for conversation continuity

### LLM Integration
- **Three distinct OpenAI clients** with specialized prompts
- **Interface-based design** for easy testing and mocking
- **Error handling** with user-friendly SSE error messages

### SSE Streaming
- Real-time status updates during processing
- Line-by-line streaming of final responses
- Proper connection handling and error propagation

## 🎨 Cultural Alignment

The API reflects a unique culture:

- **Build Fast**: Embodies entrepreneurial drive and opportunity-seeking
- **Stillness**: Reflects Tibetan Buddhist principles of ego-less collaboration
- **Zen Judge**: Synthesizes both perspectives with travel-inspired metaphors

All prompts are stored in `internal/llm/prompts.go` and can be customized to align with your organization's values.

## 🔧 Configuration

### Environment Variables

The application uses a `.env` file for configuration. Copy `.env.example` to `.env` and fill in your values:

```bash
# Required
OPENAI_API_KEY=your-openai-api-key
```

**Security Note:** The `.env` file is gitignored to prevent accidentally committing secrets. Never commit your actual API keys to version control.

### Database
- SQLite database automatically created at `./db/zenbot.db`
- Tables: `users`, `threads` with conversation history
- Thread history is a JSON array of objects, each with a `role` and `content`, e.g.:
  ```json
  [
    {"role": "user", "content": "Should we launch the new AI feature now or refine it further?"},
    {"role": "zen_judge", "content": "A clear path forward balances ambition with wisdom..."}
  ]
  ```

## 🚀 Future Enhancements

- [ ] **Enhance testing:** Expand unit, integration, coverage, and race tests for even greater reliability and confidence.
- [ ] **Slack Integration**: Connect ZenBot to Slack for seamless team dilemma resolution within your workflow.
- [ ] **Meditative Frontend**: HTML/CSS with SSE client for enhanced UX
- [ ] **Rate Limiting**: Multi-user support and API protection
- [ ] **Enhanced Prompts**: More organization-specific cultural metaphors
- [ ] **Thread History**: Separate messages table for better scalability

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Submit a pull request

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details.

---

**Built with ❤️ for mindful decision-making culture**