package data

// DeleteStandingRequest Is the info for a delete team request
type DeleteStandingRequest struct {
	SessionID string `json:"session_id"`
}

// DeleteStandingReply Is the reply to a delete team request
type DeleteStandingReply struct {
	Result Result `json:"result"`
}
