package store

import "nhlpool.com/service/go/nhlpool/data"

// PredictionStore interface of winner storage
type PredictionStore interface {
	Clean() error
	GetPredictions(league *data.League, season *data.Season) ([]data.Prediction, error)
	GetPrediction(player *data.Player, matchup *data.Matchup, league *data.League, season *data.Season) (*data.Prediction, error)
	AddPrediction(standing *data.Prediction) error
	UpdatePrediction(standing *data.Prediction) error
	DeletePrediction(standing *data.Prediction) error
}
