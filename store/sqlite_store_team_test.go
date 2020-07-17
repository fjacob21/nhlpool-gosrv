package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestNewSqliteStoreAddTeam(t *testing.T) {
	store := NewSqliteStore()
	assert.NotNil(t, store, "Should not be nil")
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.League().AddLeague(&league)
	venue := &data.Venue{ID: "id", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(venue)
	team := &data.Team{ID: "id", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: venue}
	err = store.Team().AddTeam(team)
	assert.NoError(t, err, "Should not have error")
	getTeam, err := store.Team().GetTeam(venue.ID, &league)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getTeam, "Should not be nil")
	assert.Equal(t, getTeam.ID, venue.ID, "Invalid ID")
	assert.Equal(t, getTeam.League.ID, league.ID, "Invalid league")
	assert.Equal(t, getTeam.Abbreviation, team.Abbreviation, "Invalid abbreviation")
	assert.Equal(t, getTeam.Name, team.Name, "Invalid Name")
	assert.Equal(t, getTeam.Fullname, team.Fullname, "Invalid Fullname")
	assert.Equal(t, getTeam.City, team.City, "Invalid City")
	assert.Equal(t, getTeam.Active, team.Active, "Invalid Active")
	assert.Equal(t, getTeam.CreationYear, team.CreationYear, "Invalid CreationYear")
	assert.Equal(t, getTeam.Website, team.Website, "Invalid Website")
	assert.Equal(t, getTeam.Venue.ID, team.Venue.ID, "Invalid Venue")
}

func TestNewSqliteStoreUpdateTeam(t *testing.T) {
	store := NewSqliteStore()
	assert.NotNil(t, store, "Should not be nil")
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.League().AddLeague(&league)
	venue := &data.Venue{ID: "id", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(venue)
	team := &data.Team{ID: "id", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: venue}
	err = store.Team().AddTeam(team)
	assert.NoError(t, err, "Should not have error")
	team.Abbreviation = "abbreviation2"
	team.Name = "name2"
	team.Fullname = "fullname2"
	team.City = "city2"
	team.Active = false
	team.CreationYear = "creationyear2"
	team.Website = "website2"
	err = store.Team().UpdateTeam(team)
	assert.Nil(t, err, "Should be nil")
	getTeam, err := store.Team().GetTeam(team.ID, &league)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getTeam, "Should not be nil")
	assert.Equal(t, getTeam.ID, venue.ID, "Invalid ID")
	assert.Equal(t, getTeam.League.ID, league.ID, "Invalid league")
	assert.Equal(t, getTeam.Abbreviation, team.Abbreviation, "Invalid abbreviation")
	assert.Equal(t, getTeam.Name, team.Name, "Invalid Name")
	assert.Equal(t, getTeam.Fullname, team.Fullname, "Invalid Fullname")
	assert.Equal(t, getTeam.City, team.City, "Invalid City")
	assert.Equal(t, getTeam.Active, team.Active, "Invalid Active")
	assert.Equal(t, getTeam.CreationYear, team.CreationYear, "Invalid CreationYear")
	assert.Equal(t, getTeam.Website, team.Website, "Invalid Website")
	assert.Equal(t, getTeam.Venue.ID, team.Venue.ID, "Invalid Venue")
}

func TestNewSqliteStoreDeleteTeam(t *testing.T) {
	store := NewSqliteStore()
	assert.NotNil(t, store, "Should not be nil")
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.League().AddLeague(&league)
	venue := &data.Venue{ID: "id", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(venue)
	team := &data.Team{ID: "id", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: venue}
	err = store.Team().AddTeam(team)
	assert.NoError(t, err, "Should not have error")
	getTeam, err := store.Team().GetTeam(venue.ID, &league)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getTeam, "Should not be nil")
	err = store.Team().DeleteTeam(team)
	assert.NoError(t, err, "Should not have error")
	getTeam, err = store.Team().GetTeam(team.ID, &league)
	assert.Error(t, err, "Should have error")
	assert.Nil(t, getTeam, "Should be nil")
}

func TestNewSqliteStoreGetTeamss(t *testing.T) {
	store := NewSqliteStore()
	assert.NotNil(t, store, "Should not be nil")
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	store.League().AddLeague(&league)
	venue := &data.Venue{ID: "id", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	store.Venue().AddVenue(venue)
	teams, _ := store.Team().GetTeams(&league)
	assert.Equal(t, len(teams), 0, "There should not have any team")

	team := &data.Team{ID: "id", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: venue}
	err := store.Team().AddTeam(team)
	assert.NoError(t, err, "Should not have error")

	teams, _ = store.Team().GetTeams(&league)
	assert.NotNil(t, teams, "Should not be nil")
	assert.Equal(t, len(teams), 1, "Should be only one team")
	assert.Equal(t, teams[0].ID, team.ID, "Should be the good team")
}
