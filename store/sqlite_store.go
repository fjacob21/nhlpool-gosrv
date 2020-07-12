package store

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"nhlpool.com/service/go/nhlpool/config"
	"nhlpool.com/service/go/nhlpool/data"
)

// SqliteStore Is a data store that keep it in Sqlite
type SqliteStore struct {
	database *sql.DB
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
	store.createTables()
	return store
}

func (st *SqliteStore) tableExist(table string) bool {
	// rows, err := st.database.Query("SELECT name FROM sqlite_master WHERE type ='table' AND name='" + table + "'")
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

func (st *SqliteStore) createTables() error {
	err := st.createPlayerTable()
	if err != nil {
		return err
	}
	err = st.createSessionTable()
	if err != nil {
		return err
	}
	err = st.createLeagueTable()
	if err != nil {
		return err
	}

	return nil
}

func (st *SqliteStore) createPlayerTable() error {
	statement, err := st.database.Prepare("CREATE TABLE IF NOT EXISTS player (id TEXT PRIMARY KEY, name TEXT, email TEXT, admin INTEGER, last_login TEXT, password TEXT)")
	if err != nil {
		fmt.Printf("createPlayerTable Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Printf("createPlayerTable Exec Err: %v\n", err)
		return err
	}
	return nil
}

func (st *SqliteStore) cleanPlayerTable() error {
	statement, err := st.database.Prepare("DROP TABLE player")
	if err != nil {
		fmt.Printf("cleanPlayerTable Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Printf("cleanPlayerTable Exec Err: %v\n", err)
		return err
	}
	return nil
}

func (st *SqliteStore) createSessionTable() error {
	statement, err := st.database.Prepare("CREATE TABLE IF NOT EXISTS session (id TEXT PRIMARY KEY, login_time TEXT, player_id TEXT)")
	if err != nil {
		fmt.Printf("createSessionTable Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Printf("createSessionTable Exec Err: %v\n", err)
		return err
	}
	return nil
}

func (st *SqliteStore) cleanSessionTable() error {
	statement, err := st.database.Prepare("DROP TABLE session")
	if err != nil {
		fmt.Printf("cleanSessionTable Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Printf("cleanSessionTable Exec Err: %v\n", err)
		return err
	}
	return nil
}

func (st *SqliteStore) createLeagueTable() error {
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
	return nil
}

func (st *SqliteStore) cleanLeagueTable() error {
	statement, err := st.database.Prepare("DROP TABLE league")
	if err != nil {
		fmt.Printf("cleanLeagueTable Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Printf("cleanLeagueTable Exec Err: %v\n", err)
		return err
	}
	return nil
}

// Clean Empty the store
func (st *SqliteStore) Clean() error {
	errPlayer := st.cleanPlayerTable()
	errSession := st.cleanSessionTable()
	errLeague := st.cleanLeagueTable()
	errCreate := st.createTables()
	if errPlayer != nil {
		return errPlayer
	}
	if errSession != nil {
		return errSession
	}
	if errLeague != nil {
		return errLeague
	}
	if errCreate != nil {
		return errCreate
	}
	return nil
}

// GetPlayers Return a list of all players
func (st *SqliteStore) GetPlayers() ([]data.Player, error) {
	var players []data.Player
	rows, err := st.database.Query("SELECT id, name, email, admin, last_login, password FROM player")
	if err != nil {
		fmt.Printf("GetPlayers query Err: %v\n", err)
		return []data.Player{}, err
	}
	var id string
	var name string
	var email string
	var admin int
	var lastLogin string
	var password string
	for rows.Next() {
		err := rows.Scan(&id, &name, &email, &admin, &lastLogin, &password)
		if err != nil {
			fmt.Printf("GetPlayers row scan Err: %v\n", err)
			return []data.Player{}, err
		}
		player := data.Player{}
		player.ID = id
		player.Name = name
		player.Email = email
		player.Admin = admin == 1
		if lastLogin != "" {
			tmpTime, err := time.Parse(time.RFC3339, lastLogin)
			player.LastLogin = &tmpTime
			if err != nil {
				fmt.Printf("GetPlayers bad time Err: %v\n", err)
				player.LastLogin = nil
			}
		}

		player.Password = password

		players = append(players, player)
	}
	return players, nil
}

func (st *SqliteStore) getPlayerByName(name string) *data.Player {
	players, _ := st.GetPlayers()
	for _, player := range players {
		if player.Name == name {
			return &player
		}
	}
	return nil
}

// GetPlayer Return the player of the specified ID
func (st *SqliteStore) GetPlayer(id string) *data.Player {
	players, _ := st.GetPlayers()
	for _, player := range players {
		if player.ID == id {
			return &player
		}
	}
	return st.getPlayerByName(id)
}

// AddPlayer Add a new player
func (st *SqliteStore) AddPlayer(player *data.Player) error {
	statement, err := st.database.Prepare("INSERT INTO player (id, name, email, admin, last_login, password) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		fmt.Printf("AddPlayer Prepare Err: %v\n", err)
		return err
	}
	loginDate := ""
	if player.LastLogin != nil {
		loginDate = player.LastLogin.Format(time.RFC3339)
	}
	_, err = statement.Exec(player.ID, player.Name, player.Email, player.Admin, loginDate, player.Password)
	if err != nil {
		fmt.Printf("AddPlayer Exec Err: %v\n", err)
		return err
	}
	return nil
}

// UpdatePlayer Update a player
func (st *SqliteStore) UpdatePlayer(player *data.Player) error {
	loginDate := ""
	if player.LastLogin != nil {
		loginDate = player.LastLogin.Format(time.RFC3339)
	}
	statement, err := st.database.Prepare("UPDATE player SET name=?, email=?, admin=?, last_login=?, password=? WHERE id=?")
	if err != nil {
		fmt.Printf("UpdatePlayer Prepare Err: %v\n", err)
		return err
	}
	res, err := statement.Exec(player.Name, player.Email, player.Admin, loginDate, player.Password, player.ID)
	if err != nil {
		fmt.Printf("UpdatePlayer Exec Err: %v\n", err)
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("UpdatePlayer RowsAffected Err: %v\n", err)
		return err
	}
	if row != 1 {
		fmt.Printf("UpdatePlayer Do not update any row\n")
		return errors.New("Invalid player")
	}
	return nil
}

// DeletePlayer Delete a player
func (st *SqliteStore) DeletePlayer(player *data.Player) error {
	statement, err := st.database.Prepare("DELETE FROM player WHERE id=?")
	if err != nil {
		fmt.Printf("DeletePlayer Prepare Err: %v\n", err)
		return err
	}
	res, err := statement.Exec(player.ID)
	if err != nil {
		fmt.Printf("DeletePlayer Exec Err: %v\n", err)
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("UpdatePlayer RowsAffected Err: %v\n", err)
		return err
	}
	if row != 1 {
		fmt.Printf("UpdatePlayer Do not update any row\n")
		return errors.New("Invalid player")
	}
	return nil
}

// AddSession Add a new session
func (st *SqliteStore) AddSession(session *data.LoginData) error {
	statement, err := st.database.Prepare("INSERT INTO session (id, login_time, player_id) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Printf("AddSession Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(session.SessionID, session.LoginTime.Format(time.RFC3339), session.Player.ID)
	if err != nil {
		fmt.Printf("AddSession Exec Err: %v\n", err)
		return err
	}
	return nil
}

// DeleteSession Delete a session
func (st *SqliteStore) DeleteSession(session *data.LoginData) error {
	statement, err := st.database.Prepare("DELETE FROM session WHERE id=?")
	if err != nil {
		fmt.Printf("DeleteSession Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(session.SessionID)
	if err != nil {
		fmt.Printf("DeleteSession Exec Err: %v\n", err)
		return err
	}
	return nil
}

// GetSession Return a session using it ID
func (st *SqliteStore) GetSession(sessionID string) (*data.LoginData, error) {
	row := st.database.QueryRow("SELECT id, login_time, player_id FROM session WHERE id=?", sessionID)
	var id string
	var loginTime string
	var playerID string
	if row != nil {
		session := &data.LoginData{}
		err := row.Scan(&id, &loginTime, &playerID)
		if err != nil {
			fmt.Printf("GetSession Scan Err: %v\n", err)
			return nil, err
		}
		session.SessionID = id
		session.LoginTime, err = time.Parse(time.RFC3339, loginTime)
		if err != nil {
			fmt.Printf("GetSessionByPlayer Invalid time Err: %v\n", err)
		}
		player := st.GetPlayer(playerID)
		if player == nil {
			fmt.Printf("GetSession Invalid player Err: %v\n", err)
			return nil, errors.New("Invalid player id")
		}
		session.Player = *player
		return session, nil
	}
	return nil, errors.New("Session not found")
}

// GetSessionByPlayer Return a session using it player name
func (st *SqliteStore) GetSessionByPlayer(player *data.Player) (*data.LoginData, error) {
	row := st.database.QueryRow("SELECT id, login_time, player_id FROM session WHERE player_id=?", player.ID)
	var id string
	var loginTime string
	var playerID string
	if row != nil {
		session := &data.LoginData{}
		err := row.Scan(&id, &loginTime, &playerID)
		if err != nil {
			fmt.Printf("GetSessionByPlayer Scan Err: %v\n", err)
			return nil, err
		}
		session.SessionID = id
		session.LoginTime, err = time.Parse(time.RFC3339, loginTime)
		if err != nil {
			fmt.Printf("GetSessionByPlayer Invalid time Err: %v\n", err)
		}
		player := st.GetPlayer(playerID)
		if player == nil {
			fmt.Printf("GetSessionByPlayer Invalid playe Err: %v\n", err)
			return nil, errors.New("Invalid player id")
		}
		session.Player = *player
		return session, nil
	}
	return nil, errors.New("Player not found")
}

// AddLeague Add a new league
func (st *SqliteStore) AddLeague(league *data.League) error {
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
	return nil
}

// UpdateLeague Update a league info
func (st *SqliteStore) UpdateLeague(league *data.League) error {
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
func (st *SqliteStore) DeleteLeague(league *data.League) error {
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
func (st *SqliteStore) GetLeague(leagueID string) (*data.League, error) {
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
func (st *SqliteStore) GetLeagues() ([]data.League, error) {
	var leagues []data.League
	rows, err := st.database.Query("SELECT id, name, description, website FROM league ")
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
	return leagues, nil
}
