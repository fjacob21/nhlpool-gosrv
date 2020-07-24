package controller

import (
	"nhlpool.com/service/go/nhlpool/data"
	"nhlpool.com/service/go/nhlpool/store"
)

func getStanding(team *data.Team, season *data.Season) *data.Standing {
	standing, _ := store.GetStore().Standing().GetStanding(team, season.League, season)
	return standing
}

// GetStanding Process the get Standing request
func GetStanding(leagueID string, year int, teamID string) data.GetStandingReply {
	var reply data.GetStandingReply
	league := getLeague(leagueID)
	season := getSeason(year, league)
	team := getTeam(teamID, league)
	standing := getStanding(team, season)
	if standing == nil {
		reply.Result.Code = data.NOTFOUND
		reply.Standing = data.Standing{}
		return reply
	}
	reply.Result.Code = data.SUCCESS
	reply.Standing = *standing
	return reply
}

// DeleteStanding Process the delete league request
func DeleteStanding(leagueID string, year int, teamID string, request data.DeleteStandingRequest) data.DeleteStandingReply {
	var reply data.DeleteStandingReply
	session := store.GetSessionManager().Get(request.SessionID)
	if session == nil {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	league := getLeague(leagueID)
	season := getSeason(year, league)
	team := getTeam(teamID, league)
	standing := getStanding(team, season)
	if standing == nil {
		reply.Result.Code = data.NOTFOUND
		return reply
	}
	if !session.Player.Admin {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	store.GetStore().Standing().DeleteStanding(standing)
	reply.Result.Code = data.SUCCESS
	return reply
}
