package store

import (
	"testing"

	"github.com/seuusername/guess-game/internal/game"
)

func TestNewStore(t *testing.T) {
	store := NewStore()

	if store == nil {
		t.Fatal("Expected store to be created")
	}
	if store.store == nil {
		t.Fatal("Expected internal map to be initialized")
	}
}

func TestStore_SaveAndGet(t *testing.T) {
	store := NewStore()
	g := game.NewGame("test-id", 42)

	store.Save(g)

	retrieved, ok := store.Get("test-id")
	if !ok {
		t.Fatal("Expected to find saved game")
	}
	if retrieved.ID != "test-id" {
		t.Errorf("Expected ID test-id, got %s", retrieved.ID)
	}
	if retrieved.SecretNumber != 42 {
		t.Errorf("Expected SecretNumber 42, got %d", retrieved.SecretNumber)
	}
}

func TestStore_GetNonExistent(t *testing.T) {
	store := NewStore()

	_, ok := store.Get("non-existent")
	if ok {
		t.Error("Expected false when getting non-existent game")
	}
}

func TestStore_SaveMultipleGames(t *testing.T) {
	store := NewStore()
	g1 := game.NewGame("id-1", 10)
	g2 := game.NewGame("id-2", 20)
	g3 := game.NewGame("id-3", 30)

	store.Save(g1)
	store.Save(g2)
	store.Save(g3)

	retrieved1, ok1 := store.Get("id-1")
	retrieved2, ok2 := store.Get("id-2")
	retrieved3, ok3 := store.Get("id-3")

	if !ok1 || !ok2 || !ok3 {
		t.Fatal("Expected all games to be found")
	}
	if retrieved1.SecretNumber != 10 || retrieved2.SecretNumber != 20 || retrieved3.SecretNumber != 30 {
		t.Error("Games not stored correctly")
	}
}

func TestStore_ModifyGameAfterSave(t *testing.T) {
	store := NewStore()
	g := game.NewGame("test-id", 50)

	store.Save(g)

	// Modify the game
	g.Guess(25)
	g.Guess(50)

	// Retrieve and check if changes persisted
	retrieved, _ := store.Get("test-id")
	if retrieved.AttemptsMade != 2 {
		t.Errorf("Expected 2 attempts, got %d", retrieved.AttemptsMade)
	}
	if !retrieved.CorrectGuess {
		t.Error("Expected CorrectGuess to be true")
	}
}
