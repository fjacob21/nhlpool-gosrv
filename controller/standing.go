package controller

import (
	"fmt"

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

// EditStanding Process the edit standing request
func EditStanding(leagueID string, year int, teamID string, request data.EditStandingRequest) data.EditStandingReply {
	var reply data.EditStandingReply
	session := store.GetSessionManager().Get(request.SessionID)
	if session == nil {
		fmt.Printf("Session Not found\n")
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	league := getLeague(leagueID)
	season := getSeason(year, league)
	team := getTeam(teamID, league)
	standing := getStanding(team, season)
	if standing == nil {
		fmt.Printf("Not found\n")
		reply.Result.Code = data.NOTFOUND
		return reply
	}
	if !session.Player.Admin {
		fmt.Printf("Not admin\n")
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	standing.Points = request.Points
	standing.Win = request.Win
	standing.Losses = request.Losses
	standing.OT = request.OT
	standing.GamesPlayed = request.GamesPlayed
	standing.GoalsAgainst = request.GoalsAgainst
	standing.GoalsScored = request.GoalsScored
	standing.Ranks = request.Ranks
	err := store.GetStore().Standing().UpdateStanding(standing)
	if err != nil {
		reply.Result.Code = data.ERROR
		reply.Standing = data.Standing{}
		fmt.Printf("Update err %v\n", err)
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
