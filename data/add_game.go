package data

// AddGameRequest Is the info for an add game request
type AddGameRequest struct {
	HomeID   string `json:"home_id"`
	AwayID   string `json:"away_id"`
	Date     string `json:"date"`
	Type     int    `json:"type"`
	State    int    `json:"state"`
	HomeGoal int    `json:"home_goal"`
	AwayGoal int    `json:"away_goal"`
}

// AddGameReply Is the reply to an add game request
type AddGameReply struct {
	Result Result `json:"result"`
	Game   Game   `json:"game"`
}
