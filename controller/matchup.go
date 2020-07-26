package controller

import (
	"time"

	"nhlpool.com/service/go/nhlpool/data"
	"nhlpool.com/service/go/nhlpool/store"
)

func getMatchup(season *data.Season, id string) *data.Matchup {
	matchup, _ := store.GetStore().Matchup().GetMatchup(season.League, season, id)
	return matchup
}

// GetMatchup Process the get matchup request
func GetMatchup(leagueID string, year int, ID string) data.GetMatchupReply {
	var reply data.GetMatchupReply
	league := getLeague(leagueID)
	season := getSeason(year, league)
	matchup := getMatchup(season, ID)
	if matchup == nil {
		reply.Result.Code = data.NOTFOUND
		reply.Matchup = data.Matchup{}
		return reply
	}
	reply.Result.Code = data.SUCCESS
	reply.Matchup = *matchup
	return reply
}

// EditMatchup Process the edit matchup request
func EditMatchup(leagueID string, year int, ID string, request data.EditMatchupRequest) data.EditMatchupReply {
	var reply data.EditMatchupReply
	session := store.GetSessionManager().Get(request.SessionID)
	if session == nil {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	league := getLeague(leagueID)
	season := getSeason(year, league)
	matchup := getMatchup(season, ID)
	if matchup == nil {
		reply.Result.Code = data.NOTFOUND
		return reply
	}
	if !session.Player.Admin {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	home := getTeam(request.HomeID, league)
	away := getTeam(request.AwayID, league)
	start, _ := time.Parse(time.RFC3339, request.Start)
	matchup.Home = *home
	matchup.Away = *away
	matchup.Start = start
	matchup.Round = request.Round
	err := store.GetStore().Matchup().UpdateMatchup(matchup)
	if err != nil {
		reply.Result.Code = data.ERROR
		reply.Matchup = data.Matchup{}
		return reply
	}
	reply.Result.Code = data.SUCCESS
	reply.Matchup = *matchup
	return reply
}

// DeleteMatchup Process the delete matchup request
func DeleteMatchup(leagueID string, year int, ID string, request data.DeleteMatchupRequest) data.DeleteMatchupReply {
	var reply data.DeleteMatchupReply
	session := store.GetSessionManager().Get(request.SessionID)
	if session == nil {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	league := getLeague(leagueID)
	season := getSeason(year, league)
	matchup := getMatchup(season, ID)
	if matchup == nil {
		reply.Result.Code = data.NOTFOUND
		return reply
	}
	if !session.Player.Admin {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	store.GetStore().Matchup().DeleteMatchup(matchup)
	reply.Result.Code = data.SUCCESS
	return reply
}
