package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"nhlpool.com/service/go/nhlpool/data"
)

func TestNewSqliteStoreAddSeason(t *testing.T) {
	store := NewSqliteStore()
	defer store.Close()
	store.Clean()
	assert.NotNil(t, store, "Should not be nil")
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.League().AddLeague(&league)
	season := &data.Season{Year: 2000, League: &league}
	err = store.Season().AddSeason(season)
	assert.NoError(t, err, "Should not have error")
	getSeason, err := store.Season().GetSeason(season.Year, &league)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getSeason, "Should not be nil")
	assert.Equal(t, getSeason.Year, season.Year, "Invalid year")
	assert.Equal(t, getSeason.League.ID, league.ID, "Invalid league")
}

func TestNewSqliteStoreDeleteSeason(t *testing.T) {
	store := NewSqliteStore()
	defer store.Close()
	store.Clean()
	assert.NotNil(t, store, "Should not be nil")
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.League().AddLeague(&league)
	season := &data.Season{Year: 2000, League: &league}
	err = store.Season().AddSeason(season)
	assert.NoError(t, err, "Should not have error")
	getSeason, err := store.Season().GetSeason(season.Year, &league)
	assert.NoError(t, err, "Should not have error")
	assert.NotNil(t, getSeason, "Should not be nil")
	err = store.Season().DeleteSeason(season)
	assert.NoError(t, err, "Should not have error")
	getSeason, err = store.Season().GetSeason(season.Year, &league)
	assert.Error(t, err, "Should have error")
	assert.Nil(t, getSeason, "Should be nil")
}

func TestNewSqliteStoreGetSeasons(t *testing.T) {
	store := NewSqliteStore()
	defer store.Close()
	store.Clean()
	assert.NotNil(t, store, "Should not be nil")
	league := data.League{ID: "id", Name: "name", Description: "description", Website: "website"}
	err := store.League().AddLeague(&league)
	seasons, _ := store.Season().GetSeasons(&league)
	assert.Equal(t, len(seasons), 0, "There should not have any team")

	season := &data.Season{Year: 2000, League: &league}
	err = store.Season().AddSeason(season)
	assert.NoError(t, err, "Should not have error")

	seasons, _ = store.Season().GetSeasons(&league)
	assert.NotNil(t, seasons, "Should not be nil")
	assert.Equal(t, len(seasons), 1, "Should be only one team")
	assert.Equal(t, seasons[0].Year, season.Year, "Should be the good year")
}
