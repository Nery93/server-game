package store

import (
	"github.com/seuusername/guess-game/internal/game"
)

type Store struct {
	store map[string]*game.Game
}

func (s *Store) Save(game *game.Game) {
	s.store[game.ID] = game
}

func NewStore() *Store {
	return &Store{
		store: make(map[string]*game.Game),
	}
}