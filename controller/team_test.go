package controller

import (
	"testing"

	"nhlpool.com/service/go/nhlpool/store"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestGetTeam(t *testing.T) {
	store.Clean()
	requestAddLeague := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	replyAddLeague := AddLeague(requestAddLeague)
	request := data.AddTeamRequest{ID: "id", Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: data.AddTeamVenue{ID: "id", City: "city", Name: "name", Timezone: "timezone", Address: "address"}}
	reply := AddTeam(replyAddLeague.League.ID, request)
	assert.NotNil(t, reply, "Should not be nil")
	teamReply := GetTeam(replyAddLeague.League.ID, "id")
	assert.NotNil(t, teamReply, "Should not be nil")
	assert.Equal(t, teamReply.Team.ID, reply.Team.ID, "Invalid ID")
	assert.Equal(t, teamReply.Team.Abbreviation, "abbreviation", "Invalid abbreviation")
	assert.Equal(t, teamReply.Team.Name, "name", "Invalid name")
	assert.Equal(t, teamReply.Team.Fullname, "fullname", "Invalid fullname")
	assert.Equal(t, teamReply.Team.City, "city", "Invalid city")
	assert.Equal(t, teamReply.Team.Active, true, "Invalid active")
	assert.Equal(t, teamReply.Team.CreationYear, "creationyear", "Invalid creationyear")
	assert.Equal(t, teamReply.Team.Website, "website", "Invalid website")
	assert.NotNil(t, teamReply.Team.Venue, "Should not be nil")
	assert.Equal(t, teamReply.Team.Venue.ID, "id", "Invalid venue id")
	assert.Equal(t, teamReply.Team.Venue.City, "city", "Invalid venue City")
	assert.Equal(t, teamReply.Team.Venue.Name, "name", "Invalid venue Name")
	assert.Equal(t, teamReply.Team.Venue.Timezone, "timezone", "Invalid venue Timezone")
	assert.Equal(t, teamReply.Team.Venue.Address, "address", "Invalid venue Address")
}
