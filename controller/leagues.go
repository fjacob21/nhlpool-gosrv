package controller

import (
	"log"

	"nhlpool.com/service/go/nhlpool/store"

	"nhlpool.com/service/go/nhlpool/data"
)

// GetLeagues Process the get players request
func GetLeagues() data.GetLeaguesReply {
	log.Println("Get Leagues")
	var reply data.GetLeaguesReply
	reply.Result.Code = data.SUCCESS
	reply.Leagues, _ = store.GetStore().League().GetLeagues()
	return reply
}

// AddLeague Process the add league request
func AddLeague(request data.AddLeagueRequest) data.AddLeagueReply {
	var reply data.AddLeagueReply
	log.Println("Add League", request)
	league := &data.League{ID: request.ID, Name: request.Name, Description: request.Description, Website: request.Website}
	err := store.GetStore().League().AddLeague(league)
	if err != nil {
		reply.Result.Code = data.EXISTS
		return reply
	}

	reply.Result.Code = data.SUCCESS
	reply.League = *league
	return reply
}
