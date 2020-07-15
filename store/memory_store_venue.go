package store

import (
	"errors"

	"nhlpool.com/service/go/nhlpool/data"
)

// MemoryStoreVenue Is a Venue data store that keep it only in memory
type MemoryStoreVenue struct {
	venues map[string]*data.Venue
}

// NewMemoryStoreVenue Create a new venue memory store
func NewMemoryStoreVenue() *MemoryStoreVenue {
	store := &MemoryStoreVenue{}
	store.venues = make(map[string]*data.Venue)
	return store
}

// Clean Empty the store
func (ms *MemoryStoreVenue) Clean() error {
	ms.venues = make(map[string]*data.Venue)
	return nil
}

// AddVenue Add a new venue
func (ms *MemoryStoreVenue) AddVenue(venue *data.Venue) error {
	_, ok := ms.venues[venue.ID+venue.League.ID]
	if ok {
		return errors.New("Venue already exist")
	}
	ms.venues[venue.ID+venue.League.ID] = venue
	return nil
}

// UpdateVenue Update a venue info
func (ms *MemoryStoreVenue) UpdateVenue(venue *data.Venue) error {
	_, ok := ms.venues[venue.ID+venue.League.ID]
	if !ok {
		return errors.New("Venue do not exist")
	}
	ms.venues[venue.ID+venue.League.ID] = venue
	return nil
}

// DeleteVenue Delete a venue
func (ms *MemoryStoreVenue) DeleteVenue(venue *data.Venue) error {
	_, ok := ms.venues[venue.ID+venue.League.ID]
	if !ok {
		return errors.New("Venue do not exist")
	}
	delete(ms.venues, venue.ID+venue.League.ID)
	return nil
}

// GetVenue Get a venue
func (ms *MemoryStoreVenue) GetVenue(venueID string, league *data.League) (*data.Venue, error) {
	venue, ok := ms.venues[venueID+league.ID]
	if !ok {
		return nil, errors.New("Venue do not exist")
	}
	return venue, nil
}
