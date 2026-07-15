package main

import (
	"fmt"
	"net/http"

	"github.com/seuusername/guess-game/internal/handler"
	"github.com/seuusername/guess-game/internal/store"
)

func main() {

	// Create a new store for games
	gameStore := store.NewStore()
	createHandler := handler.NewHandler(gameStore)
	
	http.HandleFunc("/game", createHandler.CreateGame)
	http.HandleFunc("/game/{id}/guess", createHandler.GuessNumber)

	// Start the HTTP server
	fmt.Println("Servidor iniciado na porta 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}