package controller

import (
	"nhlpool.com/service/go/nhlpool/data"
	"nhlpool.com/service/go/nhlpool/store"
)

func getLeague(ID string) *data.League {
	league, _ := store.GetStore().League().GetLeague(ID)
	return league
}

// GetLeague Process the get player request
func GetLeague(ID string) data.GetLeagueReply {
	var reply data.GetLeagueReply
	league := getLeague(ID)
	if league == nil {
		reply.Result.Code = data.NOTFOUND
		reply.League = data.League{}
		return reply
	}
	reply.Result.Code = data.SUCCESS
	reply.League = *league
	return reply
}

// DeleteLeague Process the delete league request
func DeleteLeague(ID string, request data.DeleteLeagueRequest) data.DeleteLeagueReply {
	var reply data.DeleteLeagueReply
	session := store.GetSessionManager().Get(request.SessionID)
	if session == nil {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	league := getLeague(ID)
	if league == nil {
		reply.Result.Code = data.NOTFOUND
		return reply
	}
	if !session.Player.Admin {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	store.GetStore().League().DeleteLeague(league)
	reply.Result.Code = data.SUCCESS
	return reply
}
