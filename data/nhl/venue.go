package nhl

import (
	"fmt"

	"nhlpool.com/service/go/nhlpool/data"
)

// Venue Info about a team venue
type Venue struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Link     string    `json:"link"`
	City     *string   `json:"city"`
	TimeZone *TimeZone `json:"timeZone"`
}

// Convert Convert to data Venue
func (t *Venue) Convert() *data.Venue {
	venue := &data.Venue{}
	venue.ID = fmt.Sprintf("%d", t.ID)
	venue.City = *t.City
	venue.Name = t.Name
	venue.Timezone = t.TimeZone.ID
	venue.Address = ""
	return venue
}
