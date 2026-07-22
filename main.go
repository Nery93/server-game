package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	server := &http.Server{
		Addr:    ":" + port,
		Handler: nil, // Use the default HTTP mux
	}

	// Register the handler functions for the endpoints
	http.HandleFunc("/game", createHandler.CreateGame)
	http.HandleFunc("/game/{id}/guess", createHandler.GuessNumber)
	http.HandleFunc("/game/{id}", createHandler.GetGameStatus)

	// Start the HTTP server
	slog.Info("Server Started", "port", port)

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed to start", "error", err)
		}
	}()

	// Wait for a termination signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	slog.Info("Shutdown signal received")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
	} else {
		slog.Info("Server shutdown gracefully")
	}
}
