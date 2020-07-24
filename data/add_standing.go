package data

// AddStandingRequest Is the info for an add team request
type AddStandingRequest struct {
	TeamID       string `json:"team_id"`
	Points       int    `json:"pts"`
	Win          int    `json:"win"`
	Losses       int    `json:"losses"`
	OT           int    `json:"ot"`
	GamesPlayed  int    `json:"games_played"`
	GoalsAgainst int    `json:"goals_against"`
	GoalsScored  int    `json:"goals_scored"`
	Ranks        int    `json:"ranks"`
}

// AddStandingReply Is the reply to an add team request
type AddStandingReply struct {
	Result   Result   `json:"result"`
	Standing Standing `json:"standing"`
}
