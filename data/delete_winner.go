package data

// DeleteWinnerRequest Is the info for a delete winner request
type DeleteWinnerRequest struct {
	SessionID string `json:"session_id"`
}

// DeleteWinnerReply Is the reply to a delete winner request
type DeleteWinnerReply struct {
	Result Result `json:"result"`
}
