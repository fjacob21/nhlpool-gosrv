package data

// DeletePredictionRequest Is the info for a delete prediction request
type DeletePredictionRequest struct {
	SessionID string `json:"session_id"`
}

// DeletePredictionReply Is the reply to a delete prediction request
type DeletePredictionReply struct {
	Result Result `json:"result"`
}
