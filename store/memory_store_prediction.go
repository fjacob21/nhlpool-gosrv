package store

import (
	"errors"

	"nhlpool.com/service/go/nhlpool/data"
)

// MemoryStorePrediction Is a Prediction data store that keep it only in memory
type MemoryStorePrediction struct {
	predictions map[string]*data.Prediction
}

// NewMemoryStorePrediction Create a new prediction memory store
func NewMemoryStorePrediction() *MemoryStorePrediction {
	store := &MemoryStorePrediction{}
	store.predictions = make(map[string]*data.Prediction)
	return store
}

// Clean Empty the store
func (ms *MemoryStorePrediction) Clean() error {
	ms.predictions = make(map[string]*data.Prediction)
	return nil
}

// AddPrediction Add a new prediction
func (ms *MemoryStorePrediction) AddPrediction(prediction *data.Prediction) error {
	_, ok := ms.predictions[prediction.League.ID+string(prediction.Season.Year)+prediction.Player.ID+prediction.Matchup.ID]
	if ok {
		return errors.New("Prediction already exist")
	}
	ms.predictions[prediction.League.ID+string(prediction.Season.Year)+prediction.Player.ID+prediction.Matchup.ID] = prediction
	return nil
}

// UpdatePrediction Update a prediction info
func (ms *MemoryStorePrediction) UpdatePrediction(prediction *data.Prediction) error {
	_, ok := ms.predictions[prediction.League.ID+string(prediction.Season.Year)+prediction.Player.ID+prediction.Matchup.ID]
	if !ok {
		return errors.New("Prediction do not exist")
	}
	ms.predictions[prediction.League.ID+string(prediction.Season.Year)+prediction.Player.ID+prediction.Matchup.ID] = prediction
	return nil
}

// DeletePrediction Delete a prediction
func (ms *MemoryStorePrediction) DeletePrediction(prediction *data.Prediction) error {
	_, ok := ms.predictions[prediction.League.ID+string(prediction.Season.Year)+prediction.Player.ID+prediction.Matchup.ID]
	if !ok {
		return errors.New("Prediction do not exist")
	}
	delete(ms.predictions, prediction.League.ID+string(prediction.Season.Year)+prediction.Player.ID+prediction.Matchup.ID)
	return nil
}

// GetPrediction Get a prediction
func (ms *MemoryStorePrediction) GetPrediction(player *data.Player, matchup *data.Matchup, league *data.League, season *data.Season) (*data.Prediction, error) {
	prediction, ok := ms.predictions[league.ID+string(season.Year)+player.ID+matchup.ID]
	if !ok {
		return nil, errors.New("Prediction do not exist")
	}
	return prediction, nil
}

// GetPredictions Get all predictions
func (ms *MemoryStorePrediction) GetPredictions(league *data.League, season *data.Season) ([]data.Prediction, error) {
	var predictions []data.Prediction
	for _, prediction := range ms.predictions {
		if prediction.League.ID == league.ID && prediction.Season.Year == season.Year {
			predictions = append(predictions, *prediction)
		}
	}
	return predictions, nil
}
