package controller

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
	"nhlpool.com/service/go/nhlpool/store"
)

func TestAddMatchup(t *testing.T) {
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
	reply := AddMatchup(replyAddLeague.League.ID, replySeason.Season.Year, request)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, reply.Matchup.League.ID, replyAddLeague.League.ID, "Invalid league")
	assert.Equal(t, reply.Matchup.Season.Year, replySeason.Season.Year, "Invalid season")
	assert.Equal(t, reply.Matchup.ID, "id", "Invalid id")
	assert.Equal(t, reply.Matchup.Home.ID, replyHomeTeam.Team.ID, "Invalid Home")
	assert.Equal(t, reply.Matchup.Away.ID, replyAwayTeam.Team.ID, "Invalid away")
	assert.Equal(t, reply.Matchup.Round, 1, "Invalid Round")
	assert.Equal(t, reply.Matchup.Start.Unix(), start.Unix(), "Invalid start")
}

func TestGetMatchups(t *testing.T) {
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
	assert.NotNil(t, replyMatchup, "Should not be nil")
	reply := GetMatchups(replyAddLeague.League.ID, replySeason.Season.Year)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, len(reply.Matchups), 1, "Should be only one team")
}
