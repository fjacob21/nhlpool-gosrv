package data

// DeleteSeasonRequest Is the info for a delete season request
type DeleteSeasonRequest struct {
	SessionID string `json:"session_id"`
}

// DeleteSeasonReply Is the reply to a delete season request
type DeleteSeasonReply struct {
	Result Result `json:"result"`
}
