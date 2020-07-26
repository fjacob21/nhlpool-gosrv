package store

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"nhlpool.com/service/go/nhlpool/data"
)

// SqliteStoreMatchup Is a matchup data store for that keep it in Sqlite
type SqliteStoreMatchup struct {
	database *sql.DB
	store    *SqliteStore
}

// NewSqliteStoreMatchup Create a new sqlite store
func NewSqliteStoreMatchup(database *sql.DB, store *SqliteStore) *SqliteStoreMatchup {
	newStore := &SqliteStoreMatchup{database: database, store: store}
	newStore.createTables()
	return newStore
}

func (st *SqliteStoreMatchup) createTables() error {
	if !st.store.tableExist("matchup") {
		err := st.createTable()
		if err != nil {
			return err
		}
	}

	return nil
}

func (st *SqliteStoreMatchup) createTable() error {
	statement, err := st.database.Prepare(`CREATE TABLE IF NOT EXISTS matchup
	(league_id TEXT NOT NULL, season_year INTEGER NOT NULL, id TEXT NOT NULL, home TEXT, away TEXT, round INTEGER,
		start TEXT, PRIMARY KEY(league_id, season_year, id))`)
	if err != nil {
		fmt.Printf("createTable Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Printf("createTable Exec Err: %v\n", err)
		return err
	}
	fmt.Printf("createTable matchup\n")
	return nil
}

func (st *SqliteStoreMatchup) cleanTable() error {
	fmt.Printf("cleanTable matchup\n")
	statement, err := st.database.Prepare("DROP TABLE matchup")
	if err != nil {
		fmt.Printf("cleanTable Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Printf("cleanTable Exec Err: %v\n", err)
		return err
	}
	return nil
}

// Clean Empty the store
func (st *SqliteStoreMatchup) Clean() error {
	errMatchup := st.cleanTable()
	errCreate := st.createTables()
	if errMatchup != nil {
		return errMatchup
	}
	if errCreate != nil {
		return errCreate
	}
	return nil
}

// AddMatchup Add a new venue
func (st *SqliteStoreMatchup) AddMatchup(matchup *data.Matchup) error {
	statement, err := st.database.Prepare(`INSERT INTO matchup (league_id, season_year, id,
		home, away, round, start) VALUES (?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		fmt.Printf("AddMatchup Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(matchup.League.ID, matchup.Season.Year, matchup.ID, matchup.Home.ID, matchup.Away.ID, matchup.Round, matchup.Start.Format(time.RFC3339))
	if err != nil {
		fmt.Printf("AddMatchup Exec Err: %v\n", err)
		return err
	}
	return nil
}

// UpdateMatchup Update a venue info
func (st *SqliteStoreMatchup) UpdateMatchup(matchup *data.Matchup) error {
	statement, err := st.database.Prepare("UPDATE matchup SET home=?, away=?, round=?, start=? WHERE league_id=? AND season_year=? AND id=?")
	if err != nil {
		fmt.Printf("UpdateMatchup Prepare Err: %v\n", err)
		return err
	}
	res, err := statement.Exec(matchup.Home.ID, matchup.Away.ID, matchup.Round, matchup.Start.Format(time.RFC3339), matchup.League.ID, matchup.Season.Year, matchup.ID)
	if err != nil {
		fmt.Printf("UpdateMatchup Exec Err: %v\n", err)
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("UpdateMatchup RowsAffected Err: %v\n", err)
		return err
	}
	if row != 1 {
		fmt.Printf("UpdateMatchup Do not update any row\n")
		return errors.New("Invalid Matchup")
	}
	return nil
}

// DeleteMatchup Delete a venue
func (st *SqliteStoreMatchup) DeleteMatchup(matchup *data.Matchup) error {
	statement, err := st.database.Prepare("DELETE FROM matchup WHERE league_id=? AND season_year=? AND id=?")
	if err != nil {
		fmt.Printf("DeleteMatchup Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(matchup.League.ID, matchup.Season.Year, matchup.ID)
	if err != nil {
		fmt.Printf("DeleteMatchup Exec Err: %v\n", err)
		return err
	}
	return nil
}

// GetMatchup Get a venue
func (st *SqliteStoreMatchup) GetMatchup(league *data.League, season *data.Season, id string) (*data.Matchup, error) {
	row := st.database.QueryRow(`SELECT league_id, season_year, id,
	home, away, round, start
	FROM matchup WHERE league_id=? AND season_year=? AND id=?`, league.ID, season.Year, id)
	var leagueID string
	var seasonYear int
	var ID string
	var homeID string
	var awayID string
	var round int
	var start string

	if row != nil {
		matchup := &data.Matchup{}
		err := row.Scan(&leagueID, &seasonYear, &ID, &homeID, &awayID, &round, &start)
		if err != nil {
			fmt.Printf("GetMatchup Scan Err: %v\n", err)
			return nil, err
		}
		league, _ := st.store.League().GetLeague(leagueID)
		season, _ := st.store.Season().GetSeason(seasonYear, league)
		home, _ := st.store.Team().GetTeam(homeID, league)
		away, _ := st.store.Team().GetTeam(awayID, league)
		matchup.League = *league
		matchup.Season = *season
		matchup.ID = ID
		matchup.Home = *home
		matchup.Away = *away
		matchup.Round = round
		matchup.Start, err = time.Parse(time.RFC3339, start)
		if err != nil {
			fmt.Printf("GetMatchup Invalid time Err: %v\n", err)

		}
		matchup.SeasonGames, _ = st.store.Game().GetSeasonGames(league, season, home, away)
		matchup.PlayoffGames, _ = st.store.Game().GetPlayoffGames(league, season, home, away)
		matchup.CalculateResult()
		return matchup, nil
	}
	return nil, errors.New("Matchup not found")
}

// GetMatchups Return a list of all team
func (st *SqliteStoreMatchup) GetMatchups(league *data.League, season *data.Season) ([]data.Matchup, error) {
	var matchups []data.Matchup
	rows, err := st.database.Query(`SELECT  league_id, season_year, id,
	home, away, round, start FROM matchup WHERE league_id=? AND season_year=?`, league.ID, season.Year)
	if err != nil {
		fmt.Printf("GetMatchups query Err: %v\n", err)
		return []data.Matchup{}, err
	}
	var leagueID string
	var seasonYear int
	var ID string
	var homeID string
	var awayID string
	var round int
	var start string
	for rows.Next() {
		matchup := &data.Matchup{}
		err := rows.Scan(&leagueID, &seasonYear, &ID, &homeID, &awayID, &round, &start)
		if err != nil {
			fmt.Printf("GetMatchup Scan Err: %v\n", err)
			return nil, err
		}
		league, _ := st.store.League().GetLeague(leagueID)
		season, _ := st.store.Season().GetSeason(seasonYear, league)
		home, _ := st.store.Team().GetTeam(homeID, league)
		away, _ := st.store.Team().GetTeam(awayID, league)
		matchup.League = *league
		matchup.Season = *season
		matchup.ID = ID
		matchup.Home = *home
		matchup.Away = *away
		matchup.Round = round
		matchup.Start, err = time.Parse(time.RFC3339, start)
		if err != nil {
			fmt.Printf("GetMatchup Invalid time Err: %v\n", err)

		}
		matchup.SeasonGames, _ = st.store.Game().GetSeasonGames(league, season, home, away)
		matchup.PlayoffGames, _ = st.store.Game().GetPlayoffGames(league, season, home, away)
		matchup.CalculateResult()
		matchups = append(matchups, *matchup)
	}
	rows.Close()
	return matchups, nil
}
