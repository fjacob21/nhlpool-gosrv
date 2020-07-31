package data

// EditGameRequest Is the info for an edit game request
type EditGameRequest struct {
	SessionID string `json:"session_id"`
	HomeID    string `json:"home_id"`
	AwayID    string `json:"away_id"`
	Date      string `json:"date"`
	Type      int    `json:"type"`
	State     int    `json:"state"`
	HomeGoal  int    `json:"home_goal"`
	AwayGoal  int    `json:"away_goal"`
}

// EditGameReply Is the reply to an edit game request
type EditGameReply struct {
	Result Result `json:"result"`
	Game   Game   `json:"game"`
}
