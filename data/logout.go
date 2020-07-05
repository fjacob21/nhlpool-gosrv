package data

// LogoutRequest Is the info for a logout request
type LogoutRequest struct {
	SessionID string `json:"session_id"`
}

// LogoutReply Is the reply to a logout request
type LogoutReply struct {
	Result Result `json:"result"`
}
