package controller

import (
	"log"
	"time"

	"nhlpool.com/service/go/nhlpool/store"

	"nhlpool.com/service/go/nhlpool/data"
)

// GetMatchups Process the get matchup request
func GetMatchups(leagueID string, year int) data.GetMatchupsReply {
	log.Println("Get Matchups")
	var reply data.GetMatchupsReply
	reply.Result.Code = data.SUCCESS
	reply.Matchups, _ = store.GetStore().Matchup().GetMatchups(getLeague(leagueID), getSeason(year, getLeague(leagueID)))
	return reply
}

// AddMatchup Process the add matchup request
func AddMatchup(leagueID string, year int, request data.AddMatchupRequest) data.AddMatchupReply {
	var reply data.AddMatchupReply
	log.Println("Add Standing", request)
	league := getLeague(leagueID)
	season := getSeason(year, league)
	home := getTeam(request.HomeID, league)
	away := getTeam(request.AwayID, league)
	start, _ := time.Parse(time.RFC3339, request.Start)
	matchup := &data.Matchup{
		League: *league,
		Season: *season,
		ID:     request.ID,
		Home:   *home,
		Away:   *away,
		Round:  request.Round,
		Start:  start,
	}

	err := store.GetStore().Matchup().AddMatchup(matchup)
	if err != nil {
		reply.Result.Code = data.EXISTS
		return reply
	}

	reply.Result.Code = data.SUCCESS
	reply.Matchup = *matchup
	return reply
}
