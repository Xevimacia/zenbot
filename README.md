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
├── internal/
│   ├── llm/            # OpenAI client and prompt management
│   ├── handler/         # HTTP handlers and SSE streaming
│   └── model/          # Shared data structures
└── README.md
```

## 🔌 API Usage

### Endpoint: `POST /zenbot`

**Request Body:**
```json
{
  "conversation_id": "unique_conversation_id_here",
  "message": {
    "role": "user",
    "content": "Should we launch the new AI feature now or refine it further?"
  }
}
```

*If `conversation_id` is omitted, a new conversation will be created.*

**Response:** Server-Sent Events (SSE) stream with real-time updates:

```
event: status
data: Build Fast argues

event: status  
data: Stillness reflects

event: status
data: Resolution forming

event: message
data: **Resolution:** A clear path forward balances ambition with wisdom. [Word-by-word streaming is enabled]
```

The **Zen Judge**'s final message will be streamed word by word to provide a more natural flow. Each word will appear with a slight delay, mimicking human-like typing.

### Example with curl:
```bash
curl -X POST http://localhost:8080/zenbot   -H "Content-Type: application/json"   -d '{
    "conversation_id": "unique_conversation_id_here",
    "message": "Should we launch the new AI feature now or refine it further?"
  }'
```

## 🧠 How It Works

1. **User submits a dilemma** via POST to `/zenbot`
2. **Build Fast (LLM1)** generates an action-oriented argument
3. **Stillness (LLM2)** provides a reflective, balanced perspective  
4. **Zen Judge (LLM3)** synthesizes both views into a professional resolution
5. **SSE streaming** delivers real-time status updates and final response

## 🧪 Testing

Run the test suite:
```bash
# Run all tests
go test ./...

# Run with race detection
go test -race ./...
```

**Test Coverage:**
- ✅ LLM service interface and error handling
- ✅ Handler integration tests (SSE, error cases)
- ✅ Race detection (thread-safe SSE streaming)

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