package store

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"nhlpool.com/service/go/nhlpool/data"
)

// SqliteStoreTeam Is a venue data store for that keep it in Sqlite
type SqliteStoreTeam struct {
	database *sql.DB
	store    *SqliteStore
}

// NewSqliteStoreTeam Create a new memory store
func NewSqliteStoreTeam(database *sql.DB, store *SqliteStore) *SqliteStoreTeam {
	newStore := &SqliteStoreTeam{database: database, store: store}
	newStore.createTables()
	return newStore
}

func (st *SqliteStoreTeam) tableExist(table string) bool {
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

func (st *SqliteStoreTeam) createTables() error {
	err := st.createTable()
	if err != nil {
		return err
	}

	return nil
}

func (st *SqliteStoreTeam) createTable() error {
	statement, err := st.database.Prepare("CREATE TABLE IF NOT EXISTS team (id TEXT NOT NULL, league_id TEXT NOT NULL, abbreviation TEXT, name TEXT, fullname TEXT, city TEXT, active INTEGER, creation_year TEXT, website TEXT, venue_id TEXT, PRIMARY KEY(id,league_id))")
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

func (st *SqliteStoreTeam) cleanTable() error {
	statement, err := st.database.Prepare("DROP TABLE team")
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
func (st *SqliteStoreTeam) Clean() error {
	errTeam := st.cleanTable()
	errCreate := st.createTables()
	if errTeam != nil {
		return errTeam
	}
	if errCreate != nil {
		return errCreate
	}
	return nil
}

// AddTeam Add a new venue
func (st *SqliteStoreTeam) AddTeam(team *data.Team) error {
	statement, err := st.database.Prepare("INSERT INTO team (id, league_id, abbreviation, name, fullname, city, active, creation_year, website, venue_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Printf("AddTeam Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(team.ID, team.League.ID, team.Abbreviation, team.Name, team.Fullname, team.City, team.Active, team.CreationYear, team.Website, team.Venue.ID)
	if err != nil {
		fmt.Printf("AddTeam Exec Err: %v\n", err)
		return err
	}
	return nil
}

// UpdateTeam Update a venue info
func (st *SqliteStoreTeam) UpdateTeam(team *data.Team) error {
	statement, err := st.database.Prepare("UPDATE team SET abbreviation=?, name=?, fullname=?, city=?, active=?, creation_year=?, website=? WHERE id=? AND league_id=?")
	if err != nil {
		fmt.Printf("UpdateTeam Prepare Err: %v\n", err)
		return err
	}
	res, err := statement.Exec(team.Abbreviation, team.Name, team.Fullname, team.City, team.Active, team.CreationYear, team.Website, team.ID, team.League.ID)
	if err != nil {
		fmt.Printf("UpdateTeam Exec Err: %v\n", err)
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("UpdateTeam RowsAffected Err: %v\n", err)
		return err
	}
	if row != 1 {
		fmt.Printf("UpdateTeam Do not update any row\n")
		return errors.New("Invalid Team")
	}
	return nil
}

// DeleteTeam Delete a venue
func (st *SqliteStoreTeam) DeleteTeam(team *data.Team) error {
	statement, err := st.database.Prepare("DELETE FROM team WHERE id=? AND league_id=?")
	if err != nil {
		fmt.Printf("DeleteTeam Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(team.ID, team.League.ID)
	if err != nil {
		fmt.Printf("DeleteTeam Exec Err: %v\n", err)
		return err
	}
	return nil
}

// GetTeam Get a venue
func (st *SqliteStoreTeam) GetTeam(ID string, league *data.League) (*data.Team, error) {
	row := st.database.QueryRow("SELECT id, league_id, abbreviation, name, fullname, city, active, creation_year, website, venue_id FROM team WHERE id=? AND league_id=?", ID, league.ID)
	var id string
	var leagueID string
	var abbreviation string
	var name string
	var fullname string
	var city string
	var active int
	var creationYear string
	var website string
	var venueID string
	if row != nil {
		team := &data.Team{}
		err := row.Scan(&id, &leagueID, &abbreviation, &name, &fullname, &city, &active, &creationYear, &website, &venueID)
		if err != nil {
			fmt.Printf("GetTeam Scan Err: %v\n", err)
			return nil, err
		}
		team.ID = id
		team.League = *league
		team.Abbreviation = abbreviation
		team.Name = name
		team.Fullname = fullname
		team.City = city
		team.Active = active == 1
		team.CreationYear = creationYear
		team.Website = website
		team.Venue, _ = st.store.Venue().GetVenue(venueID, league)

		return team, nil
	}
	return nil, errors.New("Team not found")
}

// GetTeams Return a list of all team
func (st *SqliteStoreTeam) GetTeams(league *data.League) ([]data.Team, error) {
	var teams []data.Team
	rows, err := st.database.Query("SELECT id, league_id, abbreviation, name, fullname, city, active, creation_year, website, venue_id FROM team WHERE league_id=?", league.ID)
	if err != nil {
		fmt.Printf("GetTeams query Err: %v\n", err)
		return []data.Team{}, err
	}
	var id string
	var leagueID string
	var abbreviation string
	var name string
	var fullname string
	var city string
	var active int
	var creationYear string
	var website string
	var venueID string
	for rows.Next() {
		team := data.Team{}
		err := rows.Scan(&id, &leagueID, &abbreviation, &name, &fullname, &city, &active, &creationYear, &website, &venueID)
		if err != nil {
			fmt.Printf("GetTeam Scan Err: %v\n", err)
			return nil, err
		}
		team.ID = id
		team.League = *league
		team.Abbreviation = abbreviation
		team.Name = name
		team.Fullname = fullname
		team.City = city
		team.Active = active == 1
		team.CreationYear = creationYear
		team.Website = website
		team.Venue, _ = st.store.Venue().GetVenue(venueID, league)

		teams = append(teams, team)
	}
	return teams, nil
}
