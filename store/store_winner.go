package store

import "nhlpool.com/service/go/nhlpool/data"

// WinnerStore interface of winner storage
type WinnerStore interface {
	Clean() error
	GetWinners(league *data.League, season *data.Season) ([]data.Winner, error)
	GetWinner(player *data.Player, league *data.League, season *data.Season) (*data.Winner, error)
	AddWinner(standing *data.Winner) error
	UpdateWinner(standing *data.Winner) error
	DeleteWinner(standing *data.Winner) error
}
