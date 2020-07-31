package data

// EditStandingRequest Is the info for an edit standing request
type EditStandingRequest struct {
	SessionID    string `json:"session_id"`
	Points       int    `json:"pts"`
	Win          int    `json:"win"`
	Losses       int    `json:"losses"`
	OT           int    `json:"ot"`
	GamesPlayed  int    `json:"games_played"`
	GoalsAgainst int    `json:"goals_against"`
	GoalsScored  int    `json:"goals_scored"`
	Ranks        int    `json:"ranks"`
}

// EditStandingReply Is the reply to an edit standing request
type EditStandingReply struct {
	Result   Result   `json:"result"`
	Standing Standing `json:"standing"`
}
