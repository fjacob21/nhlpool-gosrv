package store

import (
	"errors"

	"nhlpool.com/service/go/nhlpool/data"
)

// MemoryStoreWinner Is a Winner data store that keep it only in memory
type MemoryStoreWinner struct {
	winners map[string]*data.Winner
}

// NewMemoryStoreWinner Create a new winner memory store
func NewMemoryStoreWinner() *MemoryStoreWinner {
	store := &MemoryStoreWinner{}
	store.winners = make(map[string]*data.Winner)
	return store
}

// Clean Empty the store
func (ms *MemoryStoreWinner) Clean() error {
	ms.winners = make(map[string]*data.Winner)
	return nil
}

// AddWinner Add a new winner
func (ms *MemoryStoreWinner) AddWinner(winner *data.Winner) error {
	_, ok := ms.winners[winner.League.ID+string(winner.Season.Year)+winner.Player.ID]
	if ok {
		return errors.New("Winner already exist")
	}
	ms.winners[winner.League.ID+string(winner.Season.Year)+winner.Player.ID] = winner
	return nil
}

// UpdateWinner Update a winner info
func (ms *MemoryStoreWinner) UpdateWinner(winner *data.Winner) error {
	_, ok := ms.winners[winner.League.ID+string(winner.Season.Year)+winner.Player.ID]
	if !ok {
		return errors.New("Winner do not exist")
	}
	ms.winners[winner.League.ID+string(winner.Season.Year)+winner.Player.ID] = winner
	return nil
}

// DeleteWinner Delete a winner
func (ms *MemoryStoreWinner) DeleteWinner(winner *data.Winner) error {
	_, ok := ms.winners[winner.League.ID+string(winner.Season.Year)+winner.Player.ID]
	if !ok {
		return errors.New("Winner do not exist")
	}
	delete(ms.winners, winner.League.ID+string(winner.Season.Year)+winner.Player.ID)
	return nil
}

// GetWinner Get a winner
func (ms *MemoryStoreWinner) GetWinner(player *data.Player, league *data.League, season *data.Season) (*data.Winner, error) {
	venue, ok := ms.winners[league.ID+string(season.Year)+player.ID]
	if !ok {
		return nil, errors.New("Winner do not exist")
	}
	return venue, nil
}

// GetWinners Get all winners
func (ms *MemoryStoreWinner) GetWinners(league *data.League, season *data.Season) ([]data.Winner, error) {
	var winners []data.Winner
	for _, winner := range ms.winners {
		if winner.League.ID == league.ID && winner.Season.Year == season.Year {
			winners = append(winners, *winner)
		}
	}
	return winners, nil
}
