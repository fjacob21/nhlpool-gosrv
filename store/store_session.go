package store

import "nhlpool.com/service/go/nhlpool/data"

// SessionStore interface of session storage
type SessionStore interface {
	Clean() error
	AddSession(session *data.LoginData) error
	DeleteSession(session *data.LoginData) error
	GetSession(sessionID string) (*data.LoginData, error)
	GetSessionByPlayer(player *data.Player) (*data.LoginData, error)
}
