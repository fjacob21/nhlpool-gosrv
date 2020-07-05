package data

// ImportPlayerRequest Is the info for an import player request
type ImportPlayerRequest struct {
	SessionID string `json:"session_id"`
	Player    Player `json:"player"`
}

// ImportPlayerReply Is the reply to an import player request
type ImportPlayerReply struct {
	Result Result `json:"result"`
	Player Player `json:"player"`
}
