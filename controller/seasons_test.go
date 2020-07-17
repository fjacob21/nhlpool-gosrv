package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
	"nhlpool.com/service/go/nhlpool/store"
)

func TestAddSeason(t *testing.T) {
	store.Clean()
	requestAddLeague := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	replyAddLeague := AddLeague(requestAddLeague)
	request := data.AddSeasonRequest{Year: 2000}
	reply := AddSeason(replyAddLeague.League.ID, request)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, reply.Season.Year, 2000, "Invalid year")
	assert.Equal(t, reply.Season.League.ID, replyAddLeague.League.ID, "Invalid league id")
}

func TestGetSeasons(t *testing.T) {
	store.Clean()
	requestAddLeague := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	replyAddLeague := AddLeague(requestAddLeague)
	request := data.AddSeasonRequest{Year: 2000}
	AddSeason(replyAddLeague.League.ID, request)
	reply := GetSeasons(replyAddLeague.League.ID)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, len(reply.Seasons), 1, "Should be only one season")

}
