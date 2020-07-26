package store

import (
	"errors"

	"nhlpool.com/service/go/nhlpool/data"
)

// MemoryStoreMatchup Is a Matchup data store that keep it only in memory
type MemoryStoreMatchup struct {
	store    *MemoryStore
	matchups map[string]*data.Matchup
}

// NewMemoryStoreMatchup Create a new matchup memory store
func NewMemoryStoreMatchup(store *MemoryStore) *MemoryStoreMatchup {
	newStore := &MemoryStoreMatchup{}
	newStore.matchups = make(map[string]*data.Matchup)
	newStore.store = store
	return newStore
}

// Clean Empty the store
func (ms *MemoryStoreMatchup) Clean() error {
	ms.matchups = make(map[string]*data.Matchup)
	return nil
}

// AddMatchup Add a new matchup
func (ms *MemoryStoreMatchup) AddMatchup(matchup *data.Matchup) error {
	_, ok := ms.matchups[matchup.League.ID+string(matchup.Season.Year)+matchup.ID]
	if ok {
		return errors.New("Matchup already exist")
	}
	ms.matchups[matchup.League.ID+string(matchup.Season.Year)+matchup.ID] = matchup
	return nil
}

// UpdateMatchup Update a matchup info
func (ms *MemoryStoreMatchup) UpdateMatchup(matchup *data.Matchup) error {
	_, ok := ms.matchups[matchup.League.ID+string(matchup.Season.Year)+matchup.ID]
	if !ok {
		return errors.New("Matchup do not exist")
	}
	ms.matchups[matchup.League.ID+string(matchup.Season.Year)+matchup.ID] = matchup
	return nil
}

// DeleteMatchup Delete a matchup
func (ms *MemoryStoreMatchup) DeleteMatchup(matchup *data.Matchup) error {
	_, ok := ms.matchups[matchup.League.ID+string(matchup.Season.Year)+matchup.ID]
	if !ok {
		return errors.New("Matchup do not exist")
	}
	delete(ms.matchups, matchup.League.ID+string(matchup.Season.Year)+matchup.ID)
	return nil
}

// GetMatchup Get a matchup
func (ms *MemoryStoreMatchup) GetMatchup(league *data.League, season *data.Season, id string) (*data.Matchup, error) {
	matchup, ok := ms.matchups[league.ID+string(season.Year)+id]
	if !ok {
		return nil, errors.New("Matchup do not exist")
	}
	matchup.SeasonGames, _ = ms.store.Game().GetSeasonGames(league, season, &matchup.Home, &matchup.Away)
	matchup.PlayoffGames, _ = ms.store.Game().GetPlayoffGames(league, season, &matchup.Home, &matchup.Away)
	matchup.CalculateResult()
	return matchup, nil
}

// GetMatchups Get all matchups
func (ms *MemoryStoreMatchup) GetMatchups(league *data.League, season *data.Season) ([]data.Matchup, error) {
	var matchups []data.Matchup
	for _, matchup := range ms.matchups {
		if matchup.League.ID == league.ID && matchup.Season.Year == season.Year {
			matchup.SeasonGames, _ = ms.store.Game().GetSeasonGames(league, season, &matchup.Home, &matchup.Away)
			matchup.PlayoffGames, _ = ms.store.Game().GetPlayoffGames(league, season, &matchup.Home, &matchup.Away)
			matchup.CalculateResult()
			matchups = append(matchups, *matchup)
		}
	}
	return matchups, nil
}
