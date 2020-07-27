package store

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"nhlpool.com/service/go/nhlpool/data"
)

// SqliteStorePrediction Is a venue data store for that keep it in Sqlite
type SqliteStorePrediction struct {
	database *sql.DB
	store    *SqliteStore
}

// NewSqliteStorePrediction Create a new sqlite store
func NewSqliteStorePrediction(database *sql.DB, store *SqliteStore) *SqliteStorePrediction {
	newStore := &SqliteStorePrediction{database: database, store: store}
	newStore.createTables()
	return newStore
}

func (st *SqliteStorePrediction) createTables() error {
	if !st.store.tableExist("prediction") {
		err := st.createTable()
		if err != nil {
			return err
		}
	}

	return nil
}

func (st *SqliteStorePrediction) createTable() error {
	statement, err := st.database.Prepare(`CREATE TABLE IF NOT EXISTS prediction
	(league_id TEXT NOT NULL, season_year INTEGER NOT NULL, player_id TEXT NOT NULL, matchup_id TEXT NOT NULL, winner TEXT, games INTEGER, PRIMARY KEY(league_id, season_year, player_id, matchup_id))`)
	if err != nil {
		fmt.Printf("createTable Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Printf("createTable Exec Err: %v\n", err)
		return err
	}
	fmt.Printf("createTable prediction\n")
	return nil
}

func (st *SqliteStorePrediction) cleanTable() error {
	fmt.Printf("cleanTable prediction\n")
	statement, err := st.database.Prepare("DROP TABLE prediction")
	if err != nil {
		fmt.Printf("cleanTable Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Printf("cleanTable Exec Err: %v\n", err)
		return err
	}
	return nil
}

// Clean Empty the store
func (st *SqliteStorePrediction) Clean() error {
	errPrediction := st.cleanTable()
	errCreate := st.createTables()
	if errPrediction != nil {
		return errPrediction
	}
	if errCreate != nil {
		return errCreate
	}
	return nil
}

// AddPrediction Add a new venue
func (st *SqliteStorePrediction) AddPrediction(prediction *data.Prediction) error {
	statement, err := st.database.Prepare(`INSERT INTO prediction (league_id, season_year, player_id, matchup_id, winner, games) VALUES (?, ?, ?, ?, ?, ?)`)
	if err != nil {
		fmt.Printf("AddPrediction Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(prediction.League.ID, prediction.Season.Year, prediction.Player.ID, prediction.Matchup.ID, prediction.Winner.ID, prediction.Games)
	if err != nil {
		fmt.Printf("AddPrediction Exec Err: %v\n", err)
		return err
	}
	return nil
}

// UpdatePrediction Update a venue info
func (st *SqliteStorePrediction) UpdatePrediction(prediction *data.Prediction) error {
	statement, err := st.database.Prepare("UPDATE prediction SET winner=?, games=? WHERE league_id=? AND season_year=? AND player_id=? AND matchup_id=?")
	if err != nil {
		fmt.Printf("UpdatePrediction Prepare Err: %v\n", err)
		return err
	}
	res, err := statement.Exec(prediction.Winner.ID, prediction.Games, prediction.League.ID, prediction.Season.Year, prediction.Player.ID, prediction.Matchup.ID)
	if err != nil {
		fmt.Printf("UpdatePrediction Exec Err: %v\n", err)
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("UpdatePrediction RowsAffected Err: %v\n", err)
		return err
	}
	if row != 1 {
		fmt.Printf("UpdatePrediction Do not update any row\n")
		return errors.New("Invalid Prediction")
	}
	return nil
}

// DeletePrediction Delete a venue
func (st *SqliteStorePrediction) DeletePrediction(prediction *data.Prediction) error {
	statement, err := st.database.Prepare("DELETE FROM prediction WHERE league_id=? AND season_year=? AND player_id=? AND matchup_id=?")
	if err != nil {
		fmt.Printf("DeletePrediction Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(prediction.League.ID, prediction.Season.Year, prediction.Player.ID, prediction.Matchup.ID)
	if err != nil {
		fmt.Printf("DeletePrediction Exec Err: %v\n", err)
		return err
	}
	return nil
}

// GetPrediction Get a venue
func (st *SqliteStorePrediction) GetPrediction(player *data.Player, matchup *data.Matchup, league *data.League, season *data.Season) (*data.Prediction, error) {
	row := st.database.QueryRow(`SELECT league_id, season_year, player_id, matchup_id, winner, games
	FROM prediction WHERE league_id=? AND season_year=? AND player_id=? AND matchup_id=?`, league.ID, season.Year, player.ID, matchup.ID)
	var leagueID string
	var seasonYear int
	var playerID string
	var matchupID string
	var winnerSTR string
	var games int
	if row != nil {
		prediction := &data.Prediction{}
		err := row.Scan(&leagueID, &seasonYear, &playerID, &matchupID, &winnerSTR, &games)
		if err != nil {
			fmt.Printf("GetPrediction Scan Err: %v\n", err)
			return nil, err
		}
		leagueObj, _ := st.store.League().GetLeague(leagueID)
		seasonObj, _ := st.store.Season().GetSeason(seasonYear, leagueObj)
		playerObj := st.store.Player().GetPlayer(playerID)
		matchupObj, _ := st.store.Matchup().GetMatchup(leagueObj, seasonObj, matchupID)
		team, _ := st.store.Team().GetTeam(winnerSTR, leagueObj)
		prediction.League = *leagueObj
		prediction.Season = *seasonObj
		prediction.Player = playerObj
		prediction.Matchup = matchupObj
		prediction.Winner = *team
		prediction.Games = games
		return prediction, nil
	}
	return nil, errors.New("Prediction not found")
}

// GetPredictions Return a list of all team
func (st *SqliteStorePrediction) GetPredictions(league *data.League, season *data.Season) ([]data.Prediction, error) {
	var predictions []data.Prediction
	rows, err := st.database.Query("SELECT league_id, season_year, player_id, matchup_id, winner, games FROM prediction WHERE league_id=? AND season_year=?", league.ID, season.Year)
	if err != nil {
		fmt.Printf("GetPredictions query Err: %v\n", err)
		return []data.Prediction{}, err
	}
	var leagueID string
	var seasonYear int
	var playerID string
	var matchupID string
	var winnerSTR string
	var games int
	for rows.Next() {
		prediction := &data.Prediction{}
		err := rows.Scan(&leagueID, &seasonYear, &playerID, &matchupID, &winnerSTR, &games)
		if err != nil {
			fmt.Printf("GetPrediction Scan Err: %v\n", err)
			return nil, err
		}
		league, _ := st.store.League().GetLeague(leagueID)
		season, _ := st.store.Season().GetSeason(seasonYear, league)
		player := st.store.Player().GetPlayer(playerID)
		matchup, _ := st.store.Matchup().GetMatchup(league, season, matchupID)
		team, _ := st.store.Team().GetTeam(winnerSTR, league)
		prediction.League = *league
		prediction.Season = *season
		prediction.Player = player
		prediction.Matchup = matchup
		prediction.Winner = *team
		prediction.Games = games

		predictions = append(predictions, *prediction)
	}
	rows.Close()
	return predictions, nil
}
