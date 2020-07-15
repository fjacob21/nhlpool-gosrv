package store

import "nhlpool.com/service/go/nhlpool/data"

// Player interface of player storage
type Player interface {
	Clean() error
	GetPlayers() ([]data.Player, error)
	GetPlayer(id string) *data.Player
	AddPlayer(player *data.Player) error
	UpdatePlayer(player *data.Player) error
	DeletePlayer(player *data.Player) error
}
