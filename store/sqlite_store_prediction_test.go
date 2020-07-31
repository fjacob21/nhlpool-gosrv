package store

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestNewSqliteStoreAddPrediction(t *testing.T) {
	store := NewSqliteStore()
	defer store.Close()
	store.Clean()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	err := store.Player().AddPlayer(player)
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err = store.League().AddLeague(&league)
	season := &data.Season{Year: 2000, League: &league}
	err = store.Season().AddSeason(season)
	homeVenue := &data.Venue{ID: "homeid", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(homeVenue)
	homeTeam := &data.Team{ID: "homeid", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: homeVenue}
	err = store.Team().AddTeam(homeTeam)
	assert.NoError(t, err, "Should not have error")
	awayVenue := &data.Venue{ID: "awayid", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(awayVenue)
	awayTeam := &data.Team{ID: "awayid", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: awayVenue}
	err = store.Team().AddTeam(awayTeam)
	assert.NoError(t, err, "Should not have error")
	start := time.Now()
	matchup := &data.Matchup{League: league, Season: *season, ID: "id", Home: *homeTeam, Away: *awayTeam, Round: 1, Start: start}
	err = store.Matchup().AddMatchup(matchup)
	assert.NoError(t, err, "Should not have error")
	prediction := &data.Prediction{League: league, Season: *season, Player: player, Matchup: matchup, Winner: *homeTeam, Games: 4}
	err = store.Prediction().AddPrediction(prediction)
	assert.NoError(t, err, "Should not have error")
	getPrediction, err := store.Prediction().GetPrediction(player, matchup, &league, season)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getPrediction, "Should not be nil")
	assert.Equal(t, getPrediction.League.ID, league.ID, "Invalid League")
	assert.Equal(t, getPrediction.Season.Year, season.Year, "Invalid Season")
	assert.Equal(t, getPrediction.Player.ID, player.ID, "Invalid Player")
	assert.Equal(t, getPrediction.Matchup.ID, matchup.ID, "Invalid matchup")
	assert.Equal(t, getPrediction.Winner.ID, homeTeam.ID, "Invalid winner")
	assert.Equal(t, getPrediction.Games, 4, "Invalid games")
}

func TestNewSqliteStoreAddPredictionEmpty(t *testing.T) {
	store := NewSqliteStore()
	defer store.Close()
	store.Clean()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	err := store.Player().AddPlayer(player)
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err = store.League().AddLeague(&league)
	season := &data.Season{Year: 2000, League: &league}
	err = store.Season().AddSeason(season)
	matchup := &data.Matchup{League: league, Season: *season, ID: "id", Round: 1}
	err = store.Matchup().AddMatchup(matchup)
	assert.NoError(t, err, "Should not have error")
	prediction := &data.Prediction{League: league, Season: *season, Player: player, Matchup: matchup}
	err = store.Prediction().AddPrediction(prediction)
	assert.NoError(t, err, "Should not have error")
	getPrediction, err := store.Prediction().GetPrediction(player, matchup, &league, season)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getPrediction, "Should not be nil")
	assert.Equal(t, getPrediction.League.ID, league.ID, "Invalid League")
	assert.Equal(t, getPrediction.Season.Year, season.Year, "Invalid Season")
	assert.Equal(t, getPrediction.Player.ID, player.ID, "Invalid Player")
	assert.Equal(t, getPrediction.Matchup.ID, matchup.ID, "Invalid matchup")
	assert.Equal(t, getPrediction.Winner.ID, "", "Invalid winner")
	assert.Equal(t, getPrediction.Games, 0, "Invalid games")
}

