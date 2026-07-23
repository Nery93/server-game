# 🎮 Guess Game Server

REST API in Go for a number guessing game (Mastermind/Wordle-style: guess 0–100, get a higher/lower hint, limited attempts). Multiple players can create independent matches with full isolation between them. Includes a React frontend ("Neon Arcade") that consumes the API.

## 🚀 Technologies

**Backend**
- **Go 1.25+** - Programming language
- **net/http** - Native HTTP server
- **log/slog** - Structured logging
- **Docker & Docker Compose** - Containerization
- **UUID** - Unique ID generation
- **GitHub Actions** - CI (runs `go test` on every push/PR)

**Frontend**
- **React 19 + TypeScript**
- **Vite** - Dev server and build tool

## 📋 Features

- ✅ Game creation with random numbers (0-100)
- ✅ Guess validation, attempt counter and higher/lower hints
- ✅ Configurable attempt limit (10), game ends on win or exhausted attempts
- ✅ Game state query
- ✅ Complete isolation between matches, thread-safe with Mutex
- ✅ Configurable port via `PORT` env var
- ✅ Graceful shutdown on `SIGINT`/`SIGTERM`
- ✅ Playable frontend (Neon Arcade theme)

## 🏗️ Project Structure

```
/
├── main.go                    # Server entry point, graceful shutdown
├── Dockerfile                 # Container configuration
├── docker-compose.yml         # Orchestration
├── .github/workflows/         # CI (go test)
├── internal/
│   ├── game/                  # Game logic (guess, hints, attempts)
│   ├── handler/                # HTTP routes
│   └── store/                  # In-memory storage
└── frontend/                  # React + TypeScript + Vite UI
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

The server listens on `http://localhost:8080` by default (override with `PORT=3000 go run main.go`).

### Frontend (dev mode)

```bash
cd frontend
npm install
npm run dev
```

Opens on `http://localhost:5173` and proxies API calls to the Go server on port 8080 — run both at the same time to play.

## 📡 Endpoints

### Create a new game
```bash
POST /game
```

**Response:**
```json
{ "id": "abc-123-xyz" }
```

### Make a guess
```bash
POST /game/{id}/guess
Content-Type: application/json

{ "number": 42 }
```

**Response (wrong guess, game continues):**
```json
{
  "correct": false,
  "attempts": 5,
  "hint": "No, the secret number is higher.",
  "attempts_left": "You have 5 attempts left."
}
```

**Response (correct guess or attempts exhausted):** `secret_number` is included only once the match ends.
```json
{
  "correct": true,
  "attempts": 6,
  "attempts_left": "You have 4 attempts left.",
  "secret_number": 63
}
```

Guesses outside 0–100 return `400`; an unknown game id returns `404`.

### Query game state
```bash
GET /game/{id}
```

**Response:**
```json
{ "correct": false, "attempts": 5 }
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

**Backend**
```bash
go mod download   # install dependencies
go test ./...     # run tests
go build -o server .
```

**Frontend**
```bash
cd frontend
npm run build     # type-check + production build
npm run lint      # oxlint
```

## 📝 License

MIT