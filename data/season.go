package data

// Season Define the information about a year for a league
type Season struct {
	Year   int     `json:"year"`
	League *League `json:"league"`
}
