package store

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestNewSqliteStore(t *testing.T) {
	store := NewSqliteStore()
	assert.NotNil(t, store, "Should not be nil")
}

func TestNewSqliteStoreAddPlayer(t *testing.T) {
	store := NewSqliteStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	err := store.AddPlayer(player)
	assert.NoError(t, err, "Should not have error")
	getPlayer := store.GetPlayer(player.ID)
	assert.NotNil(t, getPlayer, "Should not be nil")
	assert.Equal(t, getPlayer.ID, player.ID, "Invalid ID")
	assert.Equal(t, getPlayer.Name, "name", "Invalid name")
	assert.Equal(t, getPlayer.Email, "email", "Invalid email")
	assert.Equal(t, getPlayer.Admin, false, "Invalid admin")
	assert.Equal(t, getPlayer.Password, player.Password, "Invalid Password")
	assert.Nil(t, getPlayer.LastLogin, "Should be nil")
}

func TestNewSqliteStoreDoubleAddPlayer(t *testing.T) {
	store := NewSqliteStore()
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

func TestNewSqliteStoreGetPlayers(t *testing.T) {
	store := NewSqliteStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	players, _ := store.GetPlayers()
	assert.Equal(t, len(players), 0, "There should not have any player")
	store.AddPlayer(player)
	players, _ = store.GetPlayers()
	assert.NotNil(t, players, "Should not be nil")
	assert.Equal(t, len(players), 1, "Should be only one player")
	assert.Equal(t, players[0].ID, player.ID, "Should be the good player")
}

func TestNewSqliteStoreUpdatePlayer(t *testing.T) {
	store := NewSqliteStore()
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

func TestNewSqliteStoreUpdateInvalidPlayer(t *testing.T) {
	fmt.Printf("Start\n")
	store := NewSqliteStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	players, _ := store.GetPlayers()
	fmt.Printf("Players %v\n", players)
	err := store.UpdatePlayer(player)
	assert.NotNil(t, err, "Should not be nil")
}

func TestNewSqliteStoreDeletePlayer(t *testing.T) {
	store := NewSqliteStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	store.AddPlayer(player)
	players, _ := store.GetPlayers()
	assert.NotNil(t, players, "Should not be nil")
	assert.Equal(t, len(players), 1, "Should be only one player")
	err := store.DeletePlayer(player)
	assert.Nil(t, err, "Should be nil")
	players, _ = store.GetPlayers()
	assert.Equal(t, len(players), 0, "There should not be any player")
}

func TestNewSqliteStoreDeleteInvalidPlayer(t *testing.T) {
	store := NewSqliteStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	err := store.DeletePlayer(player)
	assert.NotNil(t, err, "Should be nil")
}

func TestNewSqliteStoreAddPSession(t *testing.T) {
	store := NewSqliteStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	store.AddPlayer(player)
	session := &data.LoginData{SessionID: "id", LoginTime: time.Now(), Player: *player}
	err := store.AddSession(session)
	assert.NoError(t, err, "Should not have error")
	getSession, err := store.GetSession(session.SessionID)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getSession, "Should not be nil")
	assert.Equal(t, getSession.SessionID, session.SessionID, "Invalid ID")
	assert.Equal(t, getSession.LoginTime.Unix(), session.LoginTime.Unix(), "Invalid time")
	assert.NotNil(t, getSession.Player, "Should not be nil")
	assert.Equal(t, getSession.Player.ID, player.ID, "Invalid email")
}

func TestNewSqliteStoreDeleteSession(t *testing.T) {
	store := NewSqliteStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	store.AddPlayer(player)
	session := &data.LoginData{SessionID: "id", LoginTime: time.Now(), Player: *player}
	err := store.AddSession(session)
	assert.NoError(t, err, "Should not have error")
	getSession, err := store.GetSession(session.SessionID)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getSession, "Should not be nil")
	err = store.DeleteSession(session)
	assert.Nil(t, err, "Should be nil")
	assert.NoError(t, err, "Should not have error")
	getSession, err = store.GetSession(session.SessionID)
	assert.Error(t, err, "Should have error")
	assert.Nil(t, getSession, "Should be nil")
}
