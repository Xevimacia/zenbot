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
git clone https://github.com/Xevimacia/zenbot.git
cd zenbot

# Install dependencies
go mod tidy

# Copy the example environment file and fill in your values
cp .env.example .env
# Set your OpenAI API key in .env
OPENAI_API_KEY=your-openai-api-key

# Run the server
go run cmd/zenbot/main.go
```

The server will start on `http://localhost:8080`

### 🎯 Quick Demo
Test the API immediately with this curl command:

```bash
curl -X POST http://localhost:8080/v1/zenbot \
  -H "Content-Type: application/json" \
  -d '{
    "message": "Should I launch my startup now or wait for more validation?"
  }' \
  --no-buffer
```

Watch the real-time SSE stream with status updates and word-by-word message streaming!

## 📁 Project Structure

```
zenbot/
├── cmd/zenbot/           # Main application entrypoint
│   └── main.go          # HTTP server setup and configuration
├── internal/
│   ├── handler/         # HTTP handlers and SSE streaming
│   │   ├── routes.go    # Route registration and middleware
│   │   ├── zenbot.go    # Main zenbot handler with concurrency
│   │   └── zenbot_test.go # Integration tests for SSE streaming
│   ├── llm/            # OpenAI client and prompt management
│   │   ├── client.go   # LLM service interface and OpenAI client
│   │   ├── client_test.go # Unit tests for LLM service
│   │   └── prompts.go  # LLM prompt constants and cultural alignment
│   └── model/          # Shared data structures
│       └── zenbot.go    # Request/response models and SSE events
├── .gitignore           # Git ignore patterns
├── .env.example         # Environment variables template
├── go.mod              # Go module definition and dependencies
├── go.sum              # Dependency checksums
└── README.md           # Project documentation
```

## 🔌 API Usage

### Endpoint: `POST /v1/zenbot`

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

## 🧪 Step-by-Step Testing Guide

### 1. Start the Server
```bash
# Terminal 1: Start the ZenBot API server
go run cmd/zenbot/main.go
```
You should see: `ZenBot API starting on port 8080`

### 2. Test with curl (Detailed SSE Output)
```bash
# Terminal 2: See the full SSE stream with status updates
curl -X POST http://localhost:8080/v1/zenbot \
  -H "Content-Type: application/json" \
  -d '{
    "conversation_id": "my-team-dilemma-001",
    "message": "Do I need to create a new feature or spend more time with customers to polish the features I have?"
  }' \
  --no-buffer
```

**Expected Output:**
```
event: status
data: Processing dilemma...

event: status
data: Build Fast argues

event: status
data: Stillness reflects

event: status
data: Got response from BuildFast

event: status
data: Got response from Stillness

event: status
data: Combining results

event: status
data: Resolution forming

event: message
data: message_id: zenbot-1753630945, content: To

event: message
data: message_id: zenbot-1753630945, content: To rush

event: message
data: message_id: zenbot-1753630945, content: To rush is

event: message
data: message_id: zenbot-1753630945, content: To rush is to

event: message
data: message_id: zenbot-1753630945, content: To rush is to pause, 🌿

[... continues with word-by-word streaming ...]
```

## 🧠 How It Works

1. **User submits a dilemma** via POST to `/zenbot`
2. **HTTP Method Middleware** validates the request method (POST only)
3. **SSE Middleware** sets up streaming headers
4. **Build Fast (LLM1)** and **Stillness (LLM2)** run concurrently using goroutines
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
- ✅ Input validation (empty message handling)
- ✅ HTTP method middleware (POST only)
- ✅ Progressive message streaming
- ✅ LLM service interface and client creation
- ✅ Environment variable handling
- ✅ Agent-based generation testing
- ✅ Race condition testing (thread safety)


## 🎨 Cultural Alignment

The API reflects Acai Travel's unique culture:

- **Build Fast**: Embodies entrepreneurial drive and opportunity-seeking
- **Stillness**: Reflects Tibetan Buddhist principles of ego-less collaboration
- **Zen Judge**: Synthesizes both perspectives with travel-inspired metaphors

## 🔧 Configuration

### Environment Variables

The application requires the following environment variable:

```bash
# Required
OPENAI_API_KEY=your-openai-api-key
```

You can set this in your shell or use a `.env` file with a library like `godotenv`.

## 🚀 Future Improvements (Production-Ready Enhancements)

### **High Priority - Core Production Features**
- **Persistent Thread History**: Implement SQLite for storing conversation history across sessions, enabling team continuity and context preservation.
- **Health Check Endpoint**: Add `/health` endpoint for monitoring, load balancer integration, and production deployment readiness.
- **Enhanced Logging**: Implement structured logging with correlation IDs for debugging complex SSE streams and LLM interactions.
- **Graceful Shutdown**: Add proper shutdown handling for long-running SSE connections to prevent data loss during deployments.

### **Medium Priority - Operational Excellence**
- **Configuration Management**: Support for different environments (dev, staging, prod) with configurable LLM models, timeouts, and settings.
- **Request Validation**: Add comprehensive input validation with detailed error messages, request sanitization, and rate limiting per team.
- **Metrics & Monitoring**: Integrate basic metrics for request counts, response times, and error rates to monitor system health.
- **SSE Helper Improvements**: Add keep-alive pings and multi-line data support to prevent timeouts in production environments.

### **Nice-to-Have - Team Experience**
- **Meditative Frontend**: Add a simple HTML/CSS frontend with SSE client for direct team interaction without external tools.
- **Caching Layer**: Implement in-memory caching for frequently requested dilemmas to improve response times.
- **Separate Messages Table**: Enhance thread history with proper database schema for better querying and performance.

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