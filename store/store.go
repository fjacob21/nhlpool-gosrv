package store

// Store interface of data storage object
type Store interface {
	Close() error
	Clean() error
	Player() Player
	Session() SessionStore
	League() LeagueStore
	Venue() VenueStore
	Team() TeamStore
	Season() SeasonStore
	Standing() StandingStore
	Game() GameStore
	Matchup() MatchupStore
	Winner() WinnerStore
	Prediction() PredictionStore
	Conference() ConferenceStore
	Division() DivisionStore
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
