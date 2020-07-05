package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestNewMemoryStore(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
}

func TestNewMemoryStoreAddPlayer(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	store.AddPlayer(player)
	getPlayer := store.GetPlayer(player.ID)
	assert.NotNil(t, getPlayer, "Should not be nil")
	assert.Equal(t, getPlayer.ID, player.ID, "Invalid ID")
	assert.Equal(t, getPlayer.Name, "name", "Invalid name")
	assert.Equal(t, getPlayer.Email, "email", "Invalid email")
	assert.Equal(t, getPlayer.Admin, false, "Invalid admin")
	assert.Equal(t, getPlayer.Password, player.Password, "Invalid Password")
	assert.Nil(t, getPlayer.LastLogin, "Should be nil")
}

func TestNewMemoryStoreDoubleAddPlayer(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	err := store.AddPlayer(player)
	assert.Nil(t, err, "Should be nil")
	getPlayer := store.GetPlayer(player.ID)
	assert.NotNil(t, getPlayer, "Should not be nil")
	assert.Equal(t, getPlayer.ID, player.ID, "Invalid ID")
	assert.Equal(t, getPlayer.Name, "name", "Invalid name")
	assert.Equal(t, getPlayer.Email, "email", "Invalid email")
	assert.Equal(t, getPlayer.Admin, false, "Invalid admin")
	assert.Equal(t, getPlayer.Password, player.Password, "Invalid Password")
	assert.Nil(t, getPlayer.LastLogin, "Should be nil")
	err = store.AddPlayer(player)
	assert.NotNil(t, err, "Should not be nil")
}

func TestNewMemoryStoreGetPlayers(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	players := store.GetPlayers()
	assert.Equal(t, len(players), 0, "There should not have any player")
	store.AddPlayer(player)
	players = store.GetPlayers()
	assert.NotNil(t, players, "Should not be nil")
	assert.Equal(t, len(players), 1, "Should be only one player")
	assert.Equal(t, players[0].ID, player.ID, "Should be the good player")
}

func TestNewMemoryStoreUpdatePlayer(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	store.AddPlayer(player)
	player.Name = "name2"
	err := store.UpdatePlayer(player)
	assert.Nil(t, err, "Should be nil")
	getPlayer := store.GetPlayer(player.ID)
	assert.NotNil(t, getPlayer, "Should not be nil")
	assert.Equal(t, getPlayer.ID, player.ID, "Invalid ID")
	assert.Equal(t, getPlayer.Name, "name2", "Invalid name")
	assert.Equal(t, getPlayer.Email, "email", "Invalid email")
	assert.Equal(t, getPlayer.Admin, false, "Invalid admin")
	assert.Equal(t, getPlayer.Password, player.Password, "Invalid Password")
	assert.Nil(t, getPlayer.LastLogin, "Should be nil")
}

func TestNewMemoryStoreUpdateInvalidPlayer(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	err := store.UpdatePlayer(player)
	assert.NotNil(t, err, "Should be nil")
}

func TestNewMemoryStoreDeletePlayer(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	store.AddPlayer(player)
	players := store.GetPlayers()
	assert.NotNil(t, players, "Should not be nil")
	assert.Equal(t, len(players), 1, "Should be only one player")
	err := store.DeletePlayer(player)
	assert.Nil(t, err, "Should be nil")
	players = store.GetPlayers()
	assert.Equal(t, len(players), 0, "There should not be any player")
}

func TestNewMemoryStoreDeleteInvalidPlayer(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	err := store.DeletePlayer(player)
	assert.NotNil(t, err, "Should be nil")
}
