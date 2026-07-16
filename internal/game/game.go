package game

import (
	"strconv"
)

type Game struct {
	ID           string
	SecretNumber int
	AttemptsMade int
	MaxAttempts  int
	CorrectGuess bool
}

// Guess checks if the guessed number matches the secret number. It increments the number of attempts and returns true if the guess is correct, otherwise false. If the game has already been won, it returns true without incrementing attempts.
func (g *Game) Guess(number int) bool {

	if g.CorrectGuess {
		return true
	}
	if g.AttemptsMade >= g.MaxAttempts {
		return false
	}

	g.AttemptsMade++
	if number == g.SecretNumber {
		g.CorrectGuess = true
		return true
	}
	return false
}

// GetHint provides a hint based on the guessed number. It returns a string indicating whether the secret number is higher or lower than the guessed number.
func (g *Game) GetHint(number int) string {

	if number < g.SecretNumber {
		return "No, the secret number is higher."
	}
	return "No, the secret number is lower."

}

// Functios for number of attempts maximum 10 tentatives
func (g *Game) Tentatives() int {
	return g.MaxAttempts - g.AttemptsMade
}

// GetAttemptsMessage returns a message indicating the number of attempts left or if the maximum number of attempts has been reached.
func (g *Game) GetAttemptsMessage() string {
	if g.AttemptsMade >= g.MaxAttempts {
		return "You have reached the maximum number of attempts."
	}
	return "You have " + strconv.Itoa(g.Tentatives()) + " attempts left."
}

// NewGame creates a new game instance with the provided ID and secret number.
func NewGame(id string, secretNumber int) *Game {
	return &Game{
		ID:           id,
		SecretNumber: secretNumber,
		AttemptsMade: 0,
		MaxAttempts:  10,
		CorrectGuess: false,
	}
}
