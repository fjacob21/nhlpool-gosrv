package store

import (
	"time"

	"nhlpool.com/service/go/nhlpool/data"
)

// GameStore interface of game storage
type GameStore interface {
	Clean() error
	GetSeasonGames(league *data.League, season *data.Season, home *data.Team, away *data.Team) ([]data.Game, error)
	GetPlayoffGames(league *data.League, season *data.Season, home *data.Team, away *data.Team) ([]data.Game, error)
	GetAllGames(league *data.League, season *data.Season) ([]data.Game, error)
	GetGames(league *data.League, season *data.Season, home *data.Team, away *data.Team) ([]data.Game, error)
	GetGame(league *data.League, season *data.Season, home *data.Team, away *data.Team, date time.Time) (*data.Game, error)
	AddGame(game *data.Game) error
	UpdateGame(game *data.Game) error
	DeleteGame(game *data.Game) error
}
