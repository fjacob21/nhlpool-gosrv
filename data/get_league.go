package data

// GetLeagueReply Define the information about a league in the pool
type GetLeagueReply struct {
	Result Result `json:"result"`
	League League `json:"league"`
}
