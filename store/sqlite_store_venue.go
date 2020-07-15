package store

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"nhlpool.com/service/go/nhlpool/data"
)

// SqliteStoreVenue Is a venue data store for that keep it in Sqlite
type SqliteStoreVenue struct {
	database *sql.DB
}

// NewSqliteStoreVenue Create a new memory store
func NewSqliteStoreVenue(database *sql.DB) *SqliteStoreVenue {
	store := &SqliteStoreVenue{database: database}
	store.createTables()
	return store
}

func (st *SqliteStoreVenue) tableExist(table string) bool {
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

func (st *SqliteStoreVenue) createTables() error {
	err := st.createTable()
	if err != nil {
		return err
	}

	return nil
}

func (st *SqliteStoreVenue) createTable() error {
	statement, err := st.database.Prepare("CREATE TABLE IF NOT EXISTS venue (id TEXT NOT NULL, league_id TEXT NOT NULL, city TEXT, name TEXT, timezone TEXT, address TEXT, PRIMARY KEY(id,league_id))")
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

func (st *SqliteStoreVenue) cleanTable() error {
	statement, err := st.database.Prepare("DROP TABLE venue")
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
func (st *SqliteStoreVenue) Clean() error {
	errVenue := st.cleanTable()
	errCreate := st.createTables()
	if errVenue != nil {
		return errVenue
	}
	if errCreate != nil {
		return errCreate
	}
	return nil
}

// AddVenue Add a new venue
func (st *SqliteStoreVenue) AddVenue(venue *data.Venue) error {
	statement, err := st.database.Prepare("INSERT INTO venue (id, league_id, city, name, timezone, address) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Printf("AddVenue Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(venue.ID, venue.League.ID, venue.City, venue.Name, venue.Timezone, venue.Address)
	if err != nil {
		fmt.Printf("AddVenue Exec Err: %v\n", err)
		return err
	}
	return nil
}

// UpdateVenue Update a venue info
func (st *SqliteStoreVenue) UpdateVenue(venue *data.Venue) error {
	statement, err := st.database.Prepare("UPDATE venue SET city=?, name=?, timezone=?, address=? WHERE id=? AND league_id=?")
	if err != nil {
		fmt.Printf("UpdateVenue Prepare Err: %v\n", err)
		return err
	}
	res, err := statement.Exec(venue.City, venue.Name, venue.Timezone, venue.Address, venue.ID, venue.League.ID)
	if err != nil {
		fmt.Printf("UpdateVenue Exec Err: %v\n", err)
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("UpdateVenue RowsAffected Err: %v\n", err)
		return err
	}
	if row != 1 {
		fmt.Printf("UpdateVenue Do not update any row\n")
		return errors.New("Invalid Venue")
	}
	return nil
}

// DeleteVenue Delete a venue
func (st *SqliteStoreVenue) DeleteVenue(venue *data.Venue) error {
	statement, err := st.database.Prepare("DELETE FROM venue WHERE id=? AND league_id=?")
	if err != nil {
		fmt.Printf("DeleteVenue Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(venue.ID, venue.League.ID)
	if err != nil {
		fmt.Printf("DeleteVenue Exec Err: %v\n", err)
		return err
	}
	return nil
}

// GetVenue Get a venue
func (st *SqliteStoreVenue) GetVenue(ID string, league *data.League) (*data.Venue, error) {
	row := st.database.QueryRow("SELECT id, league_id, city, name, timezone, address FROM venue WHERE id=? AND league_id=?", ID, league.ID)
	var id string
	var leagueID string
	var city string
	var name string
	var timezone string
	var address string
	if row != nil {
		venue := &data.Venue{}
		err := row.Scan(&id, &leagueID, &city, &name, &timezone, &address)
		if err != nil {
			fmt.Printf("GetVenue Scan Err: %v\n", err)
			return nil, err
		}
		venue.ID = id
		venue.League = *league
		venue.City = city
		venue.Name = name
		venue.Timezone = timezone
		venue.Address = address
		return venue, nil
	}
	return nil, errors.New("Venue not found")
}
