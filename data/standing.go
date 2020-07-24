package data

// Standing Define the information about a team standing in a year
type Standing struct {
	League       League `json:"league"`
	Season       Season `json:"season"`
	Team         Team   `json:"team"`
	Points       int    `json:"pts"`
	Win          int    `json:"win"`
	Losses       int    `json:"losses"`
	OT           int    `json:"ot"`
	GamesPlayed  int    `json:"games_played"`
	GoalsAgainst int    `json:"goals_against"`
	GoalsScored  int    `json:"goals_scored"`
	Ranks        int    `json:"ranks"`
}
