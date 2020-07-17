package store

import (
	"errors"

	"nhlpool.com/service/go/nhlpool/data"
)

// MemoryStoreSeason Is a Season data store that keep it only in memory
type MemoryStoreSeason struct {
	seasons map[string]*data.Season
}

// NewMemoryStoreSeason Create a new team memory store
func NewMemoryStoreSeason() *MemoryStoreSeason {
	store := &MemoryStoreSeason{}
	store.seasons = make(map[string]*data.Season)
	return store
}

// Clean Empty the store
func (ms *MemoryStoreSeason) Clean() error {
	ms.seasons = make(map[string]*data.Season)
	return nil
}

// AddSeason Add a new team
func (ms *MemoryStoreSeason) AddSeason(season *data.Season) error {
	_, ok := ms.seasons[string(season.Year)+season.League.ID]
	if ok {
		return errors.New("Season already exist")
	}
	ms.seasons[string(season.Year)+season.League.ID] = season
	return nil
}

// UpdateSeason Update a team info
func (ms *MemoryStoreSeason) UpdateSeason(season *data.Season) error {
	_, ok := ms.seasons[string(season.Year)+season.League.ID]
	if !ok {
		return errors.New("Season do not exist")
	}
	ms.seasons[string(season.Year)+season.League.ID] = season
	return nil
}

// DeleteSeason Delete a team
func (ms *MemoryStoreSeason) DeleteSeason(season *data.Season) error {
	_, ok := ms.seasons[string(season.Year)+season.League.ID]
	if !ok {
		return errors.New("Season do not exist")
	}
	delete(ms.seasons, string(season.Year)+season.League.ID)
	return nil
}

// GetSeason Get a team
func (ms *MemoryStoreSeason) GetSeason(year int, league *data.League) (*data.Season, error) {
	venue, ok := ms.seasons[string(year)+league.ID]
	if !ok {
		return nil, errors.New("Season do not exist")
	}
	return venue, nil
}

// GetSeasons Get all teams
func (ms *MemoryStoreSeason) GetSeasons(league *data.League) ([]data.Season, error) {
	var seasons []data.Season
	for _, season := range ms.seasons {
		seasons = append(seasons, *season)
	}
	return seasons, nil
}
