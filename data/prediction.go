package data

// Prediction Define the information about a player prediction
type Prediction struct {
	League  League   `json:"league"`
	Season  Season   `json:"season"`
	Matchup *Matchup `json:"matchup"`
	Player  *Player  `json:"player"`
	Winner  Team     `json:"winner"`
	Games   int      `json:"games"`
}
