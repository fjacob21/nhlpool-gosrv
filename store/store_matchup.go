package store

import (
	"nhlpool.com/service/go/nhlpool/data"
)

// MatchupStore interface of matchup storage
type MatchupStore interface {
	Clean() error
	GetMatchups(league *data.League, season *data.Season) ([]data.Matchup, error)
	GetMatchup(league *data.League, season *data.Season, id string) (*data.Matchup, error)
	AddMatchup(matchup *data.Matchup) error
	UpdateMatchup(matchup *data.Matchup) error
	DeleteMatchup(matchup *data.Matchup) error
}
