package game

type Game struct {
	ID            string
	NumeroSecreto int
	JogadasFeitas int
	JogadasCertas bool
}

// Guess checks if the guessed number matches the secret number. It increments the number of attempts and returns true if the guess is correct, otherwise false. If the game has already been won, it returns true without incrementing attempts.
func (g *Game) Guess(numero int) bool {

	if g.JogadasCertas {
		return true
	}

	g.JogadasFeitas++
	if numero == g.NumeroSecreto {
		g.JogadasCertas = true
		return true
	}
	return false
}

func NewGame(id string, numeroSecreto int) *Game {
	return &Game{
		ID:            id,
		NumeroSecreto: numeroSecreto,
		JogadasFeitas: 0,
		JogadasCertas: false,
	}
}
