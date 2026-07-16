package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/seuusername/guess-game/internal/handler"
	"github.com/seuusername/guess-game/internal/store"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified in environment variables
	}

	// Create a new store for games
	gameStore := store.NewStore()
	createHandler := handler.NewHandler(gameStore)

	http.HandleFunc("/game", createHandler.CreateGame)
	http.HandleFunc("/game/{id}/guess", createHandler.GuessNumber)
	http.HandleFunc("/game/{id}", createHandler.GetGameStatus)

	// Start the HTTP server
	slog.Info("Server Started", "port", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		slog.Error("Server failed to start", "error", err)
	}
}
