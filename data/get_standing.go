package data

// GetStandingReply Define the information about a standing in the pool
type GetStandingReply struct {
	Result   Result   `json:"result"`
	Standing Standing `json:"standing"`
}
