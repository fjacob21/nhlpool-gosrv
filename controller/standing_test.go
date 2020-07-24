package controller

import (
	"testing"

	"nhlpool.com/service/go/nhlpool/store"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestGetStanding(t *testing.T) {
	store.Clean()
	requestAddLeague := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	replyAddLeague := AddLeague(requestAddLeague)
	requestSeason := data.AddSeasonRequest{Year: 2000}
	replySeason := AddSeason(replyAddLeague.League.ID, requestSeason)
	requestTeam := data.AddTeamRequest{ID: "id", Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: data.AddTeamVenue{ID: "id", City: "city", Name: "name", Timezone: "timezone", Address: "address"}}
	replyTeam := AddTeam(replyAddLeague.League.ID, requestTeam)
	request := data.AddStandingRequest{TeamID: "id", Points: 0, Win: 1, Losses: 2, OT: 3, GamesPlayed: 4, GoalsAgainst: 5, GoalsScored: 6, Ranks: 7}
	AddStanding(replyAddLeague.League.ID, replySeason.Season.Year, request)
	reply := GetStanding(replyAddLeague.League.ID, replySeason.Season.Year, replyTeam.Team.ID)
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

func TestDeleteStanding(t *testing.T) {
	store.Clean()
	requestPlayer := data.AddPlayerRequest{Name: "name", Email: "email", Admin: true, Password: "password"}
	replyPlayer := AddPlayer(requestPlayer)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(replyPlayer.Player.ID, loginReq)
	requestAddLeague := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	replyAddLeague := AddLeague(requestAddLeague)
	requestSeason := data.AddSeasonRequest{Year: 2000}
	replySeason := AddSeason(replyAddLeague.League.ID, requestSeason)
	requestTeam := data.AddTeamRequest{ID: "id", Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: data.AddTeamVenue{ID: "id", City: "city", Name: "name", Timezone: "timezone", Address: "address"}}
	replyTeam := AddTeam(replyAddLeague.League.ID, requestTeam)
	request := data.AddStandingRequest{TeamID: "id", Points: 0, Win: 1, Losses: 2, OT: 3, GamesPlayed: 4, GoalsAgainst: 5, GoalsScored: 6, Ranks: 7}
	AddStanding(replyAddLeague.League.ID, replySeason.Season.Year, request)
	deleteReq := data.DeleteStandingRequest{SessionID: loginReply.SessionID}
	deleteReply := DeleteStanding(replyAddLeague.League.ID, replySeason.Season.Year, replyTeam.Team.ID, deleteReq)
	assert.Equal(t, deleteReply.Result.Code, data.SUCCESS, "Should be a success")
	reply := GetStandings(replyAddLeague.League.ID, replySeason.Season.Year)
	assert.Equal(t, len(reply.Standings), 0, "Should have zero player")
}
