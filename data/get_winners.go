package data

// GetWinnersReply Define the information about a winner in the pool
type GetWinnersReply struct {
	Result  Result   `json:"result"`
	Winners []Winner `json:"winners"`
}
