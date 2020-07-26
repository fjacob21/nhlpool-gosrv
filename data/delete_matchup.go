package data

// DeleteMatchupRequest Is the info for a delete matchup request
type DeleteMatchupRequest struct {
	SessionID string `json:"session_id"`
}

// DeleteMatchupReply Is the reply to a delete matchup request
type DeleteMatchupReply struct {
	Result Result `json:"result"`
}
