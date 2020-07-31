package store

import "nhlpool.com/service/go/nhlpool/data"

// ConferenceStore interface of conference storage
type ConferenceStore interface {
	Clean() error
	GetConferences(league *data.League) ([]*data.Conference, error)
	GetConference(id string, league *data.League) (*data.Conference, error)
	AddConference(conference *data.Conference) error
	DeleteConference(conference *data.Conference) error
}
