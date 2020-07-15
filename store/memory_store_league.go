package store

import (
	"errors"

	"nhlpool.com/service/go/nhlpool/data"
)

// MemoryStoreLeague Is a League data store that keep it only in memory
type MemoryStoreLeague struct {
	leagues map[string]*data.League
}

// NewMemoryStoreLeague Create a new league memory store
func NewMemoryStoreLeague() *MemoryStoreLeague {
	store := &MemoryStoreLeague{}
	store.leagues = make(map[string]*data.League)
	return store
}

// Clean Empty the store
func (ms *MemoryStoreLeague) Clean() error {
	ms.leagues = make(map[string]*data.League)
	return nil
}

// AddLeague Add a new league
func (ms *MemoryStoreLeague) AddLeague(league *data.League) error {
	_, ok := ms.leagues[league.ID]
	if ok {
		return errors.New("League already exist")
	}
	ms.leagues[league.ID] = league
	return nil
}

// UpdateLeague Update a league info
func (ms *MemoryStoreLeague) UpdateLeague(league *data.League) error {
	_, ok := ms.leagues[league.ID]
	if !ok {
		return errors.New("League do not exist")
	}
	ms.leagues[league.ID] = league
	return nil
}

// DeleteLeague Delete a league
func (ms *MemoryStoreLeague) DeleteLeague(league *data.League) error {
	_, ok := ms.leagues[league.ID]
	if !ok {
		return errors.New("League do not exist")
	}
	delete(ms.leagues, league.ID)
	return nil
}

// GetLeague Get a league
func (ms *MemoryStoreLeague) GetLeague(leagueID string) (*data.League, error) {
	league, ok := ms.leagues[leagueID]
	if !ok {
		return nil, errors.New("League do not exist")
	}
	return league, nil
}

// GetLeagues Get all leagues
func (ms *MemoryStoreLeague) GetLeagues() ([]data.League, error) {
	var leagues []data.League
	for _, league := range ms.leagues {
		leagues = append(leagues, *league)
	}
	return leagues, nil
}
