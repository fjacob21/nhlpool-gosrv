package store

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"nhlpool.com/service/go/nhlpool/data"
)

// SqliteStoreWinner Is a venue data store for that keep it in Sqlite
type SqliteStoreWinner struct {
	database *sql.DB
	store    *SqliteStore
}

// NewSqliteStoreWinner Create a new sqlite store
func NewSqliteStoreWinner(database *sql.DB, store *SqliteStore) *SqliteStoreWinner {
	newStore := &SqliteStoreWinner{database: database, store: store}
	newStore.createTables()
	return newStore
}

func (st *SqliteStoreWinner) createTables() error {
	if !st.store.tableExist("winner") {
		err := st.createTable()
		if err != nil {
			return err
		}
	}

	return nil
}

func (st *SqliteStoreWinner) createTable() error {
	statement, err := st.database.Prepare(`CREATE TABLE IF NOT EXISTS winner
	(league_id TEXT NOT NULL, season_year INTEGER NOT NULL, player_id TEXT NOT NULL, winner TEXT, PRIMARY KEY(league_id, season_year, player_id))`)
	if err != nil {
		fmt.Printf("createTable Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Printf("createTable Exec Err: %v\n", err)
		return err
	}
	fmt.Printf("createTable winner\n")
	return nil
}

func (st *SqliteStoreWinner) cleanTable() error {
	fmt.Printf("cleanTable winner\n")
	statement, err := st.database.Prepare("DROP TABLE winner")
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
func (st *SqliteStoreWinner) Clean() error {
	errWinner := st.cleanTable()
	errCreate := st.createTables()
	if errWinner != nil {
		return errWinner
	}
	if errCreate != nil {
		return errCreate
	}
	return nil
}

// AddWinner Add a new venue
func (st *SqliteStoreWinner) AddWinner(winner *data.Winner) error {
	statement, err := st.database.Prepare(`INSERT INTO winner (league_id, season_year, player_id, winner) VALUES (?, ?, ?, ?)`)
	if err != nil {
		fmt.Printf("AddWinner Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(winner.League.ID, winner.Season.Year, winner.Player.ID, winner.Winner.ID)
	if err != nil {
		fmt.Printf("AddWinner Exec Err: %v\n", err)
		return err
	}
	return nil
}

// UpdateWinner Update a venue info
func (st *SqliteStoreWinner) UpdateWinner(winner *data.Winner) error {
	statement, err := st.database.Prepare("UPDATE winner SET winner=? WHERE league_id=? AND season_year=? AND player_id=?")
	if err != nil {
		fmt.Printf("UpdateWinner Prepare Err: %v\n", err)
		return err
	}
	res, err := statement.Exec(winner.Winner.ID, winner.League.ID, winner.Season.Year, winner.Player.ID)
	if err != nil {
		fmt.Printf("UpdateWinner Exec Err: %v\n", err)
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("UpdateWinner RowsAffected Err: %v\n", err)
		return err
	}
	if row != 1 {
		fmt.Printf("UpdateWinner Do not update any row\n")
		return errors.New("Invalid Winner")
	}
	return nil
}

// DeleteWinner Delete a venue
func (st *SqliteStoreWinner) DeleteWinner(winner *data.Winner) error {
	statement, err := st.database.Prepare("DELETE FROM winner WHERE league_id=? AND season_year=? AND player_id=?")
	if err != nil {
		fmt.Printf("DeleteWinner Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(winner.League.ID, winner.Season.Year, winner.Player.ID)
	if err != nil {
		fmt.Printf("DeleteWinner Exec Err: %v\n", err)
		return err
	}
	return nil
}

// GetWinner Get a venue
func (st *SqliteStoreWinner) GetWinner(player *data.Player, league *data.League, season *data.Season) (*data.Winner, error) {
	row := st.database.QueryRow(`SELECT league_id, season_year, player_id, winner
	FROM winner WHERE league_id=? AND season_year=? AND player_id=?`, league.ID, season.Year, player.ID)
	var leagueID string
	var seasonYear int
	var playerID string
	var winnerSTR string
	if row != nil {
		winner := &data.Winner{}
		err := row.Scan(&leagueID, &seasonYear, &playerID, &winnerSTR)
		if err != nil {
			fmt.Printf("GetWinner Scan Err: %v\n", err)
			return nil, err
		}
		league, _ := st.store.League().GetLeague(leagueID)
		season, _ := st.store.Season().GetSeason(seasonYear, league)
		player := st.store.Player().GetPlayer(playerID)
		team, _ := st.store.Team().GetTeam(winnerSTR, league)
		winner.League = *league
		winner.Season = *season
		winner.Player = player
		winner.Winner = *team
		return winner, nil
	}
	return nil, errors.New("Winner not found")
}

// GetWinners Return a list of all team
func (st *SqliteStoreWinner) GetWinners(league *data.League, season *data.Season) ([]data.Winner, error) {
	var winners []data.Winner
	rows, err := st.database.Query("SELECT league_id, season_year, player_id, winner FROM winner WHERE league_id=? AND season_year=?", league.ID, season.Year)
	if err != nil {
		fmt.Printf("GetWinners query Err: %v\n", err)
		return []data.Winner{}, err
	}
	var leagueID string
	var seasonYear int
	var playerID string
	var winnerSTR string
	for rows.Next() {
		winner := &data.Winner{}
		err := rows.Scan(&leagueID, &seasonYear, &playerID, &winnerSTR)
		if err != nil {
			fmt.Printf("GetWinner Scan Err: %v\n", err)
			return nil, err
		}
		winnerLeague, err := st.store.League().GetLeague(leagueID)
		if err != nil {
			fmt.Printf("GetWinner GetLeague LeagueID:%v League:%v Err: %v\n", leagueID, winnerLeague, err)
		}

		winnerSeason, err := st.store.Season().GetSeason(seasonYear, league)
		if err != nil {
			fmt.Printf("GetWinner GetSeason year:%v season:%v Err: %v\n", seasonYear, winnerSeason, err)
		}
		player := st.store.Player().GetPlayer(playerID)
		team, _ := st.store.Team().GetTeam(winnerSTR, league)
		winner.League = *league
		winner.Season = *season
		winner.Player = player
		winner.Winner = *team

		winners = append(winners, *winner)
	}
	rows.Close()
	return winners, nil
}
