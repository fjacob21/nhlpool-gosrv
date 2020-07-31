package controller

import (
	"nhlpool.com/service/go/nhlpool/data"
	"nhlpool.com/service/go/nhlpool/store"
)

func getPrediction(season *data.Season, player *data.Player, matchup *data.Matchup) *data.Prediction {
	prediction, _ := store.GetStore().Prediction().GetPrediction(player, matchup, season.League, season)
	return prediction
}

// GetPrediction Process the get prediction request
func GetPrediction(leagueID string, year int, playerID string, matchupID string) data.GetPredictionReply {
	var reply data.GetPredictionReply
	league := getLeague(leagueID)
	season := getSeason(year, league)
	player := getPlayer(playerID)
	matchup := getMatchup(season, matchupID)
	prediction := getPrediction(season, player, matchup)
	if prediction == nil {
		reply.Result.Code = data.NOTFOUND
		reply.Prediction = data.Prediction{}
		return reply
	}
	reply.Result.Code = data.SUCCESS
	reply.Prediction = *prediction
	return reply
}

// EditPrediction Process the edit prediction request
func EditPrediction(leagueID string, year int, playerID string, matchupID string, request data.EditPredictionRequest) data.EditPredictionReply {
	var reply data.EditPredictionReply
	session := store.GetSessionManager().Get(request.SessionID)
	if session == nil {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	league := getLeague(leagueID)
	season := getSeason(year, league)
	player := getPlayer(playerID)
	matchup := getMatchup(season, matchupID)
	prediction := getPrediction(season, player, matchup)
	if prediction == nil {
		reply.Result.Code = data.NOTFOUND
		return reply
	}
	if !session.Player.Admin {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	team := getTeam(request.Winner, league)
	if team != nil {
		prediction.Winner = *team
	}
	prediction.Games = request.Games
	err := store.GetStore().Prediction().UpdatePrediction(prediction)
	if err != nil {
		reply.Result.Code = data.ERROR
		reply.Prediction = data.Prediction{}
		return reply
	}
	reply.Result.Code = data.SUCCESS
	reply.Prediction = *prediction
	return reply
}

// DeletePrediction Process the delete prediction request
func DeletePrediction(leagueID string, year int, playerID string, matchupID string, request data.DeletePredictionRequest) data.DeletePredictionReply {
	var reply data.DeletePredictionReply
	session := store.GetSessionManager().Get(request.SessionID)
	if session == nil {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	league := getLeague(leagueID)
	season := getSeason(year, league)
	player := getPlayer(playerID)
	matchup := getMatchup(season, matchupID)
	prediction := getPrediction(season, player, matchup)
	if prediction == nil {
		reply.Result.Code = data.NOTFOUND
		return reply
	}
	if !session.Player.Admin {
		reply.Result.Code = data.ACCESSDENIED
		return reply
	}
	store.GetStore().Prediction().DeletePrediction(prediction)
	reply.Result.Code = data.SUCCESS
	return reply
}
