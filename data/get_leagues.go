package data

// GetLeaguesReply Define the information about a league in the pool
type GetLeaguesReply struct {
	Result  Result   `json:"result"`
	Leagues []League `json:"leagues"`
}
