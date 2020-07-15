package store

import "nhlpool.com/service/go/nhlpool/data"

// VenueStore interface of venue storage
type VenueStore interface {
	Clean() error
	GetVenue(id string, league *data.League) (*data.Venue, error)
	AddVenue(venue *data.Venue) error
	UpdateVenue(venue *data.Venue) error
	DeleteVenue(venue *data.Venue) error
}
