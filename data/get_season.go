package data

// GetSeasonReply Define the information about a season in the pool
type GetSeasonReply struct {
	Result Result `json:"result"`
	Season Season `json:"season"`
}
