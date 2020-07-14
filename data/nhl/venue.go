package nhl

// Venue Info about a team venue
type Venue struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Link     string    `json:"link"`
	City     *string   `json:"city"`
	TimeZone *TimeZone `json:"timeZone"`
}
