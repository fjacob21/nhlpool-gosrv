package data

// AddMatchupRequest Is the info for an add matchup request
type AddMatchupRequest struct {
	ID     string `json:"id"`
	HomeID string `json:"home_id"`
	AwayID string `json:"round"`
	Round  int    `json:"pts"`
	Start  string `json:"start"`
}

// AddMatchupReply Is the reply to an add matchup request
type AddMatchupReply struct {
	Result  Result  `json:"result"`
	Matchup Matchup `json:"matchup"`
}
