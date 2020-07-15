package store

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestNewMemoryStoreAddPSession(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	store.Player().AddPlayer(player)
	session := &data.LoginData{SessionID: "id", LoginTime: time.Now(), Player: *player}
	err := store.Session().AddSession(session)
	assert.NoError(t, err, "Should not have error")
	getSession, err := store.Session().GetSession(session.SessionID)
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
	store.Player().AddPlayer(player)
	session := &data.LoginData{SessionID: "id", LoginTime: time.Now(), Player: *player}
	err := store.Session().AddSession(session)
	assert.NoError(t, err, "Should not have error")
	getSession, err := store.Session().GetSession(session.SessionID)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getSession, "Should not be nil")
	err = store.Session().DeleteSession(session)
	assert.Nil(t, err, "Should be nil")
	assert.NoError(t, err, "Should not have error")
	getSession, err = store.Session().GetSession(session.SessionID)
	assert.Error(t, err, "Should have error")
	assert.Nil(t, getSession, "Should be nil")
}
