package store

import "nhlpool.com/service/go/nhlpool/data"

// DivisionStore interface of division storage
type DivisionStore interface {
	Clean() error
	GetDivisions(league *data.League) ([]*data.Division, error)
	GetDivision(id string, league *data.League) (*data.Division, error)
	AddDivision(division *data.Division) error
	DeleteDivision(division *data.Division) error
}
