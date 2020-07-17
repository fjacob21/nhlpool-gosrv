package controller

import (
	"nhlpool.com/service/go/nhlpool/data"
	"nhlpool.com/service/go/nhlpool/store"
)

func getTeam(ID string, league *data.League) *data.Team {
	team, _ := store.GetStore().Team().GetTeam(ID, league)
	return team
}

// GetTeam Process the get Team request
func GetTeam(leagueID string, ID string) data.GetTeamReply {
	var reply data.GetTeamReply
	league := getLeague(leagueID)
	team := getTeam(ID, league)
	if team == nil {
		reply.Result.Code = data.NOTFOUND
		reply.Team = data.Team{}
		return reply
	}
	reply.Result.Code = data.SUCCESS
	reply.Team = *team
	return reply
}

// DeleteTeam Process the delete league request
func DeleteTeam(leagueID string, ID string, request data.DeleteTeamRequest) data.DeleteTeamReply {
	var reply data.DeleteTeamReply
	session := store.GetSessionManager().Get(request.SessionID)
	if session == nil {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	league := getLeague(leagueID)
	team := getTeam(ID, league)
	if team == nil {
		reply.Result.Code = data.NOTFOUND
		return reply
	}
	if !session.Player.Admin {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	store.GetStore().Team().DeleteTeam(team)
	reply.Result.Code = data.SUCCESS
	return reply
}
