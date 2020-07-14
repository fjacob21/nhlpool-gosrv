package nhl

// TimeZone Represent a timezone
type TimeZone struct {
	ID     string `json:"id"`
	Offset int    `json:"offset"`
	TZ     string `json:"tz"`
}
