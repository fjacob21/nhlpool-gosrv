package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
	"nhlpool.com/service/go/nhlpool/store"
)

func TestAddStanding(t *testing.T) {
	store.Clean()
	requestAddLeague := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	replyAddLeague := AddLeague(requestAddLeague)
	requestSeason := data.AddSeasonRequest{Year: 2000}
	replySeason := AddSeason(replyAddLeague.League.ID, requestSeason)
	requestTeam := data.AddTeamRequest{ID: "id", Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: data.AddTeamVenue{ID: "id", City: "city", Name: "name", Timezone: "timezone", Address: "address"}}
	replyTeam := AddTeam(replyAddLeague.League.ID, requestTeam)
	request := data.AddStandingRequest{TeamID: replyTeam.Team.ID, Points: 0, Win: 1, Losses: 2, OT: 3, GamesPlayed: 4, GoalsAgainst: 5, GoalsScored: 6, Ranks: 7}
	reply := AddStanding(replyAddLeague.League.ID, replySeason.Season.Year, request)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, reply.Standing.League.ID, replyAddLeague.League.ID, "Invalid league")
	assert.Equal(t, reply.Standing.Season.Year, replySeason.Season.Year, "Invalid season")
	assert.Equal(t, reply.Standing.Team.ID, replyTeam.Team.ID, "Invalid team")
	assert.Equal(t, reply.Standing.Points, 0, "Invalid Point")
	assert.Equal(t, reply.Standing.Win, 1, "Invalid Win")
	assert.Equal(t, reply.Standing.Losses, 2, "Invalid Losses")
	assert.Equal(t, reply.Standing.OT, 3, "Invalid OT")
	assert.Equal(t, reply.Standing.GamesPlayed, 4, "Invalid GamesPlayed")
	assert.Equal(t, reply.Standing.GoalsAgainst, 5, "Invalid GoalsAgainst")
	assert.Equal(t, reply.Standing.GoalsScored, 6, "Invalid GoalsScored")
	assert.Equal(t, reply.Standing.Ranks, 7, "Invalid Ranks")
}

func TestGetStandings(t *testing.T) {
	store.Clean()
	requestAddLeague := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	replyAddLeague := AddLeague(requestAddLeague)
	requestSeason := data.AddSeasonRequest{Year: 2000}
	replySeason := AddSeason(replyAddLeague.League.ID, requestSeason)
	requestTeam := data.AddTeamRequest{ID: "id", Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: data.AddTeamVenue{ID: "id", City: "city", Name: "name", Timezone: "timezone", Address: "address"}}
	AddTeam(replyAddLeague.League.ID, requestTeam)
	request := data.AddStandingRequest{TeamID: "id", Points: 0, Win: 1, Losses: 2, OT: 3, GamesPlayed: 4, GoalsAgainst: 5, GoalsScored: 6, Ranks: 7}
	AddStanding(replyAddLeague.League.ID, replySeason.Season.Year, request)
	reply := GetStandings(replyAddLeague.League.ID, replySeason.Season.Year)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, len(reply.Standings), 1, "Should be only one team")
}
