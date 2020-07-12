package controller

import (
	"testing"

	"nhlpool.com/service/go/nhlpool/store"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestGetLeague(t *testing.T) {
	store.Clean()
	request := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	reply := AddLeague(request)
	assert.NotNil(t, reply, "Should not be nil")
	leagueReply := GetLeague("id")
	assert.NotNil(t, leagueReply, "Should not be nil")
	assert.Equal(t, leagueReply.League.ID, reply.League.ID, "Invalid ID")
	assert.Equal(t, leagueReply.League.Name, "name", "Invalid name")
	assert.Equal(t, leagueReply.League.Description, "description", "Invalid description")
	assert.Equal(t, leagueReply.League.Website, "website", "Invalid website")
}
