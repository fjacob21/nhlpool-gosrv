package data

// GetPlayerReply Define the information about a player in the pool
type GetPlayerReply struct {
	Result Result `json:"result"`
	Player Player `json:"player"`
}
