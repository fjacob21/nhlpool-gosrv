package data

// GetPredictionsReply Define the information about a prediction in the pool
type GetPredictionsReply struct {
	Result      Result       `json:"result"`
	Predictions []Prediction `json:"prediction"`
}
