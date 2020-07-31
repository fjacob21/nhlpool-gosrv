package store

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestNewSqliteStoreAddMatchup(t *testing.T) {
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
	start := time.Now()
	matchup := &data.Matchup{League: league, Season: *season, ID: "id", Home: *homeTeam, Away: *awayTeam, Round: 1, Start: start}
	err = store.Matchup().AddMatchup(matchup)
	assert.NoError(t, err, "Should not have error")
	getMatchup, err := store.Matchup().GetMatchup(&league, season, matchup.ID)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getMatchup, "Should not be nil")
	assert.Equal(t, getMatchup.League.ID, league.ID, "Invalid League")
	assert.Equal(t, getMatchup.Season.Year, season.Year, "Invalid Season")
	assert.Equal(t, getMatchup.Home.ID, homeTeam.ID, "Invalid Home")
	assert.Equal(t, getMatchup.Away.ID, awayTeam.ID, "Invalid Away")
	assert.Equal(t, getMatchup.Round, 1, "Invalid Round")
	assert.Equal(t, getMatchup.Start.Unix(), start.Unix(), "Invalid Start")
}

func TestNewSqliteStoreAddMatchupEmpty(t *testing.T) {
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
	matchup := &data.Matchup{League: league, Season: *season, ID: "id", Round: 1}
	err = store.Matchup().AddMatchup(matchup)
	assert.NoError(t, err, "Should not have error")
	getMatchup, err := store.Matchup().GetMatchup(&league, season, matchup.ID)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getMatchup, "Should not be nil")
	assert.Equal(t, getMatchup.League.ID, league.ID, "Invalid League")
	assert.Equal(t, getMatchup.Season.Year, season.Year, "Invalid Season")
	assert.Equal(t, getMatchup.Home.ID, "", "Invalid Home")
	assert.Equal(t, getMatchup.Away.ID, "", "Invalid Away")
	assert.Equal(t, getMatchup.Round, 1, "Invalid Round")
}

func TestNewSqliteStoreUpdateMatchup(t *testing.T) {
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
	start := time.Now()
	matchup := &data.Matchup{League: league, Season: *season, ID: "id", Round: 1, Start: start}
	err = store.Matchup().AddMatchup(matchup)
	assert.NoError(t, err, "Should not have error")
	matchup.Home = *homeTeam
	matchup.Away = *awayTeam
	matchup.Round = 2
	start = time.Now()
	matchup.Start = start
	err = store.Matchup().UpdateMatchup(matchup)
	assert.Nil(t, err, "Should be nil")
	getMatchup, err := store.Matchup().GetMatchup(&league, season, "id")
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getMatchup, "Should not be nil")
	assert.Equal(t, getMatchup.League.ID, league.ID, "Invalid League")
	assert.Equal(t, getMatchup.Season.Year, season.Year, "Invalid Season")
	assert.Equal(t, getMatchup.Home.ID, homeTeam.ID, "Invalid Home")
	assert.Equal(t, getMatchup.Away.ID, awayTeam.ID, "Invalid Away")
	assert.Equal(t, getMatchup.Round, 2, "Invalid Round")
	assert.Equal(t, getMatchup.Start.Unix(), start.Unix(), "Invalid Start")
}

func TestNewSqliteStoreUpdateMatchupEmpty(t *testing.T) {
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
	start := time.Now()
	matchup := &data.Matchup{League: league, Season: *season, ID: "id", Round: 1, Start: start}
	err = store.Matchup().AddMatchup(matchup)
	assert.NoError(t, err, "Should not have error")
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
	matchup.Home = *homeTeam
	matchup.Away = *awayTeam
	matchup.Round = 2
	start = time.Now()
	matchup.Start = start
	err = store.Matchup().UpdateMatchup(matchup)
	assert.Nil(t, err, "Should be nil")
	getMatchup, err := store.Matchup().GetMatchup(&league, season, "id")
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getMatchup, "Should not be nil")
	assert.Equal(t, getMatchup.League.ID, league.ID, "Invalid League")
	assert.Equal(t, getMatchup.Season.Year, season.Year, "Invalid Season")
	assert.Equal(t, getMatchup.Home.ID, homeTeam.ID, "Invalid Home")
	assert.Equal(t, getMatchup.Away.ID, awayTeam.ID, "Invalid Away")
	assert.Equal(t, getMatchup.Round, 2, "Invalid Round")
	assert.Equal(t, getMatchup.Start.Unix(), start.Unix(), "Invalid Start")
}

func TestNewSqliteStoreDeleteMatchup(t *testing.T) {
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
	start := time.Now()
	matchup := &data.Matchup{League: league, Season: *season, ID: "id", Home: *homeTeam, Away: *awayTeam, Round: 1, Start: start}
	err = store.Matchup().AddMatchup(matchup)
	assert.NoError(t, err, "Should not have error")
	err = store.Matchup().DeleteMatchup(matchup)
	assert.NoError(t, err, "Should not have error")
	getMatchup, err := store.Matchup().GetMatchup(&league, season, "id")
	assert.Error(t, err, "Should have error")
	assert.Nil(t, getMatchup, "Should be nil")
	getMatchups, err := store.Matchup().GetMatchups(&league, season)
	assert.NoError(t, err, "Should have error")
	assert.Equal(t, len(getMatchups), 0, "Should be zero")
}

func TestNewSqliteStoreGetMatchups(t *testing.T) {
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
	start := time.Now()
	matchup := &data.Matchup{League: league, Season: *season, ID: "id", Home: *homeTeam, Away: *awayTeam, Round: 1, Start: start}
	err = store.Matchup().AddMatchup(matchup)
	assert.NoError(t, err, "Should not have error")

	matchups, _ := store.Matchup().GetMatchups(&league, season)
	assert.NotNil(t, matchups, "Should not be nil")
	assert.Equal(t, len(matchups), 1, "Should be only one matchup")
	assert.Equal(t, matchups[0].League.ID, league.ID, "Should be the good league")
	assert.Equal(t, matchups[0].Season.Year, season.Year, "Should be the good season")
	assert.Equal(t, matchups[0].ID, matchup.ID, "Should be the good Matchup")
}

func TestNewSqliteStoreGetMatchupsEmpty(t *testing.T) {
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
	matchup := &data.Matchup{League: league, Season: *season, ID: "id", Round: 1}
	err = store.Matchup().AddMatchup(matchup)
	assert.NoError(t, err, "Should not have error")

	matchups, _ := store.Matchup().GetMatchups(&league, season)
	assert.NotNil(t, matchups, "Should not be nil")
	assert.Equal(t, len(matchups), 1, "Should be only one matchup")
	assert.Equal(t, matchups[0].League.ID, league.ID, "Should be the good league")
	assert.Equal(t, matchups[0].Season.Year, season.Year, "Should be the good season")
	assert.Equal(t, matchups[0].ID, matchup.ID, "Should be the good Matchup")
}
