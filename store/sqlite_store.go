package store

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"nhlpool.com/service/go/nhlpool/config"
)

// SqliteStore Is a data store that keep it in Sqlite
type SqliteStore struct {
	database   *sql.DB
	player     *SqliteStorePlayer
	session    *SqliteStoreSession
	league     *SqliteStoreLeague
	venue      *SqliteStoreVenue
	team       *SqliteStoreTeam
	season     *SqliteStoreSeason
	standing   *SqliteStoreStanding
	game       *SqliteStoreGame
	matchup    *SqliteStoreMatchup
	winner     *SqliteStoreWinner
	prediction *SqliteStorePrediction
	conference *SqliteStoreConference
	division   *SqliteStoreDivision
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
	store.team = NewSqliteStoreTeam(store.database, store)
	store.season = NewSqliteStoreSeason(store.database, store)
	store.standing = NewSqliteStoreStanding(store.database, store)
	store.game = NewSqliteStoreGame(store.database, store)
	store.matchup = NewSqliteStoreMatchup(store.database, store)
	store.winner = NewSqliteStoreWinner(store.database, store)
	store.prediction = NewSqliteStorePrediction(store.database, store)
	store.conference = NewSqliteStoreConference(store.database, store)
	store.division = NewSqliteStoreDivision(store.database, store)
	return store
}

// Close the database
func (st *SqliteStore) Close() error {
	return st.database.Close()
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

// Team Return the team store
func (st *SqliteStore) Team() TeamStore {
	return st.team
}

// Season Return the season store
func (st *SqliteStore) Season() SeasonStore {
	return st.season
}

// Standing Return the standing store
func (st *SqliteStore) Standing() StandingStore {
	return st.standing
}

// Game Return the game store
func (st *SqliteStore) Game() GameStore {
	return st.game
}

// Matchup Return the matchup store
func (st *SqliteStore) Matchup() MatchupStore {
	return st.matchup
}

// Winner Return the winner store
func (st *SqliteStore) Winner() WinnerStore {
	return st.winner
}

// Prediction Return the prediction store
func (st *SqliteStore) Prediction() PredictionStore {
	return st.prediction
}

// Conference Return the conference store
func (st *SqliteStore) Conference() ConferenceStore {
	return st.conference
}

// Division Return the division store
func (st *SqliteStore) Division() DivisionStore {
	return st.division
}

// Clean Empty the store
func (st *SqliteStore) Clean() error {
	errPlayer := st.player.Clean()
	errSession := st.session.Clean()
	errLeague := st.league.Clean()
	errVenue := st.venue.Clean()
	errTeam := st.team.Clean()
	errSeason := st.season.Clean()
	errStanding := st.standing.Clean()
	errGame := st.game.Clean()
	errMatchup := st.matchup.Clean()
	errWinner := st.winner.Clean()
	errPrediction := st.prediction.Clean()
	errConference := st.conference.Clean()
	errDivision := st.division.Clean()
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
	if errTeam != nil {
		return errTeam
	}
	if errSeason != nil {
		return errSeason
	}
	if errStanding != nil {
		return errStanding
	}
	if errGame != nil {
		return errGame
	}
	if errMatchup != nil {
		return errMatchup
	}
	if errWinner != nil {
		return errWinner
	}
	if errPrediction != nil {
		return errPrediction
	}
	if errConference != nil {
		return errConference
	}
	if errDivision != nil {
		return errDivision
	}
	return nil
}

func (st *SqliteStore) tableExist(table string) bool {
	row := st.database.QueryRow("SELECT name FROM sqlite_master WHERE type ='table' AND name=?", table)
	var name string

	if row != nil {
		err := row.Scan(&name)
		if err != nil {
			fmt.Printf("tableExist Scan error Table:%v Err:%v\n", table, err)
			return false
		}
		return name == table
	}
	fmt.Printf("tableExist empty Table:%v\n", table)
	return false
}
