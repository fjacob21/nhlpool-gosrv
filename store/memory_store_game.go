package store

import (
	"errors"
	"time"

	"nhlpool.com/service/go/nhlpool/data"
)

// MemoryStoreGame Is a Game data store that keep it only in memory
type MemoryStoreGame struct {
	games map[string]*data.Game
}

// NewMemoryStoreGame Create a new game memory store
func NewMemoryStoreGame() *MemoryStoreGame {
	store := &MemoryStoreGame{}
	store.games = make(map[string]*data.Game)
	return store
}

// Clean Empty the store
func (ms *MemoryStoreGame) Clean() error {
	ms.games = make(map[string]*data.Game)
	return nil
}

// AddGame Add a new game
func (ms *MemoryStoreGame) AddGame(game *data.Game) error {
	_, ok := ms.games[game.League.ID+string(game.Season.Year)+game.Home.ID+game.Away.ID+game.Date.Format(time.RFC3339)]
	if ok {
		return errors.New("Game already exist")
	}
	ms.games[game.League.ID+string(game.Season.Year)+game.Home.ID+game.Away.ID+game.Date.Format(time.RFC3339)] = game
	return nil
}

// UpdateGame Update a game info
func (ms *MemoryStoreGame) UpdateGame(game *data.Game) error {
	_, ok := ms.games[game.League.ID+string(game.Season.Year)+game.Home.ID+game.Away.ID+game.Date.Format(time.RFC3339)]
	if !ok {
		return errors.New("Game do not exist")
	}
	ms.games[game.League.ID+string(game.Season.Year)+game.Home.ID+game.Away.ID+game.Date.Format(time.RFC3339)] = game
	return nil
}

// DeleteGame Delete a game
func (ms *MemoryStoreGame) DeleteGame(game *data.Game) error {
	_, ok := ms.games[game.League.ID+string(game.Season.Year)+game.Home.ID+game.Away.ID+game.Date.Format(time.RFC3339)]
	if !ok {
		return errors.New("Game do not exist")
	}
	delete(ms.games, game.League.ID+string(game.Season.Year)+game.Home.ID+game.Away.ID+game.Date.Format(time.RFC3339))
	return nil
}

// GetGame Get a game
func (ms *MemoryStoreGame) GetGame(league *data.League, season *data.Season, home *data.Team, away *data.Team, date time.Time) (*data.Game, error) {
	venue, ok := ms.games[league.ID+string(season.Year)+home.ID+away.ID+date.Format(time.RFC3339)]
	if !ok {
		return nil, errors.New("Game do not exist")
	}
	return venue, nil
}

// GetGames Get all games
func (ms *MemoryStoreGame) GetGames(league *data.League, season *data.Season, home *data.Team, away *data.Team) ([]data.Game, error) {
	var games []data.Game
	for _, game := range ms.games {
		if game.League.ID == league.ID && game.Season.Year == season.Year && game.Home.ID == home.ID && game.Away.ID == away.ID {
			games = append(games, *game)
		}
	}
	return games, nil
}

// GetSeasonGames Get all games
func (ms *MemoryStoreGame) GetSeasonGames(league *data.League, season *data.Season) ([]data.Game, error) {
	var games []data.Game
	for _, game := range ms.games {
		if game.League.ID == league.ID && game.Season.Year == season.Year {
			games = append(games, *game)
		}
	}
	return games, nil
}
