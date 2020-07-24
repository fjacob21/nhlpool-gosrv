package store

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"nhlpool.com/service/go/nhlpool/data"
)

// SqliteStorePlayer Is a player data store for that keep it in Sqlite
type SqliteStorePlayer struct {
	database *sql.DB
}

// NewSqliteStorePlayer Create a new memory store
func NewSqliteStorePlayer(database *sql.DB) *SqliteStorePlayer {
	store := &SqliteStorePlayer{database: database}
	store.createTables()
	return store
}

func (st *SqliteStorePlayer) tableExist(table string) bool {
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

func (st *SqliteStorePlayer) createTables() error {
	if !st.tableExist("player") {
		err := st.createTable()
		if err != nil {
			return err
		}
	}
	return nil
}

func (st *SqliteStorePlayer) createTable() error {
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

func (st *SqliteStorePlayer) cleanTable() error {
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

// Clean Empty the store
func (st *SqliteStorePlayer) Clean() error {
	errPlayer := st.cleanTable()
	errCreate := st.createTables()
	if errPlayer != nil {
		return errPlayer
	}
	if errCreate != nil {
		return errCreate
	}
	return nil
}

// GetPlayers Return a list of all players
func (st *SqliteStorePlayer) GetPlayers() ([]data.Player, error) {
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
	rows.Close()
	return players, nil
}

func (st *SqliteStorePlayer) getPlayerByName(name string) *data.Player {
	players, _ := st.GetPlayers()
	for _, player := range players {
		if player.Name == name {
			return &player
		}
	}
	return nil
}

// GetPlayer Return the player of the specified ID
func (st *SqliteStorePlayer) GetPlayer(id string) *data.Player {
	players, _ := st.GetPlayers()
	for _, player := range players {
		if player.ID == id {
			return &player
		}
	}
	return st.getPlayerByName(id)
}

// AddPlayer Add a new player
func (st *SqliteStorePlayer) AddPlayer(player *data.Player) error {
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
func (st *SqliteStorePlayer) UpdatePlayer(player *data.Player) error {
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
func (st *SqliteStorePlayer) DeletePlayer(player *data.Player) error {
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
