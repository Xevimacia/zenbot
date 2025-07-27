# ZenBot API

A mindfulness-based decision-making API for Acai Travel that resolves team dilemmas through a unique dialogue between three AI perspectives: **Build Fast** (entrepreneurial drive), **Stillness** (ego-less reflection), and **Zen Judge** (wise synthesis). Built in Go with OpenAI LLMs and Server-Sent Events (SSE) streaming.

## 🎯 Project Overview

ZenBot embodies Acai Travel's dual nature as entrepreneurial builders and advocates of stillness. When teams face dilemmas, the API orchestrates a conversation between contrasting perspectives, ultimately synthesizing a balanced, professional resolution that aligns with the company's values of curiosity, generosity, and stillness.

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
│   └── main.go          # HTTP server setup
├── internal/
│   ├── handler/         # HTTP handlers and SSE streaming
│   │   ├── routes.go    # Route registration
│   │   └── zenbot.go    # Main zenbot handler
│   ├── llm/            # OpenAI client and prompt management
│   │   ├── client.go   # LLM service interface and OpenAI client
│   │   ├── client_test.go # Unit tests for LLM service
│   │   └── prompts.go  # LLM prompt constants
│   └── model/          # Shared data structures
│       └── zenbot.go    # Request/response models
├── .gitignore           # Git ignore patterns
├── go.mod               # Go module definition
└── README.md
```

## 🔌 API Usage

### Endpoint: `POST /zenbot`

**Request Body:**
```json
{
  "conversation_id": "unique_conversation_id_here",
  "message": "Should we launch the new AI feature now or refine it further?"
}
```

*If `conversation_id` is omitted, a new conversation will be created automatically.*

**Response:** Server-Sent Events (SSE) stream with real-time updates:

```
event: status
data: Processing dilemma...

event: status
data: Build Fast argues

event: status
data: Stillness reflects

event: status
data: Combining results

event: status
data: Resolution forming

event: message
data: message_id: zenbot-1234567890, content: A

event: message
data: message_id: zenbot-1234567890, content: A clear

event: message
data: message_id: zenbot-1234567890, content: A clear path
```

The **Zen Judge**'s final message will be streamed with progressive content updates, showing the message being built up word by word with a slight delay, mimicking human-like typing.

### Example with curl:
```bash
curl -X POST http://localhost:8080/zenbot \
  -H "Content-Type: application/json" \
  -d '{
    "conversation_id": "unique_conversation_id_here",
    "message": "Should we launch the new AI feature now or refine it further?"
  }'
```

## 🧠 How It Works

1. **User submits a dilemma** via POST to `/zenbot`
2. **HTTP Method Middleware** validates the request method (POST only)
3. **Build Fast (LLM1)** generates an action-oriented argument
4. **Stillness (LLM2)** provides a reflective, balanced perspective  
5. **Zen Judge (LLM3)** synthesizes both views into a professional resolution
6. **SSE streaming** delivers real-time status updates and progressive content updates

## 🧪 Testing

Run the test suite:
```bash
# Run all tests
go test ./...

# Run with race detection
go test -race ./...
```

**Current Test Coverage:**
- ✅ Basic project compilation
- ✅ HTTP server startup and response
- ✅ LLM service interface and client creation
- ✅ Mock LLM service for testing

**Planned Test Coverage:**
- 🔄 Handler integration tests (SSE, error cases)
- 🔄 Race detection (thread-safe SSE streaming)

## 🎨 Cultural Alignment

The API reflects Acai Travel's unique culture:

- **Build Fast**: Embodies entrepreneurial drive and opportunity-seeking
- **Stillness**: Reflects Tibetan Buddhist principles of ego-less collaboration
- **Zen Judge**: Synthesizes both perspectives with travel-inspired metaphors

## 🔧 Configuration

### Environment Variables

The application uses a `.env` file for configuration. Copy `.env.example` to `.env` and fill in your values:

```bash
# Required
OPENAI_API_KEY=your-openai-api-key
```

## 🚀 Future Improvements (Optional)
- **Persistent Thread History**: Implement SQLite or similar lightweight database for storing thread history.
- **Separate Messages Table**: Enhance thread history storage by using a messages table instead of a JSON array in threads.
- **SSE Helper Improvements**: Improve support for multi-line data and send multiple data lines per SSE spec.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Submit a pull request

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details.

---

**Built with ❤️ for Acai Travel's mindful decision-making culture**