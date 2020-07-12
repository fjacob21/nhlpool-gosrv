package store

import (
	"testing"
	"time"

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
	players, _ := store.GetPlayers()
	assert.Equal(t, len(players), 0, "There should not have any player")
	store.AddPlayer(player)
	players, _ = store.GetPlayers()
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
	players, _ := store.GetPlayers()
	assert.NotNil(t, players, "Should not be nil")
	assert.Equal(t, len(players), 1, "Should be only one player")
	err := store.DeletePlayer(player)
	assert.Nil(t, err, "Should be nil")
	players, _ = store.GetPlayers()
	assert.Equal(t, len(players), 0, "There should not be any player")
}

func TestNewMemoryStoreDeleteInvalidPlayer(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	err := store.DeletePlayer(player)
	assert.NotNil(t, err, "Should be nil")
}

func TestNewMemoryStoreAddPSession(t *testing.T) {
	store := NewMemoryStore()
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
	assert.Equal(t, getSession.LoginTime, session.LoginTime, "Invalid name")
	assert.NotNil(t, getSession.Player, "Should not be nil")
	assert.Equal(t, getSession.Player.ID, player.ID, "Invalid email")
}

func TestNewMemoryStoreDeleteSession(t *testing.T) {
	store := NewMemoryStore()
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

func TestNewMemoryStoreAddLeaguer(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	league := &data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.AddLeague(league)
	assert.NoError(t, err, "Should not have error")
	getLeague, err := store.GetLeague(league.ID)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getLeague, "Should not be nil")
	assert.Equal(t, getLeague.ID, league.ID, "Invalid ID")
	assert.Equal(t, getLeague.Name, league.Name, "Invalid name")
	assert.Equal(t, getLeague.Description, league.Description, "Invalid description")
	assert.Equal(t, getLeague.Website, league.Website, "Invalid we site")
}

func TestNewMemoryStoreUpdateLeague(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	league := &data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.AddLeague(league)
	assert.NoError(t, err, "Should not have error")
	league.Name = "name2"
	league.Description = "description2"
	league.Website = "website2"
	err = store.UpdateLeague(league)
	assert.Nil(t, err, "Should be nil")
	getLeague, err := store.GetLeague(league.ID)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getLeague, "Should not be nil")
	assert.Equal(t, getLeague.ID, "id", "Invalid ID")
	assert.Equal(t, getLeague.Name, "name2", "Invalid name")
	assert.Equal(t, getLeague.Description, "description2", "Invalid description")
	assert.Equal(t, getLeague.Website, "website2", "Invalid we site")
}

func TestNewMemoryStoreDeleteLeague(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	league := &data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.AddLeague(league)
	assert.NoError(t, err, "Should not have error")
	getLeague, err := store.GetLeague(league.ID)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getLeague, "Should not be nil")
	err = store.DeleteLeague(league)
	assert.NoError(t, err, "Should not have error")
	getLeague, err = store.GetLeague(league.ID)
	assert.Error(t, err, "Should have error")
	assert.Nil(t, getLeague, "Should be nil")
}

func TestNewMemoryStoreGetLeagues(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	league := &data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	leagues, _ := store.GetLeagues()
	assert.Equal(t, len(leagues), 0, "There should not have any leagues")
	err := store.AddLeague(league)
	assert.NoError(t, err, "Should not have error")
	leagues, _ = store.GetLeagues()
	assert.Equal(t, len(leagues), 1, "There should not have any leagues")
	assert.Equal(t, leagues[0].ID, league.ID, "Should be the good league")
}
