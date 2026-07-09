package store

import (
	"sync"

	"github.com/seuusername/guess-game/internal/game"
)

type Store struct {
	store map[string]*game.Game // store is a map that holds game instances, where the key is the game ID and the value is a pointer to the Game struct.
	mutex sync.Mutex
}

// Save saves a game instance in the store. It locks the store to ensure thread safety while adding the game to the map.
func (s *Store) Save(game *game.Game) {

	s.mutex.Lock()         // Lock the store to ensure thread safety while saving the game
	defer s.mutex.Unlock() // Unlock the store after saving the game to allow other operations to proceed
	s.store[game.ID] = game

}

// Get retrieves a game instance from the store by its ID. It locks the store to ensure thread safety while accessing the map and returns the game instance and a boolean indicating whether the game was found.
func (s *Store) Get(id string) (*game.Game, bool) {

	s.mutex.Lock()
	defer s.mutex.Unlock()

	game, ok := s.store[id]
	if !ok {
		return nil, false
	}
	return game, ok
}

// NewStore creates a new instance of Store with an initialized map to hold game instances. It returns a pointer to the newly created Store.
func NewStore() *Store {
	return &Store{
		store: make(map[string]*game.Game),
	}
}
