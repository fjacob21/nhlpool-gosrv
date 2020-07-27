package data

// GetWinnerReply Define the information about a winner in the pool
type GetWinnerReply struct {
	Result Result `json:"result"`
	Winner Winner `json:"winner"`
}
