package game

import "testing"

func TestNewGame(t *testing.T) {
	id := "test-id"
	secret := 42

	game := NewGame(id, secret)

	if game.ID != id {
		t.Errorf("Expected ID %s, got %s", id, game.ID)
	}
	if game.SecretNumber != secret {
		t.Errorf("Expected SecretNumber %d, got %d", secret, game.SecretNumber)
	}
	if game.AttemptsMade != 0 {
		t.Errorf("Expected AttemptsMade 0, got %d", game.AttemptsMade)
	}
	if game.CorrectGuess != false {
		t.Errorf("Expected CorrectGuess false, got %v", game.CorrectGuess)
	}
}

func TestGuess_Correct(t *testing.T) {
	game := NewGame("test", 50)

	result := game.Guess(50)

	if !result {
		t.Error("Expected true when guessing correct number")
	}
	if !game.CorrectGuess {
		t.Error("Expected CorrectGuess to be true")
	}
	if game.AttemptsMade != 1 {
		t.Errorf("Expected 1 attempt, got %d", game.AttemptsMade)
	}
}

func TestGuess_Incorrect(t *testing.T) {
	game := NewGame("test", 50)

	result := game.Guess(25)

	if result {
		t.Error("Expected false when guessing incorrect number")
	}
	if game.CorrectGuess {
		t.Error("Expected CorrectGuess to be false")
	}
	if game.AttemptsMade != 1 {
		t.Errorf("Expected 1 attempt, got %d", game.AttemptsMade)
	}
}

func TestGuess_MultipleAttempts(t *testing.T) {
	game := NewGame("test", 75)

	game.Guess(10)
	game.Guess(20)
	game.Guess(30)

	if game.AttemptsMade != 3 {
		t.Errorf("Expected 3 attempts, got %d", game.AttemptsMade)
	}
	if game.CorrectGuess {
		t.Error("Expected CorrectGuess to be false")
	}
}

func TestGuess_AfterWinning(t *testing.T) {
	game := NewGame("test", 42)

	// First guess - correct
	game.Guess(42)
	attemptsAfterWin := game.AttemptsMade

	// Second guess - should not increment
	result := game.Guess(99)

	if !result {
		t.Error("Expected true when game is already won")
	}
	if game.AttemptsMade != attemptsAfterWin {
		t.Errorf("Expected attempts to stay at %d, got %d", attemptsAfterWin, game.AttemptsMade)
	}
}
