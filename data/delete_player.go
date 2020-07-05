package data

// DeletePlayerRequest Is the info for a delete player request
type DeletePlayerRequest struct {
	SessionID string `json:"session_id"`
}

// DeletePlayerReply Is the reply to a delete player request
type DeletePlayerReply struct {
	Result Result `json:"result"`
}
