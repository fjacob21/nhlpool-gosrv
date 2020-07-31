package store

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestNewSqliteStoreAddstanding(t *testing.T) {
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
	venue := &data.Venue{ID: "id", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(venue)
	team := &data.Team{ID: "id", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: venue, Conference: conference}
	err = store.Team().AddTeam(team)
	assert.NoError(t, err, "Should not have error")
	standing := &data.Standing{League: league, Season: *season, Team: *team, Points: 0, Win: 1, Losses: 2, OT: 3, GamesPlayed: 4, GoalsAgainst: 5, GoalsScored: 6, Ranks: 7}
	err = store.Standing().AddStanding(standing)
	assert.NoError(t, err, "Should not have error")
	getStanding, err := store.Standing().GetStanding(team, &league, season)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getStanding, "Should not be nil")
	assert.Equal(t, getStanding.League.ID, league.ID, "Invalid League")
	assert.Equal(t, getStanding.Season.Year, season.Year, "Invalid Season")
	assert.Equal(t, getStanding.Team.ID, team.ID, "Invalid Team")
	assert.Equal(t, getStanding.Points, 0, "Invalid Point")
	assert.Equal(t, getStanding.Win, 1, "Invalid Win")
	assert.Equal(t, getStanding.Losses, 2, "Invalid Losses")
	assert.Equal(t, getStanding.OT, 3, "Invalid OT")
	assert.Equal(t, getStanding.GamesPlayed, 4, "Invalid GamesPlayed")
	assert.Equal(t, getStanding.GoalsAgainst, 5, "Invalid GoalsAgainst")
	assert.Equal(t, getStanding.GoalsScored, 6, "Invalid GoalsScored")
	assert.Equal(t, getStanding.Ranks, 7, "Invalid Ranks")
}

func TestNewSqliteStoreUpdatestanding(t *testing.T) {
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
	venue := &data.Venue{ID: "id", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(venue)
	team := &data.Team{ID: "id", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: venue, Conference: conference}
	err = store.Team().AddTeam(team)
	assert.NoError(t, err, "Should not have error")
	standing := &data.Standing{League: league, Season: *season, Team: *team, Points: 0, Win: 1, Losses: 2, OT: 3, GamesPlayed: 4, GoalsAgainst: 5, GoalsScored: 6, Ranks: 7}
	err = store.Standing().AddStanding(standing)
	assert.NoError(t, err, "Should not have error")
	standing.Points = 8
	standing.Win = 9
	standing.Losses = 10
	standing.OT = 11
	standing.GamesPlayed = 12
	standing.GoalsAgainst = 13
	standing.GoalsScored = 14
	standing.Ranks = 15
	err = store.Standing().UpdateStanding(standing)
	assert.Nil(t, err, "Should be nil")
	getStanding, err := store.Standing().GetStanding(team, &league, season)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getStanding, "Should not be nil")
	assert.Equal(t, getStanding.League.ID, league.ID, "Invalid League")
	assert.Equal(t, getStanding.Season.Year, season.Year, "Invalid Season")
	assert.Equal(t, getStanding.Team.ID, team.ID, "Invalid Team")
	assert.Equal(t, getStanding.Points, 8, "Invalid Point")
	assert.Equal(t, getStanding.Win, 9, "Invalid Win")
	assert.Equal(t, getStanding.Losses, 10, "Invalid Losses")
	assert.Equal(t, getStanding.OT, 11, "Invalid OT")
	assert.Equal(t, getStanding.GamesPlayed, 12, "Invalid GamesPlayed")
	assert.Equal(t, getStanding.GoalsAgainst, 13, "Invalid GoalsAgainst")
	assert.Equal(t, getStanding.GoalsScored, 14, "Invalid GoalsScored")
	assert.Equal(t, getStanding.Ranks, 15, "Invalid Ranks")
}

func TestNewSqliteStoreDeletestanding(t *testing.T) {
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
	venue := &data.Venue{ID: "id", League: league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	err = store.Venue().AddVenue(venue)
	team := &data.Team{ID: "id", League: league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: venue, Conference: conference}
	err = store.Team().AddTeam(team)
	assert.NoError(t, err, "Should not have error")
	standing := &data.Standing{League: league, Season: *season, Team: *team, Points: 0, Win: 1, Losses: 2, OT: 3, GamesPlayed: 4, GoalsAgainst: 5, GoalsScored: 6, Ranks: 7}
	err = store.Standing().AddStanding(standing)
	assert.NoError(t, err, "Should not have error")
	getStanding, err := store.Standing().GetStanding(team, &league, season)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getStanding, "Should not be nil")
	err = store.Standing().DeleteStanding(standing)
	assert.NoError(t, err, "Should not have error")
	getStanding, err = store.Standing().GetStanding(team, &league, season)
	assert.Error(t, err, "Should have error")
	assert.Nil(t, getStanding, "Should be nil")
	getStandings, err := store.Standing().GetStandings(&league, season)
	assert.NoError(t, err, "Should have error")
	assert.Equal(t, len(getStandings), 0, "Should be zero")
}

func TestNewSqliteStoreGetstandings(t *testing.T) {
	fmt.Printf("Start\n")
	store := NewSqliteStore()
	defer store.Close()
	store.Clean()
	assert.NotNil(t, store, "Should not be nil")

	league := &data.League{ID: "leagueid", Name: "name", Description: "description", Website: "website"}
	store.League().AddLeague(league)
	season := &data.Season{Year: 2000, League: league}
	store.Season().AddSeason(season)
	conference := &data.Conference{ID: "id", League: *league, Name: "name"}
	err := store.Conference().AddConference(conference)
	venue := &data.Venue{ID: "venueid", League: *league, City: "city", Name: "name", Timezone: "timezone", Address: "address"}
	store.Venue().AddVenue(venue)
	team := data.Team{ID: "teamid", League: *league, Abbreviation: "abbreviation", Name: "name", Fullname: "fullname", City: "city", Active: true, CreationYear: "creationyear", Website: "website", Venue: venue, Conference: conference}
	store.Team().AddTeam(&team)
	standing := &data.Standing{League: *league, Season: *season, Team: team, Points: 0, Win: 1, Losses: 2, OT: 3, GamesPlayed: 4, GoalsAgainst: 5, GoalsScored: 6, Ranks: 7}
	err = store.Standing().AddStanding(standing)
	assert.NoError(t, err, "Should not have error")

	standings, _ := store.Standing().GetStandings(league, season)
	assert.NotNil(t, standings, "Should not be nil")
	assert.Equal(t, len(standings), 1, "Should be only one team")
	assert.Equal(t, standings[0].League.ID, league.ID, "Should be the good league")
	assert.Equal(t, standings[0].Season.Year, season.Year, "Should be the good season")
	assert.Equal(t, standings[0].Team.ID, team.ID, "Should be the good team")
}
