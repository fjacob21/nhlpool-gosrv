package nhl

// Schedule Define the schedule for a team
type Schedule struct {
	TotalItems   int            `json:"totalItems"`
	TotalEvents  int            `json:"totalEvents"`
	TotalGames   int            `json:"totalGames"`
	TotalMatches int            `json:"totalMatches"`
	Wait         int            `json:"wait"`
	Dates        []ScheduleDate `json:"dates"`
}

// ScheduleDate Represent a date in a team schedule
type ScheduleDate struct {
	Date         string `json:"date"`
	TotalItems   int    `json:"totalItems"`
	TotalEvents  int    `json:"totalEvents"`
	TotalMatches int    `json:"totalMatches"`
	Games        []Game `json:"games"`
}
