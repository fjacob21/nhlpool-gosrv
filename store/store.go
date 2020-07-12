package store

import (
	"nhlpool.com/service/go/nhlpool/data"
)

// Store interface of data storage object
type Store interface {
	Clean() error
	GetPlayers() ([]data.Player, error)
	GetPlayer(id string) *data.Player
	AddPlayer(player *data.Player) error
	UpdatePlayer(player *data.Player) error
	DeletePlayer(player *data.Player) error
	AddSession(session *data.LoginData) error
	DeleteSession(session *data.LoginData) error
	GetSession(sessionID string) (*data.LoginData, error)
	GetSessionByPlayer(player *data.Player) (*data.LoginData, error)
	AddLeague(league *data.League) error
	UpdateLeague(league *data.League) error
	DeleteLeague(league *data.League) error
	GetLeague(leagueID string) (*data.League, error)
	GetLeagues() ([]data.League, error)
}

var activeStore = NewMemoryStore()
var sessionManager = NewSessionManager()

// SetStore Return the current data store
func SetStore(store Store) {
	activeStore = store
}

// GetStore Return the current data store
func GetStore() Store {
	return activeStore
}

// GetSessionManager Return the session manager
func GetSessionManager() *SessionManager {
	return sessionManager
}

// Clean Clean the store
func Clean() {
	activeStore.Clean()
}