func TestNewSqliteStoreUpdatePrediction(t *testing.T) {
	store := NewSqliteStore()
	defer store.Close()
	store.Clean()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	err := store.Player().AddPlayer(player)
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err = store.League().AddLeague(&league)
	season := &data.Season{Year: 2000, League: &league}
	err = store.Season().AddSeason(season)
	homeVenue := &data.Venue{ID: "homeid", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(homeVenue)
	homeTeam := &data.Team{ID: "homeid", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: homeVenue}
	err = store.Team().AddTeam(homeTeam)
	assert.NoError(t, err, "Should not have error")
	awayVenue := &data.Venue{ID: "awayid", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(awayVenue)
	awayTeam := &data.Team{ID: "awayid", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: awayVenue}
	err = store.Team().AddTeam(awayTeam)
	assert.NoError(t, err, "Should not have error")
	start := time.Now()
	matchup := &data.Matchup{League: league, Season: *season, ID: "id", Home: *homeTeam, Away: *awayTeam, Round: 1, Start: start}
	err = store.Matchup().AddMatchup(matchup)
	assert.NoError(t, err, "Should not have error")
	prediction := &data.Prediction{League: league, Season: *season, Player: player, Matchup: matchup, Winner: *homeTeam, Games: 4}
	err = store.Prediction().AddPrediction(prediction)
	assert.NoError(t, err, "Should not have error")
	prediction.Winner = *awayTeam
	prediction.Games = 5
	err = store.Prediction().UpdatePrediction(prediction)
	assert.Nil(t, err, "Should be nil")
	getPrediction, err := store.Prediction().GetPrediction(player, matchup, &league, season)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getPrediction, "Should not be nil")
	assert.Equal(t, getPrediction.League.ID, league.ID, "Invalid League")
	assert.Equal(t, getPrediction.Season.Year, season.Year, "Invalid Season")
	assert.Equal(t, getPrediction.Player.ID, player.ID, "Invalid Player")
	assert.Equal(t, getPrediction.Matchup.ID, matchup.ID, "Invalid matchup")
	assert.Equal(t, getPrediction.Winner.ID, awayTeam.ID, "Invalid winner")
	assert.Equal(t, getPrediction.Games, 5, "Invalid games")
}

func TestNewSqliteStoreDeletePrediction(t *testing.T) {
	store := NewSqliteStore()
	defer store.Close()
	store.Clean()
	assert.NotNil(t, store, "Should not be nil")
	player := data.CreatePlayer("name", "email", false, "password")
	err := store.Player().AddPlayer(player)
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err = store.League().AddLeague(&league)
	season := &data.Season{Year: 2000, League: &league}
	err = store.Season().AddSeason(season)
	homeVenue := &data.Venue{ID: "homeid", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(homeVenue)
	homeTeam := &data.Team{ID: "homeid", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: homeVenue}
	err = store.Team().AddTeam(homeTeam)
	assert.NoError(t, err, "Should not have error")
	awayVenue := &data.Venue{ID: "awayid", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(awayVenue)
	awayTeam := &data.Team{ID: "awayid", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: awayVenue}
	err = store.Team().AddTeam(awayTeam)
	assert.NoError(t, err, "Should not have error")
	start := time.Now()
	matchup := &data.Matchup{League: league, Season: *season, ID: "id", Home: *homeTeam, Away: *awayTeam, Round: 1, Start: start}
	err = store.Matchup().AddMatchup(matchup)
	assert.NoError(t, err, "Should not have error")
	prediction := &data.Prediction{League: league, Season: *season, Player: player, Matchup: matchup, Winner: *homeTeam, Games: 4}
	err = store.Prediction().AddPrediction(prediction)
	assert.NoError(t, err, "Should not have error")
	getPrediction, err := store.Prediction().GetPrediction(player, matchup, &league, season)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getPrediction, "Should not be nil")
	err = store.Prediction().DeletePrediction(getPrediction)
	assert.NoError(t, err, "Should not have error")
	getPrediction, err = store.Prediction().GetPrediction(player, matchup, &league, season)
	assert.Error(t, err, "Should have error")
	assert.Nil(t, getPrediction, "Should be nil")
	getPredictions, err := store.Prediction().GetPredictions(&league, season)
	assert.NoError(t, err, "Should have error")
	assert.Equal(t, len(getPredictions), 0, "Should be zero")
}

func TestNewSqliteStoreGetPredictions(t *testing.T) {
	fmt.Printf("Start\n")
	store := NewSqliteStore()
	defer store.Close()
	store.Clean()
	assert.NotNil(t, store, "Should not be nil")

	player := data.CreatePlayer("name", "email", false, "password")
	err := store.Player().AddPlayer(player)
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err = store.League().AddLeague(&league)
	season := &data.Season{Year: 2000, League: &league}
	err = store.Season().AddSeason(season)
	homeVenue := &data.Venue{ID: "homeid", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(homeVenue)
	homeTeam := &data.Team{ID: "homeid", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: homeVenue}
	err = store.Team().AddTeam(homeTeam)
	assert.NoError(t, err, "Should not have error")
	awayVenue := &data.Venue{ID: "awayid", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(awayVenue)
	awayTeam := &data.Team{ID: "awayid", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: awayVenue}
	err = store.Team().AddTeam(awayTeam)
	assert.NoError(t, err, "Should not have error")
	start := time.Now()
	matchup := &data.Matchup{League: league, Season: *season, ID: "id", Home: *homeTeam, Away: *awayTeam, Round: 1, Start: start}
	err = store.Matchup().AddMatchup(matchup)
	assert.NoError(t, err, "Should not have error")
	prediction := &data.Prediction{League: league, Season: *season, Player: player, Matchup: matchup, Winner: *homeTeam, Games: 4}
	err = store.Prediction().AddPrediction(prediction)
	assert.NoError(t, err, "Should not have error")

	predictions, _ := store.Prediction().GetPredictions(&league, season)
	assert.NotNil(t, predictions, "Should not be nil")
	assert.Equal(t, len(predictions), 1, "Should be only one team")
	assert.Equal(t, predictions[0].League.ID, league.ID, "Should be the good league")
	assert.Equal(t, predictions[0].Season.Year, season.Year, "Should be the good season")
	assert.Equal(t, predictions[0].Player.ID, player.ID, "Should be the good player")
	assert.Equal(t, predictions[0].Matchup.ID, matchup.ID, "Should be the good matchup")
	assert.Equal(t, predictions[0].Winner.ID, homeTeam.ID, "Should be the good winner")
	assert.Equal(t, predictions[0].Games, 4, "Should be the good matchup")
}

func TestNewSqliteStoreGetPredictionsEmpty(t *testing.T) {
	fmt.Printf("Start\n")
	store := NewSqliteStore()
	defer store.Close()
	store.Clean()
	assert.NotNil(t, store, "Should not be nil")

	player := data.CreatePlayer("name", "email", false, "password")
	err := store.Player().AddPlayer(player)
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err = store.League().AddLeague(&league)
	season := &data.Season{Year: 2000, League: &league}
	err = store.Season().AddSeason(season)
	matchup := &data.Matchup{League: league, Season: *season, ID: "id", Round: 1}
	err = store.Matchup().AddMatchup(matchup)
	assert.NoError(t, err, "Should not have error")
	prediction := &data.Prediction{League: league, Season: *season, Player: player, Matchup: matchup}
	err = store.Prediction().AddPrediction(prediction)
	assert.NoError(t, err, "Should not have error")

	predictions, _ := store.Prediction().GetPredictions(&league, season)
	assert.NotNil(t, predictions, "Should not be nil")
	assert.Equal(t, len(predictions), 1, "Should be only one team")
	assert.Equal(t, predictions[0].League.ID, league.ID, "Should be the good league")
	assert.Equal(t, predictions[0].Season.Year, season.Year, "Should be the good season")
	assert.Equal(t, predictions[0].Player.ID, player.ID, "Should be the good player")
	assert.Equal(t, predictions[0].Matchup.ID, matchup.ID, "Should be the good matchup")
	assert.Equal(t, predictions[0].Winner.ID, "", "Should be the good winner")
	assert.Equal(t, predictions[0].Games, 0, "Should be the good matchup")
}
