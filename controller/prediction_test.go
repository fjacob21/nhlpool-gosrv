package controller

import (
	"testing"
	"time"

	"nhlpool.com/service/go/nhlpool/store"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestGetPrediction(t *testing.T) {
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
	reply := GetPrediction(replyAddLeague.League.ID, replySeason.Season.Year, replyPlayer.Player.ID, replyMatchup.Matchup.ID)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, reply.Prediction.League.ID, replyAddLeague.League.ID, "Invalid league")
	assert.Equal(t, reply.Prediction.Season.Year, replySeason.Season.Year, "Invalid season")
	assert.Equal(t, reply.Prediction.Player.ID, replyPlayer.Player.ID, "Invalid player")
	assert.Equal(t, reply.Prediction.Matchup.ID, replyMatchup.Matchup.ID, "Invalid matchup")
	assert.Equal(t, reply.Prediction.Winner.ID, replyHomeTeam.Team.ID, "Invalid winner")
	assert.Equal(t, reply.Prediction.Games, 4, "Invalid games")
}

func TestDeletePrediction(t *testing.T) {
	store.Clean()
	requestPlayer := data.AddPlayerRequest{Name: "name", Email: "email", Admin: true, Password: "password"}
	replyPlayer := AddPlayer(requestPlayer)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(replyPlayer.Player.ID, loginReq)
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
	deleteReq := data.DeletePredictionRequest{SessionID: loginReply.SessionID}
	deleteReply := DeletePrediction(replyAddLeague.League.ID, replySeason.Season.Year, replyPlayer.Player.ID, replyMatchup.Matchup.ID, deleteReq)
	assert.Equal(t, deleteReply.Result.Code, data.SUCCESS, "Should be a success")
	reply := GetPredictions(replyAddLeague.League.ID, replySeason.Season.Year)
	assert.Equal(t, len(reply.Predictions), 0, "Should have zero player")
}

func TestEditPrediction(t *testing.T) {
	store.Clean()
	requestPlayer := data.AddPlayerRequest{Name: "name", Email: "email", Admin: true, Password: "password"}
	replyPlayer := AddPlayer(requestPlayer)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(replyPlayer.Player.ID, loginReq)
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
	editReq := data.EditPredictionRequest{SessionID: loginReply.SessionID, Winner: replyAwayTeam.Team.ID, Games: 5}
	editReply := EditPrediction(replyAddLeague.League.ID, replySeason.Season.Year, replyPlayer.Player.ID, replyMatchup.Matchup.ID, editReq)
	assert.Equal(t, editReply.Result.Code, data.SUCCESS, "Should be a success")
	reply := GetPredictions(replyAddLeague.League.ID, replySeason.Season.Year)
	assert.Equal(t, len(reply.Predictions), 1, "Should have zero player")
	assert.Equal(t, reply.Predictions[0].League.ID, replyAddLeague.League.ID, "Invalid league")
	assert.Equal(t, reply.Predictions[0].Season.Year, replySeason.Season.Year, "Invalid season")
	assert.Equal(t, reply.Predictions[0].Player.ID, replyPlayer.Player.ID, "Invalid player")
	assert.Equal(t, reply.Predictions[0].Matchup.ID, replyMatchup.Matchup.ID, "Invalid matchup")
	assert.Equal(t, reply.Predictions[0].Winner.ID, replyAwayTeam.Team.ID, "Invalid winner")
	assert.Equal(t, reply.Predictions[0].Games, 5, "Invalid games")
}

func TestEditPredictionEmpty(t *testing.T) {
	store.Clean()
	requestPlayer := data.AddPlayerRequest{Name: "name", Email: "email", Admin: true, Password: "password"}
	replyPlayer := AddPlayer(requestPlayer)
	loginReq := data.LoginRequest{Password: "password"}
	loginReply := Login(replyPlayer.Player.ID, loginReq)
	requestAddLeague := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	replyAddLeague := AddLeague(requestAddLeague)
	requestSeason := data.AddSeasonRequest{Year: 2000}
	replySeason := AddSeason(replyAddLeague.League.ID, requestSeason)
	requestMatchup := data.AddMatchupRequest{ID: "id", HomeID: "", AwayID: "", Round: 1, Start: ""}
	replyMatchup := AddMatchup(replyAddLeague.League.ID, replySeason.Season.Year, requestMatchup)
	request := data.AddPredictionRequest{PlayerID: replyPlayer.Player.ID, MatchupID: replyMatchup.Matchup.ID, Winner: "", Games: 4}
	replyPrediction := AddPrediction(replyAddLeague.League.ID, replySeason.Season.Year, request)
	assert.NotNil(t, replyPrediction, "Should not be nil")
	editReq := data.EditPredictionRequest{SessionID: loginReply.SessionID, Winner: "", Games: 5}
	editReply := EditPrediction(replyAddLeague.League.ID, replySeason.Season.Year, replyPlayer.Player.ID, replyMatchup.Matchup.ID, editReq)
	assert.Equal(t, editReply.Result.Code, data.SUCCESS, "Should be a success")
	reply := GetPredictions(replyAddLeague.League.ID, replySeason.Season.Year)
	assert.Equal(t, len(reply.Predictions), 1, "Should have zero player")
	assert.Equal(t, reply.Predictions[0].League.ID, replyAddLeague.League.ID, "Invalid league")
	assert.Equal(t, reply.Predictions[0].Season.Year, replySeason.Season.Year, "Invalid season")
	assert.Equal(t, reply.Predictions[0].Player.ID, replyPlayer.Player.ID, "Invalid player")
	assert.Equal(t, reply.Predictions[0].Matchup.ID, replyMatchup.Matchup.ID, "Invalid matchup")
	assert.Equal(t, reply.Predictions[0].Winner.ID, "", "Invalid winner")
	assert.Equal(t, reply.Predictions[0].Games, 5, "Invalid games")
}
