package controller

import (
	"nhlpool.com/service/go/nhlpool/data"
	"nhlpool.com/service/go/nhlpool/store"
)

func getSeason(year int, league *data.League) *data.Season {
	season, _ := store.GetStore().Season().GetSeason(year, league)
	return season
}

// GetSeason Process the get Season request
func GetSeason(leagueID string, year int) data.GetSeasonReply {
	var reply data.GetSeasonReply
	league := getLeague(leagueID)
	season := getSeason(year, league)
	if season == nil {
		reply.Result.Code = data.NOTFOUND
		reply.Season = data.Season{}
		return reply
	}
	reply.Result.Code = data.SUCCESS
	reply.Season = *season
	return reply
}

// DeleteSeason Process the delete season request
func DeleteSeason(leagueID string, year int, request data.DeleteSeasonRequest) data.DeleteSeasonReply {
	var reply data.DeleteSeasonReply
	session := store.GetSessionManager().Get(request.SessionID)
	if session == nil {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	league := getLeague(leagueID)
	season := getSeason(year, league)
	if season == nil {
		reply.Result.Code = data.NOTFOUND
		return reply
	}
	if !session.Player.Admin {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	store.GetStore().Season().DeleteSeason(season)
	reply.Result.Code = data.SUCCESS
	return reply
}
