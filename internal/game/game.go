package game

type Game struct {
	ID            string
	NumeroSecreto int
	JogadasFeitas int
	JogadasCertas bool
}

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
		JogadasCertas:   false,
	}
}