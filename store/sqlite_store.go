package store

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"nhlpool.com/service/go/nhlpool/config"
)

// SqliteStore Is a data store that keep it in Sqlite
type SqliteStore struct {
	database *sql.DB
	player   *SqliteStorePlayer
	session  *SqliteStoreSession
	league   *SqliteStoreLeague
	venue    *SqliteStoreVenue
}

// NewSqliteStore Create a new memory store
func NewSqliteStore() Store {
	configs := config.LoadConfigs()

	store := &SqliteStore{}
	var err error
	store.database, err = sql.Open("sqlite3", configs.DB)
	if err != nil {
		fmt.Printf("Cannot open db DB: %v Err:%v\n", configs.DB, err)
		return nil
	}
	store.player = NewSqliteStorePlayer(store.database)
	store.session = NewSqliteStoreSession(store.database, store)
	store.league = NewSqliteStoreLeague(store.database)
	store.venue = NewSqliteStoreVenue(store.database)
	return store
}

// Player Return the player store
func (st *SqliteStore) Player() Player {
	return st.player
}

// Session Return the session store
func (st *SqliteStore) Session() SessionStore {
	return st.session
}

// League Return the league store
func (st *SqliteStore) League() LeagueStore {
	return st.league
}

// Venue Return the venue store
func (st *SqliteStore) Venue() VenueStore {
	return st.venue
}

// Clean Empty the store
func (st *SqliteStore) Clean() error {
	errPlayer := st.player.Clean()
	errSession := st.session.Clean()
	errLeague := st.league.Clean()
	errVenue := st.venue.Clean()
	if errPlayer != nil {
		return errPlayer
	}
	if errSession != nil {
		return errSession
	}
	if errLeague != nil {
		return errLeague
	}
	if errVenue != nil {
		return errVenue
	}
	return nil
}
