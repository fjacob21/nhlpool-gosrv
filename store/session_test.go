package store

import (
	"testing"

	"nhlpool.com/service/go/nhlpool/data"

	"github.com/stretchr/testify/assert"
)

func TestNewSession(t *testing.T) {
	store.Clean()
	sessionManager := NewSessionManager()
	assert.NotNil(t, sessionManager, "Should not be nil")
}

func TestSessionLogin(t *testing.T) {
	store.Clean()
	sessionManager := NewSessionManager()
	assert.NotNil(t, sessionManager, "Should not be nil")
	player := data.CreatePlayer("name", "email", true, "password")
	sessionID := sessionManager.Login(player, "password")
	assert.NotEqual(t, sessionID, "", "Invalid session ID")
}

func TestSessionDoubleLogin(t *testing.T) {
	store.Clean()
	sessionManager := NewSessionManager()
	assert.NotNil(t, sessionManager, "Should not be nil")
	player := data.CreatePlayer("name", "email", true, "password")
	sessionID := sessionManager.Login(player, "password")
	sessionID2 := sessionManager.Login(player, "password")
	assert.NotEqual(t, sessionID, "", "Invalid session ID")
	assert.Equal(t, sessionID, sessionID2, "Session ID need to be the same")
}

func TestSessionLogout(t *testing.T) {
	store.Clean()
	sessionManager := NewSessionManager()
	assert.NotNil(t, sessionManager, "Should not be nil")
	player := data.CreatePlayer("name", "email", true, "password")
	sessionID := sessionManager.Login(player, "password")
	assert.NotEqual(t, sessionID, "", "Invalid session ID")
	err := sessionManager.Logout(sessionID)
	assert.Nil(t, err, "There should not have an error")
}

func TestSessionInvalidLogout(t *testing.T) {
	store.Clean()
	sessionManager := NewSessionManager()
	assert.NotNil(t, sessionManager, "Should not be nil")
	err := sessionManager.Logout("badessionid")
	assert.NotNil(t, err, "There should have an error")
}

func TestSessionGet(t *testing.T) {
	store.Clean()
	sessionManager := NewSessionManager()
	assert.NotNil(t, sessionManager, "Should not be nil")
	player := data.CreatePlayer("name", "email", true, "password")
	sessionID := sessionManager.Login(player, "password")
	assert.NotEqual(t, sessionID, "", "Invalid session ID")
	loginData := sessionManager.Get(sessionID)
	assert.NotNil(t, loginData, "Should not be nil")
}

func TestSessionInvalidGet(t *testing.T) {
	store.Clean()
	sessionManager := NewSessionManager()
	assert.NotNil(t, sessionManager, "Should not be nil")
	loginData := sessionManager.Get("badessionid")
	assert.Nil(t, loginData, "Should be nil")
}

func TestSessionEmptyGet(t *testing.T) {
	store.Clean()
	sessionManager := NewSessionManager()
	assert.NotNil(t, sessionManager, "Should not be nil")
	loginData := sessionManager.Get("")
	assert.Nil(t, loginData, "Should be nil")
}
