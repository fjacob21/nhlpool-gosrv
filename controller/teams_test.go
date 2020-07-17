package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
	"nhlpool.com/service/go/nhlpool/store"
)

func TestAddTeam(t *testing.T) {
	store.Clean()
	requestAddLeague := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	replyAddLeague := AddLeague(requestAddLeague)
	request := data.AddTeamRequest{ID: "id", Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: data.AddTeamVenue{ID: "id", City: "city", Name: "name", Timezone: "timezone", Address: "address"}}
	reply := AddTeam(replyAddLeague.League.ID, request)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, reply.Team.ID, "id", "Invalid id")
	assert.Equal(t, reply.Team.Abbreviation, "abbreviation", "Invalid email")
	assert.Equal(t, reply.Team.Name, "name", "Invalid admin")
	assert.Equal(t, reply.Team.Fullname, "fullname", "Password should be empty")
	assert.Equal(t, reply.Team.City, "city", "Password should be empty")
	assert.Equal(t, reply.Team.Active, true, "Active should be empty")
	assert.Equal(t, reply.Team.CreationYear, "creationyear", "CreationYear should be empty")
	assert.Equal(t, reply.Team.Website, "website", "Website should be empty")
	assert.NotNil(t, reply.Team.Venue, "Should not be nil")
	assert.Equal(t, reply.Team.Venue.ID, "id", "Invalid venue id")
	assert.Equal(t, reply.Team.Venue.City, "city", "Invalid venue City")
	assert.Equal(t, reply.Team.Venue.Name, "name", "Invalid venue Name")
	assert.Equal(t, reply.Team.Venue.Timezone, "timezone", "Invalid venue Timezone")
	assert.Equal(t, reply.Team.Venue.Address, "address", "Invalid venue Address")
}

func TestGetTeams(t *testing.T) {
	store.Clean()
	requestAddLeague := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	replyAddLeague := AddLeague(requestAddLeague)
	request := data.AddTeamRequest{ID: "id", Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: data.AddTeamVenue{ID: "id", City: "city", Name: "name", Timezone: "timezone", Address: "address"}}
	AddTeam(replyAddLeague.League.ID, request)
	reply := GetTeams(replyAddLeague.League.ID)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, len(reply.Teams), 1, "Should be only one team")

}
