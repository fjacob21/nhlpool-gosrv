package data

// GetTeamsReply Define the information about a team in the pool
type GetTeamsReply struct {
	Result Result `json:"result"`
	Teams  []Team `json:"teams"`
}
