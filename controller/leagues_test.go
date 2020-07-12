package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
	"nhlpool.com/service/go/nhlpool/store"
)

func TestAddLeague(t *testing.T) {
	store.Clean()
	request := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	reply := AddLeague(request)
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, reply.League.ID, "id", "Invalid id")
	assert.Equal(t, reply.League.Name, "name", "Invalid name")
	assert.Equal(t, reply.League.Description, "description", "Invalid description")
	assert.Equal(t, reply.League.Website, "website", "Invalid website")
}

func TestGetLeagues(t *testing.T) {
	store.Clean()
	request := data.AddLeagueRequest{ID: "id", Name: "name", Description: "description", Website: "website"}
	AddLeague(request)
	reply := GetLeagues()
	assert.NotNil(t, reply, "Should not be nil")
	assert.Equal(t, len(reply.Leagues), 1, "Should be only one player")

}
