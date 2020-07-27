package controller

import (
	"log"

	"nhlpool.com/service/go/nhlpool/store"

	"nhlpool.com/service/go/nhlpool/data"
)

// GetWinners Process the get players request
func GetWinners(leagueID string, year int) data.GetWinnersReply {
	log.Println("Get Winners")
	var reply data.GetWinnersReply
	reply.Result.Code = data.SUCCESS
	reply.Winners, _ = store.GetStore().Winner().GetWinners(getLeague(leagueID), getSeason(year, getLeague(leagueID)))
	return reply
}

// AddWinner Process the add team request
func AddWinner(leagueID string, year int, request data.AddWinnerRequest) data.AddWinnerReply {
	var reply data.AddWinnerReply
	log.Println("Add Winner", request)
	league := getLeague(leagueID)
	season := getSeason(year, league)
	player := getPlayer(request.PlayerID)
	team := getTeam(request.Winner, league)
	winner := &data.Winner{
		League: *league,
		Season: *season,
		Player: player,
		Winner: *team,
	}

	err := store.GetStore().Winner().AddWinner(winner)
	if err != nil {
		reply.Result.Code = data.EXISTS
		return reply
	}

	reply.Result.Code = data.SUCCESS
	reply.Winner = *winner
	return reply
}
