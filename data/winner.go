package data

// Winner Define the information about a player winner prediction
type Winner struct {
	League League  `json:"league"`
	Season Season  `json:"season"`
	Player *Player `json:"player"`
	Winner Team    `json:"winner"`
}
