package store

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"nhlpool.com/service/go/nhlpool/data"
)

// SqliteStoreSession Is a data store that keep it in Sqlite
type SqliteStoreSession struct {
	database *sql.DB
	store    *SqliteStore
}

// NewSqliteStoreSession Create a new memory store
func NewSqliteStoreSession(database *sql.DB, store *SqliteStore) *SqliteStoreSession {
	newStore := &SqliteStoreSession{database: database, store: store}
	newStore.createTables()
	return newStore
}

func (st *SqliteStoreSession) tableExist(table string) bool {
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

func (st *SqliteStoreSession) createTables() error {
	if !st.tableExist("session") {
		err := st.createTable()
		if err != nil {
			return err
		}
	}

	return nil
}

func (st *SqliteStoreSession) createTable() error {
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

func (st *SqliteStoreSession) cleanTable() error {
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

// Clean Empty the store
func (st *SqliteStoreSession) Clean() error {
	errSession := st.cleanTable()
	errCreate := st.createTables()
	if errSession != nil {
		return errSession
	}
	if errCreate != nil {
		return errCreate
	}
	return nil
}

// AddSession Add a new session
func (st *SqliteStoreSession) AddSession(session *data.LoginData) error {
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
func (st *SqliteStoreSession) DeleteSession(session *data.LoginData) error {
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
func (st *SqliteStoreSession) GetSession(sessionID string) (*data.LoginData, error) {
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
		player := st.store.player.GetPlayer(playerID)
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
func (st *SqliteStoreSession) GetSessionByPlayer(player *data.Player) (*data.LoginData, error) {
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
		player := st.store.player.GetPlayer(playerID)
		if player == nil {
			fmt.Printf("GetSessionByPlayer Invalid playe Err: %v\n", err)
			return nil, errors.New("Invalid player id")
		}
		session.Player = *player
		return session, nil
	}
	return nil, errors.New("Player not found")
}
