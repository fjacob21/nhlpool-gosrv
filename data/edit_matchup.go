package data

// EditMatchupRequest Is the info for an edit matchup request
type EditMatchupRequest struct {
	SessionID string `json:"session_id"`
	HomeID    string `json:"home_id"`
	AwayID    string `json:"round"`
	Round     int    `json:"pts"`
	Start     string `json:"start"`
}

// EditMatchupReply Is the reply to an edit matchup request
type EditMatchupReply struct {
	Result  Result  `json:"result"`
	Matchup Matchup `json:"matchup"`
}
