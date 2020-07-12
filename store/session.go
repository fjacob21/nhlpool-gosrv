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

// Login to the system
func (sm *SessionManager) Login(player *data.Player, password string) string {
	pswOK := player.PasswordOK(password)
	if !pswOK {
		fmt.Printf("Bad psw\n")
		return ""
	}
	session, _ := GetStore().GetSessionByPlayer(player)
	if session != nil {
		return session.SessionID
	}
	data := &data.LoginData{}
	data.LoginTime = time.Now()
	data.Player = *player
	data.SessionID = generateSessionID(data)
	GetStore().AddSession(data)
	return data.SessionID
}

// Logout the session
func (sm *SessionManager) Logout(sessionID string) error {
	session := sm.Get(sessionID)
	if session == nil {
		return errors.New("Invalid result")
	}
	GetStore().DeleteSession(session)
	return nil

}

// Get the sesson from the ID
func (sm *SessionManager) Get(sessionID string) *data.LoginData {
	session, _ := GetStore().GetSession(sessionID)
	return session
}
