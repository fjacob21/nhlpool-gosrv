package data

// GetStandingsReply Define the information about a standings in the pool
type GetStandingsReply struct {
	Result    Result     `json:"result"`
	Standings []Standing `json:"standings"`
}
