package data

// GetTeamReply Define the information about a team in the pool
type GetTeamReply struct {
	Result Result `json:"result"`
	Team   Team   `json:"team"`
}
