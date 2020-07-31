package store

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"nhlpool.com/service/go/nhlpool/data"
)

// SqliteStoreConference Is a conference data store for that keep it in Sqlite
type SqliteStoreConference struct {
	database *sql.DB
	store    *SqliteStore
}

// NewSqliteStoreConference Create a new memory store
func NewSqliteStoreConference(database *sql.DB, store *SqliteStore) *SqliteStoreConference {
	newStore := &SqliteStoreConference{database: database, store: store}
	newStore.createTables()
	return newStore
}

func (st *SqliteStoreConference) createTables() error {
	if !st.store.tableExist("conference") {
		err := st.createTable()
		if err != nil {
			return err
		}
	}

	return nil
}

func (st *SqliteStoreConference) createTable() error {
	statement, err := st.database.Prepare("CREATE TABLE IF NOT EXISTS conference (id TEXT NOT NULL, league_id TEXT NOT NULL, name TEXT, PRIMARY KEY(id,league_id))")
	if err != nil {
		fmt.Printf("createTable Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Printf("createTable Exec Err: %v\n", err)
		return err
	}
	return nil
}

func (st *SqliteStoreConference) cleanTable() error {
	statement, err := st.database.Prepare("DROP TABLE conference")
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
func (st *SqliteStoreConference) Clean() error {
	errConference := st.cleanTable()
	errCreate := st.createTables()
	if errConference != nil {
		return errConference
	}
	if errCreate != nil {
		return errCreate
	}
	return nil
}

// AddConference Add a new conference
func (st *SqliteStoreConference) AddConference(conference *data.Conference) error {
	statement, err := st.database.Prepare("INSERT INTO conference (id, league_id, name) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Printf("AddConference Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(conference.ID, conference.League.ID, conference.Name)
	if err != nil {
		fmt.Printf("AddConference Exec Err: %v\n", err)
		return err
	}
	return nil
}

// DeleteConference Delete a conference
func (st *SqliteStoreConference) DeleteConference(conference *data.Conference) error {
	statement, err := st.database.Prepare("DELETE FROM conference WHERE id=? AND league_id=?")
	if err != nil {
		fmt.Printf("DeleteConference Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(conference.ID, conference.League.ID)
	if err != nil {
		fmt.Printf("DeleteConference Exec Err: %v\n", err)
		return err
	}
	return nil
}

// GetConferences Return a list of all conference
func (st *SqliteStoreConference) GetConferences(league *data.League) ([]*data.Conference, error) {
	var conferences []*data.Conference
	rows, err := st.database.Query("SELECT id, league_id, name FROM conference WHERE league_id=?", league.ID)
	if err != nil {
		fmt.Printf("GetConferences query Err: %v\n", err)
		return []*data.Conference{}, err
	}
	var id string
	var leagueID string
	var name string
	for rows.Next() {
		conference := &data.Conference{}
		err := rows.Scan(&id, &leagueID, &name)
		if err != nil {
			fmt.Printf("GetConferences Scan Err: %v\n", err)
			return nil, err
		}
		conference.ID = id
		conference.League = *league
		conference.Name = name

		conferences = append(conferences, conference)
	}
	rows.Close()
	return conferences, nil
}

// GetConference Get a conference
func (st *SqliteStoreConference) GetConference(ID string, league *data.League) (*data.Conference, error) {
	row := st.database.QueryRow("SELECT id, league_id, name FROM conference WHERE id=? AND league_id=?", ID, league.ID)
	var id string
	var leagueID string
	var name string
	if row != nil {
		conference := &data.Conference{}
		err := row.Scan(&id, &leagueID, &name)
		if err != nil {
			fmt.Printf("GetConference Scan Err: %v\n", err)
			return nil, err
		}
		conference.ID = id
		conference.League = *league
		conference.Name = name
		return conference, nil
	}
	return nil, errors.New("Conference not found")
}
