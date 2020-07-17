package store

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"nhlpool.com/service/go/nhlpool/data"
)

// SqliteStoreSeason Is a venue data store for that keep it in Sqlite
type SqliteStoreSeason struct {
	database *sql.DB
	store    *SqliteStore
}

// NewSqliteStoreSeason Create a new memory store
func NewSqliteStoreSeason(database *sql.DB, store *SqliteStore) *SqliteStoreSeason {
	newStore := &SqliteStoreSeason{database: database, store: store}
	newStore.createTables()
	return newStore
}

func (st *SqliteStoreSeason) tableExist(table string) bool {
	row := st.database.QueryRow("SELECT name FROM sqlite_master WHERE type ='table' AND name=?", table)
	var name string

	if row != nil {
		err := row.Scan(&name)
		if err != nil {
			return false
		}
		return name == table
	}

	return false
}

func (st *SqliteStoreSeason) createTables() error {
	err := st.createTable()
	if err != nil {
		return err
	}

	return nil
}

func (st *SqliteStoreSeason) createTable() error {
	statement, err := st.database.Prepare("CREATE TABLE IF NOT EXISTS season (year INT NOT NULL, league_id TEXT NOT NULL, PRIMARY KEY(year, league_id))")
	if err != nil {
		fmt.Printf("createTable season Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Printf("createTable season Exec Err: %v\n", err)
		return err
	}
	return nil
}

func (st *SqliteStoreSeason) cleanTable() error {
	statement, err := st.database.Prepare("DROP TABLE season")
	if err != nil {
		fmt.Printf("cleanTable season Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Printf("cleanTable season Exec Err: %v\n", err)
		return err
	}
	return nil
}

// Clean Empty the store
func (st *SqliteStoreSeason) Clean() error {
	errSeason := st.cleanTable()
	errCreate := st.createTables()
	if errSeason != nil {
		return errSeason
	}
	if errCreate != nil {
		return errCreate
	}
	return nil
}

// AddSeason Add a new venue
func (st *SqliteStoreSeason) AddSeason(season *data.Season) error {
	statement, err := st.database.Prepare("INSERT INTO season (year, league_id) VALUES (?, ?)")
	if err != nil {
		fmt.Printf("AddSeason Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(season.Year, season.League.ID)
	if err != nil {
		fmt.Printf("AddSeason Exec Err: %v\n", err)
		return err
	}
	return nil
}

// DeleteSeason Delete a venue
func (st *SqliteStoreSeason) DeleteSeason(season *data.Season) error {
	statement, err := st.database.Prepare("DELETE FROM season WHERE year=? AND league_id=?")
	if err != nil {
		fmt.Printf("DeleteSeason Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(season.Year, season.League.ID)
	if err != nil {
		fmt.Printf("DeleteSeason Exec Err: %v\n", err)
		return err
	}
	return nil
}

// GetSeason Get a venue
func (st *SqliteStoreSeason) GetSeason(year int, league *data.League) (*data.Season, error) {
	row := st.database.QueryRow("SELECT year, league_id FROM season WHERE year=? AND league_id=?", year, league.ID)
	var yearValue int
	var leagueID string
	if row != nil {
		season := &data.Season{}
		err := row.Scan(&yearValue, &leagueID)
		if err != nil {
			fmt.Printf("GetSeason Scan Err: %v\n", err)
			return nil, err
		}
		season.Year = yearValue
		season.League = *league

		return season, nil
	}
	return nil, errors.New("Season not found")
}

// GetSeasons Return a list of all team
func (st *SqliteStoreSeason) GetSeasons(league *data.League) ([]data.Season, error) {
	var seasons []data.Season
	rows, err := st.database.Query("SELECT year, league_id FROM season WHERE league_id=?", league.ID)
	if err != nil {
		fmt.Printf("GetSeasons query Err: %v\n", err)
		return []data.Season{}, err
	}
	var year int
	var leagueID string
	for rows.Next() {
		season := data.Season{}
		err := rows.Scan(&year, &leagueID)
		if err != nil {
			fmt.Printf("GetSeason Scan Err: %v\n", err)
			return nil, err
		}
		season.Year = year
		season.League = *league

		seasons = append(seasons, season)
	}
	return seasons, nil
}
