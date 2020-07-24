package store

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"nhlpool.com/service/go/nhlpool/data"
)

// SqliteStoreLeague Is a league data store for that keep it in Sqlite
type SqliteStoreLeague struct {
	database *sql.DB
}

// NewSqliteStoreLeague Create a new memory store
func NewSqliteStoreLeague(database *sql.DB) *SqliteStoreLeague {
	store := &SqliteStoreLeague{database: database}
	store.createTables()
	return store
}

func (st *SqliteStoreLeague) tableExist(table string) bool {
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

func (st *SqliteStoreLeague) createTables() error {
	if !st.tableExist("league") {
		err := st.createTable()
		if err != nil {
			return err
		}
	}

	return nil
}

func (st *SqliteStoreLeague) createTable() error {
	statement, err := st.database.Prepare("CREATE TABLE IF NOT EXISTS league (id TEXT PRIMARY KEY, name TEXT, description TEXT, website TEXT)")
	if err != nil {
		fmt.Printf("createLeagueTable Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Printf("createLeagueTable Exec Err: %v\n", err)
		return err
	}
	fmt.Printf("createTable league\n")
	return nil
}

func (st *SqliteStoreLeague) cleanTable() error {
	statement, err := st.database.Prepare("DROP TABLE league")
	if err != nil {
		fmt.Printf("cleanTable Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Printf("cleanTable Exec Err: %v\n", err)
		return err
	}
	fmt.Printf("cleanTable league\n")
	return nil
}

// Clean Empty the store
func (st *SqliteStoreLeague) Clean() error {
	errLeague := st.cleanTable()
	errCreate := st.createTables()
	if errLeague != nil {
		return errLeague
	}
	if errCreate != nil {
		return errCreate
	}
	return nil
}

// AddLeague Add a new league
func (st *SqliteStoreLeague) AddLeague(league *data.League) error {
	statement, err := st.database.Prepare("INSERT INTO league (id, name, description, website) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Printf("AddLeague Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(league.ID, league.Name, league.Description, league.Website)
	if err != nil {
		fmt.Printf("AddLeague Exec Err: %v\n", err)
		return err
	}
	fmt.Printf("AddLeague %v\n", league)
	return nil
}

// UpdateLeague Update a league info
func (st *SqliteStoreLeague) UpdateLeague(league *data.League) error {
	statement, err := st.database.Prepare("UPDATE league SET name=?, description=?, website=? WHERE id=?")
	if err != nil {
		fmt.Printf("UpdateLeague Prepare Err: %v\n", err)
		return err
	}
	res, err := statement.Exec(league.Name, league.Description, league.Website, league.ID)
	if err != nil {
		fmt.Printf("UpdateLeague Exec Err: %v\n", err)
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("UpdateLeague RowsAffected Err: %v\n", err)
		return err
	}
	if row != 1 {
		fmt.Printf("UpdateLeague Do not update any row\n")
		return errors.New("Invalid player")
	}
	return nil
}

// DeleteLeague Delete a league
func (st *SqliteStoreLeague) DeleteLeague(league *data.League) error {
	statement, err := st.database.Prepare("DELETE FROM league WHERE id=?")
	if err != nil {
		fmt.Printf("DeleteLeague Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(league.ID)
	if err != nil {
		fmt.Printf("DeleteLeague Exec Err: %v\n", err)
		return err
	}
	return nil
}

// GetLeague Get a league
func (st *SqliteStoreLeague) GetLeague(leagueID string) (*data.League, error) {
	row := st.database.QueryRow("SELECT id, name, description, website FROM league WHERE id=?", leagueID)
	var id string
	var name string
	var description string
	var website string
	if row != nil {
		league := &data.League{}
		err := row.Scan(&id, &name, &description, &website)
		if err != nil {
			fmt.Printf("GetLeague Scan Err: %v\n", err)
			return nil, err
		}
		league.ID = id
		league.Name = name
		league.Description = description
		league.Website = website
		return league, nil
	}
	return nil, errors.New("League not found")
}

// GetLeagues Get all leagues
func (st *SqliteStoreLeague) GetLeagues() ([]data.League, error) {
	var leagues []data.League
	rows, err := st.database.Query("SELECT id, name, description, website FROM league")
	if err != nil {
		fmt.Printf("GetLeagues query Err: %v\n", err)
		return []data.League{}, err
	}
	var id string
	var name string
	var description string
	var website string
	for rows.Next() {
		league := data.League{}
		err := rows.Scan(&id, &name, &description, &website)
		if err != nil {
			fmt.Printf("GetLeagues Scan Err: %v\n", err)
			return nil, err
		}
		league.ID = id
		league.Name = name
		league.Description = description
		league.Website = website

		leagues = append(leagues, league)
	}
	rows.Close()
	return leagues, nil
}
