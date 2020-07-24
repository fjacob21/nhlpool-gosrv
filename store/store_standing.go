package store

import "nhlpool.com/service/go/nhlpool/data"

// StandingStore interface of standing storage
type StandingStore interface {
	Clean() error
	GetStandings(league *data.League, season *data.Season) ([]data.Standing, error)
	GetStanding(team *data.Team, league *data.League, season *data.Season) (*data.Standing, error)
	AddStanding(standing *data.Standing) error
	UpdateStanding(standing *data.Standing) error
	DeleteStanding(standing *data.Standing) error
}
