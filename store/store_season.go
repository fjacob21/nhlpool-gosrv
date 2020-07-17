package store

import "nhlpool.com/service/go/nhlpool/data"

// SeasonStore interface of season storage
type SeasonStore interface {
	Clean() error
	GetSeasons(league *data.League) ([]data.Season, error)
	GetSeason(id int, league *data.League) (*data.Season, error)
	AddSeason(season *data.Season) error
	DeleteSeason(season *data.Season) error
}
