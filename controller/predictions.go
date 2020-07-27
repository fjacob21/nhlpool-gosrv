package controller

import (
	"log"

	"nhlpool.com/service/go/nhlpool/store"

	"nhlpool.com/service/go/nhlpool/data"
)

// GetPredictions Process the get prediction request
func GetPredictions(leagueID string, year int) data.GetPredictionsReply {
	log.Println("Get Predictions")
	var reply data.GetPredictionsReply
	reply.Result.Code = data.SUCCESS
	reply.Predictions, _ = store.GetStore().Prediction().GetPredictions(getLeague(leagueID), getSeason(year, getLeague(leagueID)))
	return reply
}

// AddPrediction Process the add prediction request
func AddPrediction(leagueID string, year int, request data.AddPredictionRequest) data.AddPredictionReply {
	var reply data.AddPredictionReply
	log.Println("Add Prediction", request)
	league := getLeague(leagueID)
	season := getSeason(year, league)
	player := getPlayer(request.PlayerID)
	matchup := getMatchup(season, request.MatchupID)
	team := getTeam(request.Winner, league)
	prediction := &data.Prediction{
		League:  *league,
		Season:  *season,
		Player:  player,
		Matchup: matchup,
		Winner:  *team,
		Games:   request.Games,
	}

	err := store.GetStore().Prediction().AddPrediction(prediction)
	if err != nil {
		reply.Result.Code = data.EXISTS
		return reply
	}

	reply.Result.Code = data.SUCCESS
	reply.Prediction = *prediction
	return reply
}
