package controller

import (
	"testing"

	"nhlpool.com/service/go/nhlpool/store"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestGetWinner(t *testing.T) {
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
	reply := GetWinner(replyAddLeague.League.ID, replySeason.Season.Year, replyPlayer.Player.ID)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, reply.Winner.League.ID, replyAddLeague.League.ID, "Invalid league")
	assert.Equal(t, reply.Winner.Season.Year, replySeason.Season.Year, "Invalid season")
	assert.Equal(t, reply.Winner.Player.ID, replyPlayer.Player.ID, "Invalid player")
	assert.Equal(t, reply.Winner.Winner.ID, replyTeam.Team.ID, "Invalid winner")
}

func TestDeleteWinner(t *testing.T) {
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
	request := data.AddWinnerRequest{PlayerID: replyPlayer.Player.ID, Winner: replyTeam.Team.ID}
	AddWinner(replyAddLeague.League.ID, replySeason.Season.Year, request)
	deleteReq := data.DeleteWinnerRequest{SessionID: loginReply.SessionID}
	deleteReply := DeleteWinner(replyAddLeague.League.ID, replySeason.Season.Year, replyPlayer.Player.ID, deleteReq)
	assert.Equal(t, deleteReply.Result.Code, data.SUCCESS, "Should be a success")
	reply := GetWinners(replyAddLeague.League.ID, replySeason.Season.Year)
	assert.Equal(t, len(reply.Winners), 0, "Should have zero player")
}

func TestEditWinner(t *testing.T) {
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
	requestTeam2 := data.AddTeamRequest{ID: "id2", Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: data.AddTeamVenue{ID: "id", City: "city", Name: "name", Timezone: "timezone", Address: "address"}}
	replyTeam2 := AddTeam(replyAddLeague.League.ID, requestTeam2)
	request := data.AddWinnerRequest{PlayerID: replyPlayer.Player.ID, Winner: replyTeam.Team.ID}
	AddWinner(replyAddLeague.League.ID, replySeason.Season.Year, request)
	editReq := data.EditWinnerRequest{SessionID: loginReply.SessionID, Winner: replyTeam2.Team.ID}
	editReply := EditWinner(replyAddLeague.League.ID, replySeason.Season.Year, replyPlayer.Player.ID, editReq)
	assert.Equal(t, editReply.Result.Code, data.SUCCESS, "Should be a success")
	reply := GetWinners(replyAddLeague.League.ID, replySeason.Season.Year)
	assert.Equal(t, len(reply.Winners), 1, "Should have zero player")
	assert.Equal(t, reply.Winners[0].League.ID, replyAddLeague.League.ID, "Invalid league")
	assert.Equal(t, reply.Winners[0].Season.Year, replySeason.Season.Year, "Invalid season")
	assert.Equal(t, reply.Winners[0].Player.ID, replyPlayer.Player.ID, "Invalid player")
	assert.Equal(t, reply.Winners[0].Winner.ID, replyTeam2.Team.ID, "Invalid Winner")
}
