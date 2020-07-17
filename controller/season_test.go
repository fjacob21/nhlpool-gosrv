package controller

import (
	"testing"

	"nhlpool.com/service/go/nhlpool/store"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestGetSeason(t *testing.T) {
	store.Clean()
	requestAddLeague := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	replyAddLeague := AddLeague(requestAddLeague)
	request := data.AddSeasonRequest{Year: 2000}
	reply := AddSeason(replyAddLeague.League.ID, request)
	assert.NotNil(t, reply, "Should not be nil")
	seasonReply := GetSeason(replyAddLeague.League.ID, 2000)
	assert.NotNil(t, seasonReply, "Should not be nil")
	assert.Equal(t, seasonReply.Season.Year, reply.Season.Year, "Invalid year")
	assert.Equal(t, seasonReply.Season.League.ID, replyAddLeague.League.ID, "Invalid league")
}
