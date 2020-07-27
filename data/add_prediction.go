package data

// AddPredictionRequest Is the info for an add prediction request
type AddPredictionRequest struct {
	PlayerID  string `json:"player_id"`
	MatchupID string `json:"matchup_id"`
	Winner    string `json:"winner"`
	Games     int    `json:"games"`
}

// AddPredictionReply Is the reply to an add prediction request
type AddPredictionReply struct {
	Result     Result     `json:"result"`
	Prediction Prediction `json:"prediction"`
}
