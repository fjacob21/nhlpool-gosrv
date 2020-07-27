package data

// GetPredictionReply Define the information about a prediction in the pool
type GetPredictionReply struct {
	Result     Result     `json:"result"`
	Prediction Prediction `json:"prediction"`
}
