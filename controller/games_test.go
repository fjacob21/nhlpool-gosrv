package controller

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
	"nhlpool.com/service/go/nhlpool/store"
)

func TestAddGame(t *testing.T) {
	store.Clean()
	requestAddLeague := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	replyAddLeague := AddLeague(requestAddLeague)
	requestSeason := data.AddSeasonRequest{Year: 2000}
	replySeason := AddSeason(replyAddLeague.League.ID, requestSeason)
	requestHomeTeam := data.AddTeamRequest{ID: "homeid", Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: data.AddTeamVenue{ID: "id", City: "city", Name: "name", Timezone: "timezone", Address: "address"}}
	replyHomeTeam := AddTeam(replyAddLeague.League.ID, requestHomeTeam)
	requestAwayTeam := data.AddTeamRequest{ID: "awayid", Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: data.AddTeamVenue{ID: "id", City: "city", Name: "name", Timezone: "timezone", Address: "address"}}
	replyAwayTeam := AddTeam(replyAddLeague.League.ID, requestAwayTeam)
	date := time.Now()
	request := data.AddGameRequest{HomeID: replyHomeTeam.Team.ID, AwayID: replyAwayTeam.Team.ID, Date: date.Format(time.RFC3339), Type: data.GameTypeRegular, State: data.GameStateInProgress, HomeGoal: 0, AwayGoal: 1}
	reply := AddGame(replyAddLeague.League.ID, replySeason.Season.Year, request)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, reply.Game.League.ID, replyAddLeague.League.ID, "Invalid league")
	assert.Equal(t, reply.Game.Season.Year, replySeason.Season.Year, "Invalid season")
	assert.Equal(t, reply.Game.Home.ID, replyHomeTeam.Team.ID, "Invalid home")
	assert.Equal(t, reply.Game.Away.ID, replyAwayTeam.Team.ID, "Invalid away")
	assert.Equal(t, reply.Game.Date.Unix(), date.Unix(), "Invalid date")
	assert.Equal(t, reply.Game.Type, data.GameTypeRegular, "Invalid Type")
	assert.Equal(t, reply.Game.State, data.GameStateInProgress, "Invalid State")
	assert.Equal(t, reply.Game.HomeGoal, 0, "Invalid Home goal")
	assert.Equal(t, reply.Game.AwayGoal, 1, "Invalid Away goal")
}
