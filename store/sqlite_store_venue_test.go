package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestNewSqliteStoreAddVenue(t *testing.T) {
	store := NewSqliteStore()
	assert.NotNil(t, store, "Should not be nil")
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.League().AddLeague(&league)
	venue := &data.Venue{ID: "id", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(venue)
	assert.NoError(t, err, "Should not have error")
	getVenue, err := store.Venue().GetVenue(venue.ID, &league)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getVenue, "Should not be nil")
	assert.Equal(t, getVenue.ID, venue.ID, "Invalid ID")
	assert.Equal(t, getVenue.City, venue.City, "Invalid city")
	assert.Equal(t, getVenue.Name, venue.Name, "Invalid name")
	assert.Equal(t, getVenue.Timezone, venue.Timezone, "Invalid timezone")
	assert.Equal(t, getVenue.Address, venue.Address, "Invalid address")
}

func TestNewSqliteStoreUpdateVenue(t *testing.T) {
	store := NewSqliteStore()
	assert.NotNil(t, store, "Should not be nil")
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.League().AddLeague(&league)
	venue := &data.Venue{ID: "id", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(venue)
	assert.NoError(t, err, "Should not have error")
	venue.City = "city2"
	venue.Name = "name2"
	venue.Timezone = "timezone2"
	venue.Address = "address2"
	err = store.Venue().UpdateVenue(venue)
	assert.Nil(t, err, "Should be nil")
	getVenue, err := store.Venue().GetVenue(venue.ID, &league)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getVenue, "Should not be nil")
	assert.Equal(t, getVenue.ID, venue.ID, "Invalid ID")
	assert.Equal(t, getVenue.City, venue.City, "Invalid city")
	assert.Equal(t, getVenue.Name, venue.Name, "Invalid name")
	assert.Equal(t, getVenue.Timezone, venue.Timezone, "Invalid timezone")
	assert.Equal(t, getVenue.Address, venue.Address, "Invalid address")
}

func TestNewSqliteStoreDeleteVenue(t *testing.T) {
	store := NewSqliteStore()
	assert.NotNil(t, store, "Should not be nil")
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.League().AddLeague(&league)
	venue := &data.Venue{ID: "id", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(venue)
	assert.NoError(t, err, "Should not have error")
	getVenue, err := store.Venue().GetVenue(venue.ID, &league)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getVenue, "Should not be nil")
	err = store.Venue().DeleteVenue(venue)
	assert.NoError(t, err, "Should not have error")
	getVenue, err = store.Venue().GetVenue(venue.ID, &league)
	assert.Error(t, err, "Should have error")
	assert.Nil(t, getVenue, "Should be nil")
}
