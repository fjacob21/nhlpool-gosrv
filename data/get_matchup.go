package data

// GetMatchupReply Define the information about a matchup in the pool
type GetMatchupReply struct {
	Result  Result  `json:"result"`
	Matchup Matchup `json:"matchup"`
}
