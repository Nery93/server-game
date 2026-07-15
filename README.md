# 🎮 Guess Game Server

REST API in Go for a number guessing game. Multiple players can create independent matches and try to guess the secret number.

## 🚀 Technologies

- **Go 1.23+** - Programming language
- **net/http** - Native HTTP server
- **Docker & Docker Compose** - Containerization
- **UUID** - Unique ID generation

## 📋 Features

- ✅ Game creation with random numbers (0-100)
- ✅ Guess system with attempt counter
- ✅ Game state query
- ✅ Complete isolation between matches
- ✅ Thread-safe with Mutex

## 🏗️ Project Structure

```
/
├── main.go              # Server entry point
├── Dockerfile           # Container configuration
├── docker-compose.yml   # Orchestration
└── internal/
    ├── game/           # Game logic
    ├── handler/        # HTTP routes
    └── store/          # In-memory storage
```

## 🎯 How to Run

### Option 1: With Docker (Recommended)

```bash
docker compose up --build
```

### Option 2: Direct with Go

```bash
go run main.go
```

The server will be available at `http://localhost:8080`

## 📡 Endpoints

### Create a new game
```bash
POST /game
```

**Response:**
```json
{
  "id": "abc-123-xyz"
}
```

### Make a guess
```bash
POST /game/{id}/guess
Content-Type: application/json

{
  "number": 42
}
```

**Response:**
```json
{
  "correct": false,
  "attempts": 5
}
```

### Query game state
```bash
GET /game/{id}
```

**Response:**
```json
{
  "correct": false,
  "attempts": 5
}
```

## 🧪 Usage Example

```bash
# 1. Create game
curl -X POST http://localhost:8080/game

# 2. Make a guess (replace {id} with returned ID)
curl -X POST http://localhost:8080/game/{id}/guess \
  -H "Content-Type: application/json" \
  -d '{"number": 50}'

# 3. Check state
curl http://localhost:8080/game/{id}
```

## 🛠️ Development

```bash
# Install dependencies
go mod download

# Run tests (when implemented)
go test ./...

# Build
go build -o server .
```

## 📝 License

MIT
