package data

import "time"

// LoginData Is data related to a login
type LoginData struct {
	SessionID string    `json:"session_id"`
	LoginTime time.Time `json:"login_time"`
	Player    Player    `json:"player"`
}
