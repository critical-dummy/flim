# FLIM - Flash Live Instant Messenger

A high-performance instant messaging server built with Go, gRPC, and MongoDB.

## Features

- [x] User Authentication with JWT
- [x] One-to-One Messaging
- [x] Real-time Bidirectional Streaming (gRPC)
- [x] MongoDB for Message Persistence
- [x] Efficient Protocol Buffer Serialization

## Tech Stack

- **Language**: Go 1.21
- **RPC**: gRPC
- **Database**: MongoDB
- **Authentication**: JWT + bcrypt
- **Serialization**: Protocol Buffers

## Project Structure

```
flim/
├── cmd/
│   └── server/
│       └── main.go              # Server entry point
├── proto/
│   ├── message.proto            # Message definitions
│   └── messenger.proto          # Service definitions
├── internal/
│   ├── auth/
│   │   └── jwt.go               # JWT token management
│   ├── db/
│   │   └── mongo.go             # MongoDB client
│   ├── models/
│   │   ├── user.go              # User model
│   │   └── message.go           # Message model
│   └── server/
│       └── messenger.go         # Messenger service implementation
├── go.mod
└── README.md
```

## Installation

### Prerequisites

- Go 1.21+
- MongoDB 5.0+
- Protocol Buffer Compiler

### Setup

1. Clone the repository:
```bash
git clone https://github.com/critical-dummy/flim.git
cd flim
```

2. Install dependencies:
```bash
go mod download
```

3. Generate Protocol Buffer files:
```bash
protoc --go_out=. --go-grpc_out=. proto/*.proto
```

4. Start MongoDB:
```bash
mongod
```

5. Run the server:
```bash
go run cmd/server/main.go
```

Server will start on `:50051`

## API Endpoints (gRPC)

### Authentication

- `Register(AuthRequest) -> AuthResponse` - Create new user
- `Login(AuthRequest) -> AuthResponse` - User login

### Messaging

- `SendMessage(SendMessageRequest) -> ChatMessage` - Send a message
- `StreamMessages(Empty) -> stream ChatMessage` - Stream incoming messages
- `GetConversation(GetConversationRequest) -> stream ChatMessage` - Get conversation history

## Environment Variables

```bash
MONGODB_URI=mongodb://localhost:27017
JWT_SECRET=your-secret-key
GRPC_PORT=50051
```

## Development

### Testing

```bash
go test ./...
```

### Building

```bash
go build -o bin/flim cmd/server/main.go
```

## TODO

- [ ] Complete gRPC service registration
- [ ] Implement streaming message handlers
- [ ] Add connection management
- [ ] Add message delivery confirmation
- [ ] Add typing indicators
- [ ] Add user presence tracking
- [ ] Add group messaging support
- [ ] Add unit tests
- [ ] Add Docker configuration
- [ ] Add CI/CD pipeline

## License

MIT
