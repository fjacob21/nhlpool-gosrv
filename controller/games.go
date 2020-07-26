package controller

import (
	"fmt"
	"log"
	"time"

	"nhlpool.com/service/go/nhlpool/store"

	"nhlpool.com/service/go/nhlpool/data"
)

// GetGames Process the get games request
func GetGames(leagueID string, year int) data.GetGamesReply {
	log.Println("Get Games")
	var reply data.GetGamesReply
	reply.Result.Code = data.SUCCESS
	reply.Games, _ = store.GetStore().Game().GetAllGames(getLeague(leagueID), getSeason(year, getLeague(leagueID)))
	return reply
}

// AddGame Process the add game request
func AddGame(leagueID string, year int, request data.AddGameRequest) data.AddGameReply {
	var reply data.AddGameReply
	log.Println("Add Game", request)
	league := getLeague(leagueID)
	season := getSeason(year, league)
	home := getTeam(request.HomeID, league)
	away := getTeam(request.AwayID, league)
	date, err := time.Parse(time.RFC3339, request.Date)
	if err != nil {
		fmt.Printf("GetGame Invalid time Err: %v\n", err)
	}
	game := &data.Game{
		League:   *league,
		Season:   *season,
		Home:     *home,
		Away:     *away,
		Date:     date,
		Type:     request.Type,
		State:    request.State,
		HomeGoal: request.HomeGoal,
		AwayGoal: request.AwayGoal,
	}

	err = store.GetStore().Game().AddGame(game)
	if err != nil {
		reply.Result.Code = data.EXISTS
		return reply
	}

	reply.Result.Code = data.SUCCESS
	reply.Game = *game
	return reply
}
