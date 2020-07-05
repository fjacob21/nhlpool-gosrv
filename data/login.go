package data

// LoginRequest Is the info for a login request
type LoginRequest struct {
	Password string `json:"password"`
}

// LoginReply Is the reply to a login request
type LoginReply struct {
	Result    Result `json:"result"`
	SessionID string `json:"session_id"`
}
