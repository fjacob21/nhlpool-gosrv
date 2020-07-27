package data

// AddWinnerRequest Is the info for an add winner request
type AddWinnerRequest struct {
	PlayerID string `json:"player_id"`
	Winner   string `json:"winner"`
}

// AddWinnerReply Is the reply to an add winner request
type AddWinnerReply struct {
	Result Result `json:"result"`
	Winner Winner `json:"winner"`
}
