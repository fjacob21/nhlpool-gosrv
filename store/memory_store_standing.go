package store

import (
	"errors"

	"nhlpool.com/service/go/nhlpool/data"
)

// MemoryStoreStanding Is a Standing data store that keep it only in memory
type MemoryStoreStanding struct {
	standings map[string]*data.Standing
}

// NewMemoryStoreStanding Create a new standing memory store
func NewMemoryStoreStanding() *MemoryStoreStanding {
	store := &MemoryStoreStanding{}
	store.standings = make(map[string]*data.Standing)
	return store
}

// Clean Empty the store
func (ms *MemoryStoreStanding) Clean() error {
	ms.standings = make(map[string]*data.Standing)
	return nil
}

// AddStanding Add a new standing
func (ms *MemoryStoreStanding) AddStanding(standing *data.Standing) error {
	_, ok := ms.standings[standing.League.ID+string(standing.Season.Year)+standing.Team.ID]
	if ok {
		return errors.New("Standing already exist")
	}
	ms.standings[standing.League.ID+string(standing.Season.Year)+standing.Team.ID] = standing
	return nil
}

// UpdateStanding Update a standing info
func (ms *MemoryStoreStanding) UpdateStanding(standing *data.Standing) error {
	_, ok := ms.standings[standing.League.ID+string(standing.Season.Year)+standing.Team.ID]
	if !ok {
		return errors.New("Standing do not exist")
	}
	ms.standings[standing.League.ID+string(standing.Season.Year)+standing.Team.ID] = standing
	return nil
}

// DeleteStanding Delete a standing
func (ms *MemoryStoreStanding) DeleteStanding(standing *data.Standing) error {
	_, ok := ms.standings[standing.League.ID+string(standing.Season.Year)+standing.Team.ID]
	if !ok {
		return errors.New("Standing do not exist")
	}
	delete(ms.standings, standing.League.ID+string(standing.Season.Year)+standing.Team.ID)
	return nil
}

// GetStanding Get a standing
func (ms *MemoryStoreStanding) GetStanding(team *data.Team, league *data.League, season *data.Season) (*data.Standing, error) {
	venue, ok := ms.standings[league.ID+string(season.Year)+team.ID]
	if !ok {
		return nil, errors.New("Standing do not exist")
	}
	return venue, nil
}

// GetStandings Get all standings
func (ms *MemoryStoreStanding) GetStandings(league *data.League, season *data.Season) ([]data.Standing, error) {
	var standings []data.Standing
	for _, standing := range ms.standings {
		if standing.League.ID == league.ID && standing.Season.Year == season.Year {
			standings = append(standings, *standing)
		}
	}
	return standings, nil
}
