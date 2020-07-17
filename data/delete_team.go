package data

// DeleteTeamRequest Is the info for a delete team request
type DeleteTeamRequest struct {
	SessionID string `json:"session_id"`
}

// DeleteTeamReply Is the reply to a delete team request
type DeleteTeamReply struct {
	Result Result `json:"result"`
}
