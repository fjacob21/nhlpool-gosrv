package data

// DeleteLeagueRequest Is the info for a delete league request
type DeleteLeagueRequest struct {
	SessionID string `json:"session_id"`
}

// DeleteLeagueReply Is the reply to a delete league request
type DeleteLeagueReply struct {
	Result Result `json:"result"`
}
