package store

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"time"

	"nhlpool.com/service/go/nhlpool/data"
)

// SessionManager Store info about connection
type SessionManager struct {
	sessions map[string]data.LoginData
}

// NewSessionManager Create a new sessions manager
func NewSessionManager() *SessionManager {
	manager := &SessionManager{}
	return manager
}

func generateSessionID(loginData *data.LoginData) string {
	hash := sha256.Sum256([]byte(loginData.Player.Name + loginData.LoginTime.String()))
	return fmt.Sprintf("%x", hash)
}

// Clean Empty the sessions
func (sm *SessionManager) Clean() {
	sm.sessions = make(map[string]data.LoginData)
}

// Login to the system
func (sm *SessionManager) Login(player *data.Player, password string) string {
	sm.load()
	session := sm.find(player)
	if session != nil {
		fmt.Printf("Found session\n")
		return session.SessionID
	}
	pswOK := player.PasswordOK(password)
	if !pswOK {
		fmt.Printf("Bad psw\n")
		return ""
	}
	var data data.LoginData
	data.LoginTime = time.Now()
	data.Player = *player
	data.SessionID = generateSessionID(&data)
	sm.sessions[data.SessionID] = data
	sm.store()
	return data.SessionID
}

// Logout the session
func (sm *SessionManager) Logout(SessionID string) error {
	sm.load()
	session := sm.Get(SessionID)
	if session == nil {
		return errors.New("Invalid result")
	}
	delete(sm.sessions, SessionID)
	sm.store()
	return nil

}

// Get the sesson from the ID
func (sm *SessionManager) Get(SessionID string) *data.LoginData {
	sm.load()
	session, ok := sm.sessions[SessionID]
	if !ok {
		return nil
	}
	return &session
}

func (sm *SessionManager) load() {
	sm.sessions = store.LoadSessions()

}

func (sm *SessionManager) store() {
	store.StoreSessions(sm.sessions)
}

func (sm *SessionManager) find(player *data.Player) *data.LoginData {
	for _, session := range sm.sessions {
		if session.Player.ID == player.ID {
			return &session
		}
	}
	return nil
}
