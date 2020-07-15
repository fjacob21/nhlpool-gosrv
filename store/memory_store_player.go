package store

import (
	"errors"

	"nhlpool.com/service/go/nhlpool/data"
)

// MemoryStorePlayer Is a Player data store that keep it only in memory
type MemoryStorePlayer struct {
	players map[string]*data.Player
}

// NewMemoryStorePlayer Create a new player memory store
func NewMemoryStorePlayer() *MemoryStorePlayer {
	store := &MemoryStorePlayer{}
	store.players = make(map[string]*data.Player)
	return store
}

// Clean Empty the store
func (ms *MemoryStorePlayer) Clean() error {
	ms.players = make(map[string]*data.Player)
	return nil
}

// GetPlayers Return a list of all players
func (ms *MemoryStorePlayer) GetPlayers() ([]data.Player, error) {
	var players []data.Player
	for _, player := range ms.players {
		players = append(players, *player)
	}
	return players, nil
}

func (ms *MemoryStorePlayer) getPlayerByName(name string) *data.Player {
	for _, player := range ms.players {
		if player.Name == name {
			return player
		}
	}
	return nil
}

// GetPlayer Return the player of the specified ID
func (ms *MemoryStorePlayer) GetPlayer(id string) *data.Player {
	player, ok := ms.players[id]
	if !ok {
		return ms.getPlayerByName(id)
	}
	return player
}

// AddPlayer Add a new player
func (ms *MemoryStorePlayer) AddPlayer(player *data.Player) error {
	_, ok := ms.players[player.ID]
	if ok {
		return errors.New("Player already exist")
	}
	ms.players[player.ID] = player
	return nil
}

// UpdatePlayer Update a player
func (ms *MemoryStorePlayer) UpdatePlayer(player *data.Player) error {
	_, ok := ms.players[player.ID]
	if !ok {
		return errors.New("Player do not exist")
	}
	ms.players[player.ID] = player
	return nil
}

// DeletePlayer Delete a player
func (ms *MemoryStorePlayer) DeletePlayer(player *data.Player) error {
	_, ok := ms.players[player.ID]
	if !ok {
		return errors.New("Player do not exist")
	}
	delete(ms.players, player.ID)
	return nil
}
