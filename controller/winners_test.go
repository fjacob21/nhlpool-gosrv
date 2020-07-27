package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
	"nhlpool.com/service/go/nhlpool/store"
)

func TestAddWinner(t *testing.T) {
	store.Clean()
	requestPlayer := data.AddPlayerRequest{Name: "name", Email: "email", Admin: true, Password: "password"}
	replyPlayer := AddPlayer(requestPlayer)
	loginReq := data.LoginRequest{Password: "password"}
	Login(replyPlayer.Player.ID, loginReq)
	requestAddLeague := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	replyAddLeague := AddLeague(requestAddLeague)
	requestSeason := data.AddSeasonRequest{Year: 2000}
	replySeason := AddSeason(replyAddLeague.League.ID, requestSeason)
	requestTeam := data.AddTeamRequest{ID: "id", Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: data.AddTeamVenue{ID: "id", City: "city", Name: "name", Timezone: "timezone", Address: "address"}}
	replyTeam := AddTeam(replyAddLeague.League.ID, requestTeam)
	request := data.AddWinnerRequest{PlayerID: replyPlayer.Player.ID, Winner: replyTeam.Team.ID}
	reply := AddWinner(replyAddLeague.League.ID, replySeason.Season.Year, request)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, reply.Winner.League.ID, replyAddLeague.League.ID, "Invalid league")
	assert.Equal(t, reply.Winner.Season.Year, replySeason.Season.Year, "Invalid season")
	assert.Equal(t, reply.Winner.Player.ID, replyPlayer.Player.ID, "Invalid player")
	assert.Equal(t, reply.Winner.Winner.ID, replyTeam.Team.ID, "Invalid winner")
}

func TestGetWinners(t *testing.T) {
	store.Clean()
	requestPlayer := data.AddPlayerRequest{Name: "name", Email: "email", Admin: true, Password: "password"}
	replyPlayer := AddPlayer(requestPlayer)
	loginReq := data.LoginRequest{Password: "password"}
	Login(replyPlayer.Player.ID, loginReq)
	requestAddLeague := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	replyAddLeague := AddLeague(requestAddLeague)
	requestSeason := data.AddSeasonRequest{Year: 2000}
	replySeason := AddSeason(replyAddLeague.League.ID, requestSeason)
	requestTeam := data.AddTeamRequest{ID: "id", Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: data.AddTeamVenue{ID: "id", City: "city", Name: "name", Timezone: "timezone", Address: "address"}}
	replyTeam := AddTeam(replyAddLeague.League.ID, requestTeam)
	request := data.AddWinnerRequest{PlayerID: replyPlayer.Player.ID, Winner: replyTeam.Team.ID}
	AddWinner(replyAddLeague.League.ID, replySeason.Season.Year, request)
	reply := GetWinners(replyAddLeague.League.ID, replySeason.Season.Year)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, len(reply.Winners), 1, "Should be only one team")
}
