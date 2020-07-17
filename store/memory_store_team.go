package store

import (
	"errors"

	"nhlpool.com/service/go/nhlpool/data"
)

// MemoryStoreTeam Is a Team data store that keep it only in memory
type MemoryStoreTeam struct {
	teams map[string]*data.Team
}

// NewMemoryStoreTeam Create a new team memory store
func NewMemoryStoreTeam() *MemoryStoreTeam {
	store := &MemoryStoreTeam{}
	store.teams = make(map[string]*data.Team)
	return store
}

// Clean Empty the store
func (ms *MemoryStoreTeam) Clean() error {
	ms.teams = make(map[string]*data.Team)
	return nil
}

// AddTeam Add a new team
func (ms *MemoryStoreTeam) AddTeam(team *data.Team) error {
	_, ok := ms.teams[team.ID+team.League.ID]
	if ok {
		return errors.New("Team already exist")
	}
	ms.teams[team.ID+team.League.ID] = team
	return nil
}

// UpdateTeam Update a team info
func (ms *MemoryStoreTeam) UpdateTeam(team *data.Team) error {
	_, ok := ms.teams[team.ID+team.League.ID]
	if !ok {
		return errors.New("Team do not exist")
	}
	ms.teams[team.ID+team.League.ID] = team
	return nil
}

// DeleteTeam Delete a team
func (ms *MemoryStoreTeam) DeleteTeam(team *data.Team) error {
	_, ok := ms.teams[team.ID+team.League.ID]
	if !ok {
		return errors.New("Team do not exist")
	}
	delete(ms.teams, team.ID+team.League.ID)
	return nil
}

// GetTeam Get a team
func (ms *MemoryStoreTeam) GetTeam(ID string, league *data.League) (*data.Team, error) {
	venue, ok := ms.teams[ID+league.ID]
	if !ok {
		return nil, errors.New("Team do not exist")
	}
	return venue, nil
}

// GetTeams Get all teams
func (ms *MemoryStoreTeam) GetTeams(league *data.League) ([]data.Team, error) {
	var teams []data.Team
	for _, team := range ms.teams {
		teams = append(teams, *team)
	}
	return teams, nil
}
