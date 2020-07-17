package store

import "nhlpool.com/service/go/nhlpool/data"

// TeamStore interface of team storage
type TeamStore interface {
	Clean() error
	GetTeams(league *data.League) ([]data.Team, error)
	GetTeam(id string, league *data.League) (*data.Team, error)
	AddTeam(team *data.Team) error
	UpdateTeam(team *data.Team) error
	DeleteTeam(team *data.Team) error
}
