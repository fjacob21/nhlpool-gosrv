package data

// EditWinnerRequest Is the info for an edit winner request
type EditWinnerRequest struct {
	SessionID string `json:"session_id"`
	Winner    string `json:"winner"`
}

// EditWinnerReply Is the reply to an edit winner request
type EditWinnerReply struct {
	Result Result `json:"result"`
	Winner Winner `json:"winner"`
}
