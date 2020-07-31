package data

// Team Define the information about a team in a league
type Team struct {
	ID           string      `json:"id"`
	League       League      `json:"league"`
	Abbreviation string      `json:"abbreviation"`
	Name         string      `json:"name"`
	Fullname     string      `json:"fullname"`
	City         string      `json:"city"`
	Active       bool        `json:"active"`
	CreationYear string      `json:"creation_year"`
	Website      string      `json:"website"`
	Venue        *Venue      `json:"venue"`
	Conference   *Conference `json:"conference"`
	Division     *Division   `json:"division"`
}
