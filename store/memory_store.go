package store

import (
	"errors"

	"nhlpool.com/service/go/nhlpool/data"
)

// MemoryStore Is a data store that keep it only in memory
type MemoryStore struct {
	players  map[string]*data.Player
	sessions map[string]data.LoginData
}

// NewMemoryStore Create a new memory store
func NewMemoryStore() *MemoryStore {
	store := &MemoryStore{players: make(map[string]*data.Player), sessions: make(map[string]data.LoginData)}
	return store
}

// Clean Empty the store
func (ms *MemoryStore) Clean() {
	ms.players = make(map[string]*data.Player)
	ms.sessions = make(map[string]data.LoginData)
}

// GetPlayers Return a list of all players
func (ms *MemoryStore) GetPlayers() []data.Player {
	var players []data.Player
	for _, player := range ms.players {
		players = append(players, *player)
	}
	return players
}

// GetPlayer Return the player of the specified ID
func (ms *MemoryStore) GetPlayer(id string) *data.Player {
	player, ok := ms.players[id]
	if !ok {
		return nil
	}
	return player
}

// AddPlayer Add a new player
func (ms *MemoryStore) AddPlayer(player *data.Player) error {
	_, ok := ms.players[player.ID]
	if ok {
		return errors.New("Player already exist")
	}
	ms.players[player.ID] = player
	return nil
}

// UpdatePlayer Update a player
func (ms *MemoryStore) UpdatePlayer(player *data.Player) error {
	_, ok := ms.players[player.ID]
	if !ok {
		return errors.New("Player do not exist")
	}
	ms.players[player.ID] = player
	return nil
}

// DeletePlayer Delete a player
func (ms *MemoryStore) DeletePlayer(player *data.Player) error {
	_, ok := ms.players[player.ID]
	if !ok {
		return errors.New("Player do not exist")
	}
	delete(ms.players, player.ID)
	return nil
}

// StoreSessions Store the sessions table
func (ms *MemoryStore) StoreSessions(sessions map[string]data.LoginData) {
	ms.sessions = sessions
}

// LoadSessions Load the sessions table
func (ms *MemoryStore) LoadSessions() map[string]data.LoginData {
	return ms.sessions
}
