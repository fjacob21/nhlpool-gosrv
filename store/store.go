package store

import (
	"nhlpool.com/service/go/nhlpool/data"
)

// Store interface of data storage object
type Store interface {
	Clean()
	GetPlayers() []data.Player
	GetPlayer(id string) *data.Player
	AddPlayer(player *data.Player) error
	UpdatePlayer(player *data.Player) error
	DeletePlayer(player *data.Player) error
	StoreSessions(sessions map[string]data.LoginData)
	LoadSessions() map[string]data.LoginData
}

var store = NewMemoryStore()
var sessionManager = NewSessionManager()

// GetStore Return the current data store
func GetStore() Store {
	return store
}

// GetSessionManager Return the session manager
func GetSessionManager() *SessionManager {
	return sessionManager
}

// Clean Clean the store
func Clean() {
	store.Clean()
	sessionManager.Clean()
}
