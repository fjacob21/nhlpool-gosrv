package controller

import (
	"testing"
	"time"

	"nhlpool.com/service/go/nhlpool/store"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestGetMatchup(t *testing.T) {
	store.Clean()
	requestAddLeague := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	replyAddLeague := AddLeague(requestAddLeague)
	requestSeason := data.AddSeasonRequest{Year: 2000}
	replySeason := AddSeason(replyAddLeague.League.ID, requestSeason)
	requestHomeTeam := data.AddTeamRequest{ID: "homeid", Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: data.AddTeamVenue{ID: "id", City: "city", Name: "name", Timezone: "timezone", Address: "address"}}
	replyHomeTeam := AddTeam(replyAddLeague.League.ID, requestHomeTeam)
	requestAwayTeam := data.AddTeamRequest{ID: "awayid", Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: data.AddTeamVenue{ID: "id", City: "city", Name: "name", Timezone: "timezone", Address: "address"}}
	replyAwayTeam := AddTeam(replyAddLeague.League.ID, requestAwayTeam)
	start := time.Now()
	request := data.AddMatchupRequest{ID: "id", HomeID: replyHomeTeam.Team.ID, AwayID: replyAwayTeam.Team.ID, Round: 1, Start: start.Format(time.RFC3339)}
	replyMatchup := AddMatchup(replyAddLeague.League.ID, replySeason.Season.Year, request)
	reply := GetMatchup(replyAddLeague.League.ID, replySeason.Season.Year, replyMatchup.Matchup.ID)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, reply.Matchup.League.ID, replyAddLeague.League.ID, "Invalid league")
	assert.Equal(t, reply.Matchup.Season.Year, replySeason.Season.Year, "Invalid season")
	assert.Equal(t, reply.Matchup.ID, "id", "Invalid id")
	assert.Equal(t, reply.Matchup.Home.ID, replyHomeTeam.Team.ID, "Invalid Home")
	assert.Equal(t, reply.Matchup.Away.ID, replyAwayTeam.Team.ID, "Invalid away")
	assert.Equal(t, reply.Matchup.Round, 1, "Invalid Round")
}

func TestDeleteMatchup(t *testing.T) {
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
	request := data.AddMatchupRequest{ID: "id", HomeID: replyHomeTeam.Team.ID, AwayID: replyAwayTeam.Team.ID, Round: 1, Start: start.Format(time.RFC3339)}
	replyMatchup := AddMatchup(replyAddLeague.League.ID, replySeason.Season.Year, request)
	deleteReq := data.DeleteMatchupRequest{SessionID: loginReply.SessionID}
	deleteReply := DeleteMatchup(replyAddLeague.League.ID, replySeason.Season.Year, replyMatchup.Matchup.ID, deleteReq)
	assert.Equal(t, deleteReply.Result.Code, data.SUCCESS, "Should be a success")
	reply := GetMatchups(replyAddLeague.League.ID, replySeason.Season.Year)
	assert.Equal(t, len(reply.Matchups), 0, "Should have zero player")
}

func TestEditMatchup(t *testing.T) {
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
	request := data.AddMatchupRequest{ID: "id", HomeID: replyHomeTeam.Team.ID, AwayID: replyAwayTeam.Team.ID, Round: 1, Start: start.Format(time.RFC3339)}
	replyMatchup := AddMatchup(replyAddLeague.League.ID, replySeason.Season.Year, request)
	start = time.Now()
	editReq := data.EditMatchupRequest{SessionID: loginReply.SessionID, HomeID: replyAwayTeam.Team.ID, AwayID: replyHomeTeam.Team.ID, Round: 2, Start: start.Format(time.RFC3339)}
	editReply := EditMatchup(replyAddLeague.League.ID, replySeason.Season.Year, replyMatchup.Matchup.ID, editReq)
	assert.Equal(t, editReply.Result.Code, data.SUCCESS, "Should be a success")
	reply := GetMatchups(replyAddLeague.League.ID, replySeason.Season.Year)
	assert.Equal(t, len(reply.Matchups), 1, "Should have zero player")
	assert.Equal(t, reply.Matchups[0].League.ID, replyAddLeague.League.ID, "Invalid league")
	assert.Equal(t, reply.Matchups[0].Season.Year, replySeason.Season.Year, "Invalid season")
	assert.Equal(t, reply.Matchups[0].ID, "id", "Invalid id")
	assert.Equal(t, reply.Matchups[0].Home.ID, replyAwayTeam.Team.ID, "Invalid Home")
	assert.Equal(t, reply.Matchups[0].Away.ID, replyHomeTeam.Team.ID, "Invalid away")
	assert.Equal(t, reply.Matchups[0].Round, 2, "Invalid Round")
	assert.Equal(t, reply.Matchups[0].Start.Unix(), start.Unix(), "Invalid start")
}
