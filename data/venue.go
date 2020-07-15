package data

// Venue Define the information about a team venue
type Venue struct {
	ID       string `json:"id"`
	League   League `json:"league"`
	City     string `json:"city"`
	Name     string `json:"name"`
	Timezone string `json:"timezone"`
	Address  string `json:"address"`
}
