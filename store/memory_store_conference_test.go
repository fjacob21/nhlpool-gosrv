package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestNewMemoryStoreAddConference(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.League().AddLeague(&league)
	conference := &data.Conference{ID: "id", League: league, Name: "name"}
	err = store.Conference().AddConference(conference)
	assert.NoError(t, err, "Should not have error")
	getConference, err := store.Conference().GetConference(conference.ID, &league)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getConference, "Should not be nil")
	assert.Equal(t, getConference.ID, conference.ID, "Invalid ID")
	assert.Equal(t, getConference.League.ID, league.ID, "Invalid league")
	assert.Equal(t, getConference.Name, conference.Name, "Invalid name")
}

func TestNewMemoryStoreDeleteConference(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.League().AddLeague(&league)
	conference := &data.Conference{ID: "id", League: league, Name: "name"}
	err = store.Conference().AddConference(conference)
	assert.NoError(t, err, "Should not have error")
	getConference, err := store.Conference().GetConference(conference.ID, &league)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getConference, "Should not be nil")
	err = store.Conference().DeleteConference(conference)
	assert.NoError(t, err, "Should not have error")
	getConference, err = store.Conference().GetConference(conference.ID, &league)
	assert.Error(t, err, "Should have error")
	assert.Nil(t, getConference, "Should be nil")
}
