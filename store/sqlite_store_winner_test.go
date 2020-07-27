package store

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestNewSqliteStoreAddWinner(t *testing.T) {
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
	venue := &data.Venue{ID: "id", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(venue)
	team := &data.Team{ID: "id", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: venue}
	err = store.Team().AddTeam(team)
	assert.NoError(t, err, "Should not have error")
	winner := &data.Winner{League: league, Season: *season, Player: player, Winner: *team}
	err = store.Winner().AddWinner(winner)
	assert.NoError(t, err, "Should not have error")
	getWinner, err := store.Winner().GetWinner(player, &league, season)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getWinner, "Should not be nil")
	assert.Equal(t, getWinner.League.ID, league.ID, "Invalid League")
	assert.Equal(t, getWinner.Season.Year, season.Year, "Invalid Season")
	assert.Equal(t, getWinner.Player.ID, player.ID, "Invalid Player")
	assert.Equal(t, getWinner.Winner.ID, team.ID, "Invalid Winner")
}

func TestNewSqliteStoreUpdateWinner(t *testing.T) {
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
	venue := &data.Venue{ID: "id", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(venue)
	team := &data.Team{ID: "id", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: venue}
	err = store.Team().AddTeam(team)
	team2 := &data.Team{ID: "id2", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: venue}
	err = store.Team().AddTeam(team2)
	assert.NoError(t, err, "Should not have error")
	winner := &data.Winner{League: league, Season: *season, Player: player, Winner: *team}
	err = store.Winner().AddWinner(winner)
	assert.NoError(t, err, "Should not have error")
	winner.Winner = *team2
	err = store.Winner().UpdateWinner(winner)
	assert.Nil(t, err, "Should be nil")
	getWinner, err := store.Winner().GetWinner(player, &league, season)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getWinner, "Should not be nil")
	assert.Equal(t, getWinner.League.ID, league.ID, "Invalid League")
	assert.Equal(t, getWinner.Season.Year, season.Year, "Invalid Season")
	assert.Equal(t, getWinner.Player.ID, player.ID, "Invalid Player")
	assert.Equal(t, getWinner.Winner.ID, team2.ID, "Invalid Winner")
}

func TestNewSqliteStoreDeleteWinner(t *testing.T) {
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
	venue := &data.Venue{ID: "id", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(venue)
	team := &data.Team{ID: "id", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: venue}
	err = store.Team().AddTeam(team)
	assert.NoError(t, err, "Should not have error")
	winner := &data.Winner{League: league, Season: *season, Player: player, Winner: *team}
	err = store.Winner().AddWinner(winner)
	assert.NoError(t, err, "Should not have error")
	getWinner, err := store.Winner().GetWinner(player, &league, season)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getWinner, "Should not be nil")
	err = store.Winner().DeleteWinner(winner)
	assert.NoError(t, err, "Should not have error")
	getWinner, err = store.Winner().GetWinner(player, &league, season)
	assert.Error(t, err, "Should have error")
	assert.Nil(t, getWinner, "Should be nil")
	getWinners, err := store.Winner().GetWinners(&league, season)
	assert.NoError(t, err, "Should have error")
	assert.Equal(t, len(getWinners), 0, "Should be zero")
}

func TestNewSqliteStoreGetWinners(t *testing.T) {
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
	venue := &data.Venue{ID: "id", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(venue)
	team := &data.Team{ID: "id", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: venue}
	err = store.Team().AddTeam(team)
	assert.NoError(t, err, "Should not have error")
	winner := &data.Winner{League: league, Season: *season, Player: player, Winner: *team}
	err = store.Winner().AddWinner(winner)
	assert.NoError(t, err, "Should not have error")

	winners, _ := store.Winner().GetWinners(&league, season)
	assert.NotNil(t, winners, "Should not be nil")
	assert.Equal(t, len(winners), 1, "Should be only one team")
	assert.Equal(t, winners[0].League.ID, league.ID, "Should be the good league")
	assert.Equal(t, winners[0].Season.Year, season.Year, "Should be the good season")
	assert.Equal(t, winners[0].Player.ID, player.ID, "Should be the good player")
	assert.Equal(t, winners[0].Winner.ID, team.ID, "Should be the good winner")
}
