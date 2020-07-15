package store

import (
	"errors"

	"nhlpool.com/service/go/nhlpool/data"
)

// MemoryStoreSession Is a session data store that keep it only in memory
type MemoryStoreSession struct {
	sessions map[string]data.LoginData
}

// NewMemoryStoreSession Create a new session memory store
func NewMemoryStoreSession() *MemoryStoreSession {
	store := &MemoryStoreSession{}
	store.sessions = make(map[string]data.LoginData)
	return store
}

// Clean Empty the store
func (ms *MemoryStoreSession) Clean() error {
	ms.sessions = make(map[string]data.LoginData)
	return nil
}

// AddSession Add a new session
func (ms *MemoryStoreSession) AddSession(session *data.LoginData) error {
	ms.sessions[session.SessionID] = *session
	return nil
}

// DeleteSession Delete a session
func (ms *MemoryStoreSession) DeleteSession(session *data.LoginData) error {
	delete(ms.sessions, session.SessionID)
	return nil
}

// GetSession Return a session using it ID
func (ms *MemoryStoreSession) GetSession(sessionID string) (*data.LoginData, error) {
	session, ok := ms.sessions[sessionID]
	if !ok {
		return nil, errors.New("Do not exist")
	}
	return &session, nil
}

// GetSessionByPlayer Return a session using it player name
func (ms *MemoryStoreSession) GetSessionByPlayer(player *data.Player) (*data.LoginData, error) {
	for _, session := range ms.sessions {
		if session.Player.ID == player.ID {
			return &session, nil
		}
	}
	return nil, errors.New("Do not exist")
}
