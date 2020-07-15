package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestNewMemoryStoreAddLeague(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	league := &data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.League().AddLeague(league)
	assert.NoError(t, err, "Should not have error")
	getLeague, err := store.League().GetLeague(league.ID)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getLeague, "Should not be nil")
	assert.Equal(t, getLeague.ID, league.ID, "Invalid ID")
	assert.Equal(t, getLeague.Name, league.Name, "Invalid name")
	assert.Equal(t, getLeague.Description, league.Description, "Invalid description")
	assert.Equal(t, getLeague.Website, league.Website, "Invalid we site")
}

func TestNewMemoryStoreUpdateLeague(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	league := &data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.League().AddLeague(league)
	assert.NoError(t, err, "Should not have error")
	league.Name = "name2"
	league.Description = "description2"
	league.Website = "website2"
	err = store.League().UpdateLeague(league)
	assert.Nil(t, err, "Should be nil")
	getLeague, err := store.League().GetLeague(league.ID)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getLeague, "Should not be nil")
	assert.Equal(t, getLeague.ID, "id", "Invalid ID")
	assert.Equal(t, getLeague.Name, "name2", "Invalid name")
	assert.Equal(t, getLeague.Description, "description2", "Invalid description")
	assert.Equal(t, getLeague.Website, "website2", "Invalid we site")
}

func TestNewMemoryStoreDeleteLeague(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	league := &data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.League().AddLeague(league)
	assert.NoError(t, err, "Should not have error")
	getLeague, err := store.League().GetLeague(league.ID)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getLeague, "Should not be nil")
	err = store.League().DeleteLeague(league)
	assert.NoError(t, err, "Should not have error")
	getLeague, err = store.League().GetLeague(league.ID)
	assert.Error(t, err, "Should have error")
	assert.Nil(t, getLeague, "Should be nil")
}

func TestNewMemoryStoreGetLeagues(t *testing.T) {
	store := NewMemoryStore()
	assert.NotNil(t, store, "Should not be nil")
	league := &data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	leagues, _ := store.League().GetLeagues()
	assert.Equal(t, len(leagues), 0, "There should not have any leagues")
	err := store.League().AddLeague(league)
	assert.NoError(t, err, "Should not have error")
	leagues, _ = store.League().GetLeagues()
	assert.Equal(t, len(leagues), 1, "There should not have any leagues")
	assert.Equal(t, leagues[0].ID, league.ID, "Should be the good league")
}
