package controller

import (
	"log"

	"nhlpool.com/service/go/nhlpool/store"

	"nhlpool.com/service/go/nhlpool/data"
)

// GetStandings Process the get players request
func GetStandings(leagueID string, year int) data.GetStandingsReply {
	log.Println("Get Standings")
	var reply data.GetStandingsReply
	reply.Result.Code = data.SUCCESS
	reply.Standings, _ = store.GetStore().Standing().GetStandings(getLeague(leagueID), getSeason(year, getLeague(leagueID)))
	return reply
}

// AddStanding Process the add team request
func AddStanding(leagueID string, year int, request data.AddStandingRequest) data.AddStandingReply {
	var reply data.AddStandingReply
	log.Println("Add Standing", request)
	league := getLeague(leagueID)
	season := getSeason(year, league)
	team := getTeam(request.TeamID, league)
	standing := &data.Standing{
		League:       *league,
		Season:       *season,
		Team:         *team,
		Points:       request.Points,
		Win:          request.Win,
		Losses:       request.Losses,
		OT:           request.OT,
		GamesPlayed:  request.GamesPlayed,
		GoalsAgainst: request.GoalsAgainst,
		GoalsScored:  request.GoalsScored,
		Ranks:        request.Ranks,
	}

	err := store.GetStore().Standing().AddStanding(standing)
	if err != nil {
		reply.Result.Code = data.EXISTS
		return reply
	}

	reply.Result.Code = data.SUCCESS
	reply.Standing = *standing
	return reply
}
