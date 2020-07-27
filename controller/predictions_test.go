package controller

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
	"nhlpool.com/service/go/nhlpool/store"
)

func TestAddPrediction(t *testing.T) {
	store.Clean()
	requestPlayer := data.AddPlayerRequest{Name: "name", Email: "email", Admin: true, Password: "password"}
	replyPlayer := AddPlayer(requestPlayer)
	loginReq := data.LoginRequest{Password: "password"}
	Login(replyPlayer.Player.ID, loginReq)
	requestAddLeague := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	replyAddLeague := AddLeague(requestAddLeague)
	requestSeason := data.AddSeasonRequest{Year: 2000}
	replySeason := AddSeason(replyAddLeague.League.ID, requestSeason)
	requestHomeTeam := data.AddTeamRequest{ID: "homeid", Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: data.AddTeamVenue{ID: "id", City: "city", Name: "name", Timezone: "timezone", Address: "address"}}
	replyHomeTeam := AddTeam(replyAddLeague.League.ID, requestHomeTeam)
	requestAwayTeam := data.AddTeamRequest{ID: "awayid", Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: data.AddTeamVenue{ID: "id", City: "city", Name: "name", Timezone: "timezone", Address: "address"}}
	replyAwayTeam := AddTeam(replyAddLeague.League.ID, requestAwayTeam)
	start := time.Now()
	requestMatchup := data.AddMatchupRequest{ID: "id", HomeID: replyHomeTeam.Team.ID, AwayID: replyAwayTeam.Team.ID, Round: 1, Start: start.Format(time.RFC3339)}
	replyMatchup := AddMatchup(replyAddLeague.League.ID, replySeason.Season.Year, requestMatchup)
	request := data.AddPredictionRequest{PlayerID: replyPlayer.Player.ID, MatchupID: replyMatchup.Matchup.ID, Winner: replyHomeTeam.Team.ID, Games: 4}
	reply := AddPrediction(replyAddLeague.League.ID, replySeason.Season.Year, request)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, reply.Prediction.League.ID, replyAddLeague.League.ID, "Invalid league")
	assert.Equal(t, reply.Prediction.Season.Year, replySeason.Season.Year, "Invalid season")
	assert.Equal(t, reply.Prediction.Player.ID, replyPlayer.Player.ID, "Invalid player")
	assert.Equal(t, reply.Prediction.Matchup.ID, replyMatchup.Matchup.ID, "Invalid matchup")
	assert.Equal(t, reply.Prediction.Winner.ID, replyHomeTeam.Team.ID, "Invalid winner")
	assert.Equal(t, reply.Prediction.Games, 4, "Invalid games")
}

func TestGetPredictions(t *testing.T) {
	store.Clean()
	requestPlayer := data.AddPlayerRequest{Name: "name", Email: "email", Admin: true, Password: "password"}
	replyPlayer := AddPlayer(requestPlayer)
	loginReq := data.LoginRequest{Password: "password"}
	Login(replyPlayer.Player.ID, loginReq)
	requestAddLeague := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	replyAddLeague := AddLeague(requestAddLeague)
	requestSeason := data.AddSeasonRequest{Year: 2000}
	replySeason := AddSeason(replyAddLeague.League.ID, requestSeason)
	requestHomeTeam := data.AddTeamRequest{ID: "homeid", Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: data.AddTeamVenue{ID: "id", City: "city", Name: "name", Timezone: "timezone", Address: "address"}}
	replyHomeTeam := AddTeam(replyAddLeague.League.ID, requestHomeTeam)
	requestAwayTeam := data.AddTeamRequest{ID: "awayid", Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: data.AddTeamVenue{ID: "id", City: "city", Name: "name", Timezone: "timezone", Address: "address"}}
	replyAwayTeam := AddTeam(replyAddLeague.League.ID, requestAwayTeam)
	start := time.Now()
	requestMatchup := data.AddMatchupRequest{ID: "id", HomeID: replyHomeTeam.Team.ID, AwayID: replyAwayTeam.Team.ID, Round: 1, Start: start.Format(time.RFC3339)}
	replyMatchup := AddMatchup(replyAddLeague.League.ID, replySeason.Season.Year, requestMatchup)
	request := data.AddPredictionRequest{PlayerID: replyPlayer.Player.ID, MatchupID: replyMatchup.Matchup.ID, Winner: replyHomeTeam.Team.ID, Games: 4}
	replyPrediction := AddPrediction(replyAddLeague.League.ID, replySeason.Season.Year, request)
	assert.NotNil(t, replyPrediction, "Should not be nil")
	reply := GetPredictions(replyAddLeague.League.ID, replySeason.Season.Year)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, len(reply.Predictions), 1, "Should be only one team")
}
