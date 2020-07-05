package data

// GetPlayersReply Define the information about a player in the pool
type GetPlayersReply struct {
	Result  Result   `json:"result"`
	Players []Player `json:"players"`
}
