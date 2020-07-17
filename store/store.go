package store

// Store interface of data storage object
type Store interface {
	Clean() error
	Player() Player
	Session() SessionStore
	League() LeagueStore
	Venue() VenueStore
	Team() TeamStore
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
