package data

// GetGamesReply Define the information about games in the pool
type GetGamesReply struct {
	Result Result `json:"result"`
	Games  []Game `json:"games"`
}
