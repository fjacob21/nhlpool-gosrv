package store

import "nhlpool.com/service/go/nhlpool/data"

// LeagueStore interface of league storage
type LeagueStore interface {
	Clean() error
	AddLeague(league *data.League) error
	UpdateLeague(league *data.League) error
	DeleteLeague(league *data.League) error
	GetLeague(leagueID string) (*data.League, error)
	GetLeagues() ([]data.League, error)
}
