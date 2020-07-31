package store

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestNewSqliteStoreAddGame(t *testing.T) {
	store := NewSqliteStore()
	defer store.Close()
	store.Clean()
	assert.NotNil(t, store, "Should not be nil")
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.League().AddLeague(&league)
	season := &data.Season{Year: 2000, League: &league}
	err = store.Season().AddSeason(season)
	conference := &data.Conference{ID: "id", League: league, Name: "name"}
	err = store.Conference().AddConference(conference)
	homeVenue := &data.Venue{ID: "homeid", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(homeVenue)
	homeTeam := &data.Team{ID: "homeid", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: homeVenue, Conference: conference}
	err = store.Team().AddTeam(homeTeam)
	assert.NoError(t, err, "Should not have error")
	awayVenue := &data.Venue{ID: "awayid", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(awayVenue)
	awayTeam := &data.Team{ID: "awayid", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: awayVenue, Conference: conference}
	err = store.Team().AddTeam(awayTeam)
	assert.NoError(t, err, "Should not have error")
	date := time.Now()
	game := &data.Game{League: league, Season: *season, Home: *homeTeam, Away: *awayTeam, Date: date, Type: data.GameTypeRegular, State: data.GameStateScheduled, HomeGoal: 0, AwayGoal: 1}
	err = store.Game().AddGame(game)
	assert.NoError(t, err, "Should not have error")
	getGame, err := store.Game().GetGame(&league, season, homeTeam, awayTeam, date)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getGame, "Should not be nil")
	assert.Equal(t, getGame.League.ID, league.ID, "Invalid League")
	assert.Equal(t, getGame.Season.Year, season.Year, "Invalid Season")
	assert.Equal(t, getGame.Home.ID, homeTeam.ID, "Invalid Home")
	assert.Equal(t, getGame.Away.ID, awayTeam.ID, "Invalid Away")
	assert.Equal(t, getGame.Date.Unix(), date.Unix(), "Invalid date")
	assert.Equal(t, getGame.Type, data.GameTypeRegular, "Invalid type")
	assert.Equal(t, getGame.State, data.GameStateScheduled, "Invalid state")
	assert.Equal(t, getGame.HomeGoal, 0, "Invalid Home goal")
	assert.Equal(t, getGame.AwayGoal, 1, "Invalid Away goal")
}

func TestNewSqliteStoreUpdateGame(t *testing.T) {
	store := NewSqliteStore()
	defer store.Close()
	store.Clean()
	assert.NotNil(t, store, "Should not be nil")
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.League().AddLeague(&league)
	season := &data.Season{Year: 2000, League: &league}
	err = store.Season().AddSeason(season)
	conference := &data.Conference{ID: "id", League: league, Name: "name"}
	err = store.Conference().AddConference(conference)
	homeVenue := &data.Venue{ID: "homeid", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(homeVenue)
	homeTeam := &data.Team{ID: "homeid", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: homeVenue, Conference: conference}
	err = store.Team().AddTeam(homeTeam)
	assert.NoError(t, err, "Should not have error")
	awayVenue := &data.Venue{ID: "awayid", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(awayVenue)
	awayTeam := &data.Team{ID: "awayid", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: awayVenue, Conference: conference}
	err = store.Team().AddTeam(awayTeam)
	assert.NoError(t, err, "Should not have error")
	date := time.Now()
	game := &data.Game{League: league, Season: *season, Home: *homeTeam, Away: *awayTeam, Date: date, Type: data.GameTypeRegular, State: data.GameStateScheduled, HomeGoal: 0, AwayGoal: 1}
	err = store.Game().AddGame(game)
	assert.NoError(t, err, "Should not have error")
	game.State = data.GameStateInProgress
	game.HomeGoal = 2
	game.AwayGoal = 3
	err = store.Game().UpdateGame(game)
	assert.Nil(t, err, "Should be nil")
	getGame, err := store.Game().GetGame(&league, season, homeTeam, awayTeam, date)
	assert.NoError(t, err, "Should not have error")
	assert.Equal(t, getGame.League.ID, league.ID, "Invalid League")
	assert.Equal(t, getGame.Season.Year, season.Year, "Invalid Season")
	assert.Equal(t, getGame.Home.ID, homeTeam.ID, "Invalid Home")
	assert.Equal(t, getGame.Away.ID, awayTeam.ID, "Invalid Away")
	assert.Equal(t, getGame.Date.Unix(), date.Unix(), "Invalid date")
	assert.Equal(t, getGame.Type, data.GameTypeRegular, "Invalid type")
	assert.Equal(t, getGame.State, data.GameStateInProgress, "Invalid state")
	assert.Equal(t, getGame.HomeGoal, 2, "Invalid Home goal")
	assert.Equal(t, getGame.AwayGoal, 3, "Invalid Away goal")
}

func TestNewSqliteStoreDeleteGame(t *testing.T) {
	store := NewSqliteStore()
	defer store.Close()
	store.Clean()
	assert.NotNil(t, store, "Should not be nil")
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.League().AddLeague(&league)
	season := &data.Season{Year: 2000, League: &league}
	err = store.Season().AddSeason(season)
	conference := &data.Conference{ID: "id", League: league, Name: "name"}
	err = store.Conference().AddConference(conference)
	homeVenue := &data.Venue{ID: "homeid", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(homeVenue)
	homeTeam := &data.Team{ID: "homeid", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: homeVenue, Conference: conference}
	err = store.Team().AddTeam(homeTeam)
	assert.NoError(t, err, "Should not have error")
	awayVenue := &data.Venue{ID: "awayid", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(awayVenue)
	awayTeam := &data.Team{ID: "awayid", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: awayVenue, Conference: conference}
	err = store.Team().AddTeam(awayTeam)
	assert.NoError(t, err, "Should not have error")
	date := time.Now()
	game := &data.Game{League: league, Season: *season, Home: *homeTeam, Away: *awayTeam, Date: date, Type: data.GameTypeRegular, State: data.GameStateScheduled, HomeGoal: 0, AwayGoal: 1}
	err = store.Game().AddGame(game)
	assert.NoError(t, err, "Should not have error")
	err = store.Game().DeleteGame(game)
	assert.NoError(t, err, "Should not have error")
	getGame, err := store.Game().GetGame(&league, season, homeTeam, awayTeam, date)
	assert.Error(t, err, "Should have error")
	assert.Nil(t, getGame, "Should be nil")
	getGames, err := store.Game().GetGames(&league, season, homeTeam, awayTeam)
	assert.NoError(t, err, "Should have error")
	assert.Equal(t, len(getGames), 0, "Should be zero")
}

func TestNewSqliteStoreGetGames(t *testing.T) {
	store := NewSqliteStore()
	defer store.Close()
	store.Clean()
	assert.NotNil(t, store, "Should not be nil")
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.League().AddLeague(&league)
	season := &data.Season{Year: 2000, League: &league}
	err = store.Season().AddSeason(season)
	conference := &data.Conference{ID: "id", League: league, Name: "name"}
	err = store.Conference().AddConference(conference)
	homeVenue := &data.Venue{ID: "homeid", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(homeVenue)
	homeTeam := &data.Team{ID: "homeid", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: homeVenue, Conference: conference}
	err = store.Team().AddTeam(homeTeam)
	assert.NoError(t, err, "Should not have error")
	awayVenue := &data.Venue{ID: "awayid", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(awayVenue)
	awayTeam := &data.Team{ID: "awayid", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: awayVenue, Conference: conference}
	err = store.Team().AddTeam(awayTeam)
	assert.NoError(t, err, "Should not have error")
	date := time.Now()
	game := &data.Game{League: league, Season: *season, Home: *homeTeam, Away: *awayTeam, Date: date, Type: data.GameTypeRegular, State: data.GameStateScheduled, HomeGoal: 0, AwayGoal: 1}
	err = store.Game().AddGame(game)
	assert.NoError(t, err, "Should not have error")

	games, err := store.Game().GetGames(&league, season, homeTeam, awayTeam)
	assert.NotNil(t, games, "Should not be nil")
	assert.Equal(t, len(games), 1, "Should be only one game")
	assert.Equal(t, games[0].League.ID, league.ID, "Should be the good league")
	assert.Equal(t, games[0].Season.Year, season.Year, "Should be the good season")
	assert.Equal(t, games[0].Home.ID, homeTeam.ID, "Should be the good home team")
	assert.Equal(t, games[0].Away.ID, awayTeam.ID, "Should be the good away team")
}
