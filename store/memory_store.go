package store

// MemoryStore Is a data store that keep it only in memory
type MemoryStore struct {
	player     *MemoryStorePlayer
	session    *MemoryStoreSession
	league     *MemoryStoreLeague
	venue      *MemoryStoreVenue
	team       *MemoryStoreTeam
	season     *MemoryStoreSeason
	standing   *MemoryStoreStanding
	game       *MemoryStoreGame
	matchup    *MemoryStoreMatchup
	winner     *MemoryStoreWinner
	prediction *MemoryStorePrediction
}

// NewMemoryStore Create a new memory store
func NewMemoryStore() Store {
	store := &MemoryStore{}
	store.player = NewMemoryStorePlayer()
	store.session = NewMemoryStoreSession()
	store.league = NewMemoryStoreLeague()
	store.venue = NewMemoryStoreVenue()
	store.team = NewMemoryStoreTeam()
	store.season = NewMemoryStoreSeason()
	store.standing = NewMemoryStoreStanding()
	store.game = NewMemoryStoreGame()
	store.matchup = NewMemoryStoreMatchup(store)
	store.winner = NewMemoryStoreWinner()
	store.prediction = NewMemoryStorePrediction()
	return store
}

// Close the database
func (ms *MemoryStore) Close() error {
	ms.Clean()
	return nil
}

// Clean Empty the store
func (ms *MemoryStore) Clean() error {
	ms.player.Clean()
	ms.session.Clean()
	ms.league.Clean()
	ms.venue.Clean()
	ms.team.Clean()
	ms.season.Clean()
	ms.standing.Clean()
	ms.game.Clean()
	ms.matchup.Clean()
	ms.winner.Clean()
	ms.prediction.Clean()
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

// Team Return the team store
func (ms *MemoryStore) Team() TeamStore {
	return ms.team
}

// Season Return the season store
func (ms *MemoryStore) Season() SeasonStore {
	return ms.season
}

// Standing Return the standing store
func (ms *MemoryStore) Standing() StandingStore {
	return ms.standing
}

// Game Return the game store
func (ms *MemoryStore) Game() GameStore {
	return ms.game
}

// Matchup Return the matchup store
func (ms *MemoryStore) Matchup() MatchupStore {
	return ms.matchup
}

// Winner Return the winner store
func (ms *MemoryStore) Winner() WinnerStore {
	return ms.winner
}

// Prediction Return the prediction store
func (ms *MemoryStore) Prediction() PredictionStore {
	return ms.prediction
}
