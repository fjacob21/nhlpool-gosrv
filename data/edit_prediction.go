package data

// EditPredictionRequest Is the info for an edit prediction request
type EditPredictionRequest struct {
	SessionID string `json:"session_id"`
	Winner    string `json:"winner"`
	Games     int    `json:"games"`
}

// EditPredictionReply Is the reply to an edit prediction request
type EditPredictionReply struct {
	Result     Result     `json:"result"`
	Prediction Prediction `json:"prediction"`
}
