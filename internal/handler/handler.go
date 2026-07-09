package handler

import (
	"math/rand"
	"net/http"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/seuusername/guess-game/internal/game"
	"github.com/seuusername/guess-game/internal/store"
)

type Handler struct {
	store *store.Store	// store is a reference to the game storage, allowing the handler to save and retrieve game instances.
}

	// CreateGame handles the creation of a new game. It generates a unique ID and a random secret number, saves the game in the store, and responds with the game ID in JSON format.
func (h *Handler) CreateGame(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Generate a unique ID for the game and a random secret number between 0 and 100
	id := uuid.New().String()
	numeroAleatorio := rand.Intn(100+1)

	// Create a new game instance and save it in the store
	game := game.NewGame(id, numeroAleatorio)
	h.store.Save(game)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Respond with the game ID in JSON format
	err := json.NewEncoder(w).Encode(map[string]string{"id": id})
	if err != nil {
		http.Error(w, "Erro ao criar o jogo", http.StatusInternalServerError)
		return
	}

}

// NewHandler creates a new instance of Handler with a reference to the game storage
func NewHandler(store *store.Store) *Handler { 
	return &Handler{
		store: store,
	}
}