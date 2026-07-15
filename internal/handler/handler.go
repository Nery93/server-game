package handler

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/google/uuid"
	"github.com/seuusername/guess-game/internal/game"
	"github.com/seuusername/guess-game/internal/store"
)

type Handler struct {
	store *store.Store // store is a reference to the game storage, allowing the handler to save and retrieve game instances.
}

// CreateGame handles the creation of a new game. It generates a unique ID and a random secret number, saves the game in the store, and responds with the game ID in JSON format.
func (h *Handler) CreateGame(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Generate a unique ID for the game and a random secret number between 0 and 100
	id := uuid.New().String()
	randomNumber := rand.Intn(100 + 1)

	// Create a new game instance and save it in the store
	game := game.NewGame(id, randomNumber)
	h.store.Save(game)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Respond with the game ID in JSON format
	err := json.NewEncoder(w).Encode(map[string]string{"id": id})
	if err != nil {
		http.Error(w, "Error creating game", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) GuessNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the game ID from the URL path
	id := r.PathValue("id")

	// Retrieve the game instance from the store using the extracted ID
	game, ok := h.store.Get(id)
	if !ok {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	var request struct {
		Number int `json:"number"` // Define a struct to parse the guessed number from the request body
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if the guessed number is correct and respond accordingly
	correct := game.Guess(request.Number)
	response := map[string]interface{}{
		"correct":  correct,
		"attempts": game.AttemptsMade,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Error processing guess", http.StatusInternalServerError)
		return
	}
}

// NewHandler creates a new instance of Handler with a reference to the game storage
func NewHandler(store *store.Store) *Handler {
	return &Handler{
		store: store,
	}
}
