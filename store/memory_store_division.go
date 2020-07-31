package store

import (
	"errors"

	"nhlpool.com/service/go/nhlpool/data"
)

// MemoryStoreDivision Is a Division data store that keep it only in memory
type MemoryStoreDivision struct {
	divisions map[string]*data.Division
}

// NewMemoryStoreDivision Create a new division memory store
func NewMemoryStoreDivision() *MemoryStoreDivision {
	store := &MemoryStoreDivision{}
	store.divisions = make(map[string]*data.Division)
	return store
}

// Clean Empty the store
func (ms *MemoryStoreDivision) Clean() error {
	ms.divisions = make(map[string]*data.Division)
	return nil
}

// AddDivision Add a new division
func (ms *MemoryStoreDivision) AddDivision(division *data.Division) error {
	_, ok := ms.divisions[division.League.ID+division.ID]
	if ok {
		return errors.New("Division already exist")
	}
	ms.divisions[division.League.ID+division.ID] = division
	return nil
}

// DeleteDivision Delete a division
func (ms *MemoryStoreDivision) DeleteDivision(division *data.Division) error {
	_, ok := ms.divisions[division.League.ID+division.ID]
	if !ok {
		return errors.New("Division do not exist")
	}
	delete(ms.divisions, division.League.ID+division.ID)
	return nil
}

// GetDivision Get a division
func (ms *MemoryStoreDivision) GetDivision(ID string, league *data.League) (*data.Division, error) {
	venue, ok := ms.divisions[league.ID+ID]
	if !ok {
		return nil, errors.New("Division do not exist")
	}
	return venue, nil
}

// GetDivisions Get all divisions
func (ms *MemoryStoreDivision) GetDivisions(league *data.League) ([]*data.Division, error) {
	var divisions []*data.Division
	for _, division := range ms.divisions {
		if division.League.ID == league.ID {
			divisions = append(divisions, division)
		}
	}
	return divisions, nil
}
