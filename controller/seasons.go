package controller

import (
	"log"

	"nhlpool.com/service/go/nhlpool/store"

	"nhlpool.com/service/go/nhlpool/data"
)

// GetSeasons Process the get seasons request
func GetSeasons(leagueID string) data.GetSeasonsReply {
	log.Println("Get Seasons")
	var reply data.GetSeasonsReply
	reply.Result.Code = data.SUCCESS
	reply.Seasons, _ = store.GetStore().Season().GetSeasons(getLeague(leagueID))
	return reply
}

// AddSeason Process the add season request
func AddSeason(leagueID string, request data.AddSeasonRequest) data.AddSeasonReply {
	var reply data.AddSeasonReply
	log.Println("Add Season", request)
	league := getLeague(leagueID)
	season := &data.Season{Year: request.Year, League: *league}

	err := store.GetStore().Season().AddSeason(season)
	if err != nil {
		reply.Result.Code = data.EXISTS
		return reply
	}

	reply.Result.Code = data.SUCCESS
	reply.Season = *season
	return reply
}
