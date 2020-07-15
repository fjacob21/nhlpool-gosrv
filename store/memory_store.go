package store

// MemoryStore Is a data store that keep it only in memory
type MemoryStore struct {
	player  *MemoryStorePlayer
	session *MemoryStoreSession
	league  *MemoryStoreLeague
	venue   *MemoryStoreVenue
}

// NewMemoryStore Create a new memory store
func NewMemoryStore() Store {
	store := &MemoryStore{}
	store.player = NewMemoryStorePlayer()
	store.session = NewMemoryStoreSession()
	store.league = NewMemoryStoreLeague()
	store.venue = NewMemoryStoreVenue()
	return store
}

// Clean Empty the store
func (ms *MemoryStore) Clean() error {
	ms.player.Clean()
	ms.session.Clean()
	ms.league.Clean()
	ms.venue.Clean()
	return nil
}

// Player Return the player store
func (ms *MemoryStore) Player() Player {
	return ms.player
}

// Session Return the session store
func (ms *MemoryStore) Session() SessionStore {
	return ms.session
}

// League Return the league store
func (ms *MemoryStore) League() LeagueStore {
	return ms.league
}

// Venue Return the venue store
func (ms *MemoryStore) Venue() VenueStore {
	return ms.venue
}
