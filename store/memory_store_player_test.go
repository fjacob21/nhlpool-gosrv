package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestNewMemoryStoreAddPlayer(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	store.Player().AddPlayer(player)
	getPlayer := store.Player().GetPlayer(player.ID)
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
	err := store.Player().AddPlayer(player)
	assert.Nil(t, err, "Should be nil")
	getPlayer := store.Player().GetPlayer(player.ID)
	assert.NotNil(t, getPlayer, "Should not be nil")
	assert.Equal(t, getPlayer.ID, player.ID, "Invalid ID")
	assert.Equal(t, getPlayer.Name, "name", "Invalid name")
	assert.Equal(t, getPlayer.Email, "email", "Invalid email")
	assert.Equal(t, getPlayer.Admin, false, "Invalid admin")
	assert.Equal(t, getPlayer.Password, player.Password, "Invalid Password")
	assert.Nil(t, getPlayer.LastLogin, "Should be nil")
	err = store.Player().AddPlayer(player)
	assert.NotNil(t, err, "Should not be nil")
}

func TestNewMemoryStoreGetPlayers(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	players, _ := store.Player().GetPlayers()
	assert.Equal(t, len(players), 0, "There should not have any player")
	store.Player().AddPlayer(player)
	players, _ = store.Player().GetPlayers()
	assert.NotNil(t, players, "Should not be nil")
	assert.Equal(t, len(players), 1, "Should be only one player")
	assert.Equal(t, players[0].ID, player.ID, "Should be the good player")
}

func TestNewMemoryStoreUpdatePlayer(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	store.Player().AddPlayer(player)
	player.Name = "name2"
	err := store.Player().UpdatePlayer(player)
	assert.Nil(t, err, "Should be nil")
	getPlayer := store.Player().GetPlayer(player.ID)
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
	err := store.Player().UpdatePlayer(player)
	assert.NotNil(t, err, "Should be nil")
}

func TestNewMemoryStoreDeletePlayer(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	store.Player().AddPlayer(player)
	players, _ := store.Player().GetPlayers()
	assert.NotNil(t, players, "Should not be nil")
	assert.Equal(t, len(players), 1, "Should be only one player")
	err := store.Player().DeletePlayer(player)
	assert.Nil(t, err, "Should be nil")
	players, _ = store.Player().GetPlayers()
	assert.Equal(t, len(players), 0, "There should not be any player")
}

func TestNewMemoryStoreDeleteInvalidPlayer(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	err := store.Player().DeletePlayer(player)
	assert.NotNil(t, err, "Should be nil")
}
