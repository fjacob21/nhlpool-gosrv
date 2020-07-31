package store

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"nhlpool.com/service/go/nhlpool/data"
)

// SqliteStoreDivision Is a division data store for that keep it in Sqlite
type SqliteStoreDivision struct {
	database *sql.DB
	store    *SqliteStore
}

// NewSqliteStoreDivision Create a new memory store
func NewSqliteStoreDivision(database *sql.DB, store *SqliteStore) *SqliteStoreDivision {
	newStore := &SqliteStoreDivision{database: database, store: store}
	newStore.createTables()
	return newStore
}

func (st *SqliteStoreDivision) createTables() error {
	if !st.store.tableExist("division") {
		err := st.createTable()
		if err != nil {
			return err
		}
	}

	return nil
}

func (st *SqliteStoreDivision) createTable() error {
	statement, err := st.database.Prepare("CREATE TABLE IF NOT EXISTS division (id TEXT NOT NULL, league_id TEXT NOT NULL, name TEXT, PRIMARY KEY(id,league_id))")
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

func (st *SqliteStoreDivision) cleanTable() error {
	statement, err := st.database.Prepare("DROP TABLE division")
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
func (st *SqliteStoreDivision) Clean() error {
	errDivision := st.cleanTable()
	errCreate := st.createTables()
	if errDivision != nil {
		return errDivision
	}
	if errCreate != nil {
		return errCreate
	}
	return nil
}

// AddDivision Add a new division
func (st *SqliteStoreDivision) AddDivision(division *data.Division) error {
	statement, err := st.database.Prepare("INSERT INTO division (id, league_id, name) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Printf("AddDivision Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(division.ID, division.League.ID, division.Name)
	if err != nil {
		fmt.Printf("AddDivision Exec Err: %v\n", err)
		return err
	}
	return nil
}

// DeleteDivision Delete a division
func (st *SqliteStoreDivision) DeleteDivision(division *data.Division) error {
	statement, err := st.database.Prepare("DELETE FROM division WHERE id=? AND league_id=?")
	if err != nil {
		fmt.Printf("DeleteDivision Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(division.ID, division.League.ID)
	if err != nil {
		fmt.Printf("DeleteDivision Exec Err: %v\n", err)
		return err
	}
	return nil
}

// GetDivisions Return a list of all division
func (st *SqliteStoreDivision) GetDivisions(league *data.League) ([]*data.Division, error) {
	var divisions []*data.Division
	rows, err := st.database.Query("SELECT id, league_id, name FROM division WHERE league_id=?", league.ID)
	if err != nil {
		fmt.Printf("GetDivisions query Err: %v\n", err)
		return []*data.Division{}, err
	}
	var id string
	var leagueID string
	var name string
	for rows.Next() {
		division := &data.Division{}
		err := rows.Scan(&id, &leagueID, &name)
		if err != nil {
			fmt.Printf("GetDivisions Scan Err: %v\n", err)
			return nil, err
		}
		division.ID = id
		division.League = *league
		division.Name = name

		divisions = append(divisions, division)
	}
	rows.Close()
	return divisions, nil
}

// GetDivision Get a division
func (st *SqliteStoreDivision) GetDivision(ID string, league *data.League) (*data.Division, error) {
	row := st.database.QueryRow("SELECT id, league_id, name FROM division WHERE id=? AND league_id=?", ID, league.ID)
	var id string
	var leagueID string
	var name string
	if row != nil {
		division := &data.Division{}
		err := row.Scan(&id, &leagueID, &name)
		if err != nil {
			fmt.Printf("GetDivision Scan Err: %v\n", err)
			return nil, err
		}
		division.ID = id
		division.League = *league
		division.Name = name
		return division, nil
	}
	return nil, errors.New("Division not found")
}
