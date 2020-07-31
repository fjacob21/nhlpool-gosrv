package store

import (
	"errors"

	"nhlpool.com/service/go/nhlpool/data"
)

// MemoryStoreConference Is a Conference data store that keep it only in memory
type MemoryStoreConference struct {
	conferences map[string]*data.Conference
}

// NewMemoryStoreConference Create a new conference memory store
func NewMemoryStoreConference() *MemoryStoreConference {
	store := &MemoryStoreConference{}
	store.conferences = make(map[string]*data.Conference)
	return store
}

// Clean Empty the store
func (ms *MemoryStoreConference) Clean() error {
	ms.conferences = make(map[string]*data.Conference)
	return nil
}

// AddConference Add a new conference
func (ms *MemoryStoreConference) AddConference(conference *data.Conference) error {
	_, ok := ms.conferences[conference.League.ID+conference.ID]
	if ok {
		return errors.New("Conference already exist")
	}
	ms.conferences[conference.League.ID+conference.ID] = conference
	return nil
}

// DeleteConference Delete a conference
func (ms *MemoryStoreConference) DeleteConference(conference *data.Conference) error {
	_, ok := ms.conferences[conference.League.ID+conference.ID]
	if !ok {
		return errors.New("Conference do not exist")
	}
	delete(ms.conferences, conference.League.ID+conference.ID)
	return nil
}

// GetConference Get a conference
func (ms *MemoryStoreConference) GetConference(ID string, league *data.League) (*data.Conference, error) {
	venue, ok := ms.conferences[league.ID+ID]
	if !ok {
		return nil, errors.New("Conference do not exist")
	}
	return venue, nil
}

// GetConferences Get all conferences
func (ms *MemoryStoreConference) GetConferences(league *data.League) ([]*data.Conference, error) {
	var conferences []*data.Conference
	for _, conference := range ms.conferences {
		if conference.League.ID == league.ID {
			conferences = append(conferences, conference)
		}
	}
	return conferences, nil
}
