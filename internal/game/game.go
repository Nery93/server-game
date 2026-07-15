package game

type Game struct {
	ID            string
	SecretNumber  int
	AttemptsMade  int
	CorrectGuess  bool
}

// Guess checks if the guessed number matches the secret number. It increments the number of attempts and returns true if the guess is correct, otherwise false. If the game has already been won, it returns true without incrementing attempts.
func (g *Game) Guess(number int) bool {

	if g.CorrectGuess {
		return true
	}

	g.AttemptsMade++
	if number == g.SecretNumber {
		g.CorrectGuess = true
		return true
	}
	return false
}

func NewGame(id string, secretNumber int) *Game {
	return &Game{
		ID:            id,
		SecretNumber: secretNumber,
		AttemptsMade: 0,
		CorrectGuess: false,
	}
}
