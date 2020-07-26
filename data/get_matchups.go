package data

// GetMatchupsReply Define the information about a Matchup in the pool
type GetMatchupsReply struct {
	Result   Result    `json:"result"`
	Matchups []Matchup `json:"matchups"`
}
