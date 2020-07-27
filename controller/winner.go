package controller

import (
	"nhlpool.com/service/go/nhlpool/data"
	"nhlpool.com/service/go/nhlpool/store"
)

func getWinner(season *data.Season, player *data.Player) *data.Winner {
	winner, _ := store.GetStore().Winner().GetWinner(player, season.League, season)
	return winner
}

// GetWinner Process the get winner request
func GetWinner(leagueID string, year int, playerID string) data.GetWinnerReply {
	var reply data.GetWinnerReply
	league := getLeague(leagueID)
	season := getSeason(year, league)
	player := getPlayer(playerID)
	winner := getWinner(season, player)
	if winner == nil {
		reply.Result.Code = data.NOTFOUND
		reply.Winner = data.Winner{}
		return reply
	}
	reply.Result.Code = data.SUCCESS
	reply.Winner = *winner
	return reply
}

// EditWinner Process the edit winner request
func EditWinner(leagueID string, year int, playerID string, request data.EditWinnerRequest) data.EditWinnerReply {
	var reply data.EditWinnerReply
	session := store.GetSessionManager().Get(request.SessionID)
	if session == nil {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	league := getLeague(leagueID)
	season := getSeason(year, league)
	player := getPlayer(playerID)
	winner := getWinner(season, player)
	if winner == nil {
		reply.Result.Code = data.NOTFOUND
		return reply
	}
	if !session.Player.Admin {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	team := getTeam(request.Winner, league)
	winner.Winner = *team
	err := store.GetStore().Winner().UpdateWinner(winner)
	if err != nil {
		reply.Result.Code = data.ERROR
		reply.Winner = data.Winner{}
		return reply
	}
	reply.Result.Code = data.SUCCESS
	reply.Winner = *winner
	return reply
}

// DeleteWinner Process the delete winner request
func DeleteWinner(leagueID string, year int, playerID string, request data.DeleteWinnerRequest) data.DeleteWinnerReply {
	var reply data.DeleteWinnerReply
	session := store.GetSessionManager().Get(request.SessionID)
	if session == nil {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	league := getLeague(leagueID)
	season := getSeason(year, league)
	player := getPlayer(playerID)
	winner := getWinner(season, player)
	if winner == nil {
		reply.Result.Code = data.NOTFOUND
		return reply
	}
	if !session.Player.Admin {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	store.GetStore().Winner().DeleteWinner(winner)
	reply.Result.Code = data.SUCCESS
	return reply
}
