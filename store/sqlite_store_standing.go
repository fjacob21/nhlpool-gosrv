package store

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"nhlpool.com/service/go/nhlpool/data"
)

// SqliteStoreStanding Is a venue data store for that keep it in Sqlite
type SqliteStoreStanding struct {
	database *sql.DB
	store    *SqliteStore
}

// NewSqliteStoreStanding Create a new sqlite store
func NewSqliteStoreStanding(database *sql.DB, store *SqliteStore) *SqliteStoreStanding {
	newStore := &SqliteStoreStanding{database: database, store: store}
	newStore.createTables()
	return newStore
}

func (st *SqliteStoreStanding) createTables() error {
	if !st.store.tableExist("standing") {
		err := st.createTable()
		if err != nil {
			return err
		}
	}

	return nil
}

func (st *SqliteStoreStanding) createTable() error {
	statement, err := st.database.Prepare(`CREATE TABLE IF NOT EXISTS standing
	(league_id TEXT NOT NULL, season_year TEXT NOT NULL, team_id TEXT NOT NULL, points INTEGER,
		win INTEGER, losses INTEGER, ot INTEGER, games_played INTEGER, goals_against INTEGER, goals_scored INTEGER,
		ranks INTEGER, PRIMARY KEY(league_id, season_year, team_id))`)
	if err != nil {
		fmt.Printf("createTable Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Printf("createTable Exec Err: %v\n", err)
		return err
	}
	fmt.Printf("createTable standing\n")
	return nil
}

func (st *SqliteStoreStanding) cleanTable() error {
	fmt.Printf("cleanTable standing\n")
	statement, err := st.database.Prepare("DROP TABLE standing")
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
func (st *SqliteStoreStanding) Clean() error {
	errStanding := st.cleanTable()
	errCreate := st.createTables()
	if errStanding != nil {
		return errStanding
	}
	if errCreate != nil {
		return errCreate
	}
	return nil
}

// AddStanding Add a new venue
func (st *SqliteStoreStanding) AddStanding(standing *data.Standing) error {
	statement, err := st.database.Prepare(`INSERT INTO standing (league_id, season_year, team_id, points,
		win, losses, ot, games_played, goals_against, goals_scored, ranks) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		fmt.Printf("AddStanding Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(standing.League.ID, standing.Season.Year, standing.Team.ID, standing.Points, standing.Win, standing.Losses, standing.OT, standing.GamesPlayed, standing.GoalsAgainst, standing.GoalsScored, standing.Ranks)
	if err != nil {
		fmt.Printf("AddStanding Exec Err: %v\n", err)
		return err
	}
	return nil
}

// UpdateStanding Update a venue info
func (st *SqliteStoreStanding) UpdateStanding(standing *data.Standing) error {
	statement, err := st.database.Prepare("UPDATE standing SET points=?, win=?, losses=?, ot=?, games_played=?, goals_against=?, goals_scored=?, ranks=? WHERE league_id=? AND season_year=? AND team_id=?")
	if err != nil {
		fmt.Printf("UpdateStanding Prepare Err: %v\n", err)
		return err
	}
	res, err := statement.Exec(standing.Points, standing.Win, standing.Losses, standing.OT, standing.GamesPlayed, standing.GoalsAgainst, standing.GoalsScored, standing.Ranks, standing.League.ID, standing.Season.Year, standing.Team.ID)
	if err != nil {
		fmt.Printf("UpdateStanding Exec Err: %v\n", err)
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("UpdateStanding RowsAffected Err: %v\n", err)
		return err
	}
	if row != 1 {
		fmt.Printf("UpdateStanding Do not update any row\n")
		return errors.New("Invalid Standing")
	}
	return nil
}

// DeleteStanding Delete a venue
func (st *SqliteStoreStanding) DeleteStanding(standing *data.Standing) error {
	statement, err := st.database.Prepare("DELETE FROM standing WHERE league_id=? AND season_year=? AND team_id=?")
	if err != nil {
		fmt.Printf("DeleteStanding Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(standing.League.ID, standing.Season.Year, standing.Team.ID)
	if err != nil {
		fmt.Printf("DeleteStanding Exec Err: %v\n", err)
		return err
	}
	return nil
}

// GetStanding Get a venue
func (st *SqliteStoreStanding) GetStanding(team *data.Team, league *data.League, season *data.Season) (*data.Standing, error) {
	row := st.database.QueryRow(`SELECT league_id, season_year, team_id, points,
	win, losses, ot, games_played, goals_against, goals_scored, ranks
	FROM standing WHERE league_id=? AND season_year=? AND team_id=?`, league.ID, season.Year, team.ID)
	var leagueID string
	var seasonYear int
	var teamID string
	var points int
	var win int
	var losses int
	var ot int
	var gamesPlayed int
	var goalsAgainst int
	var goalsScored int
	var ranks int
	if row != nil {
		standing := &data.Standing{}
		err := row.Scan(&leagueID, &seasonYear, &teamID, &points, &win, &losses, &ot, &gamesPlayed, &goalsAgainst, &goalsScored, &ranks)
		if err != nil {
			fmt.Printf("GetStanding Scan Err: %v\n", err)
			return nil, err
		}
		league, _ := st.store.League().GetLeague(leagueID)
		season, _ := st.store.Season().GetSeason(seasonYear, league)
		team, _ := st.store.Team().GetTeam(teamID, league)
		standing.League = *league
		standing.Season = *season
		standing.Team = *team
		standing.Points = points
		standing.Win = win
		standing.Losses = losses
		standing.OT = ot
		standing.GamesPlayed = gamesPlayed
		standing.GoalsAgainst = goalsAgainst
		standing.GoalsScored = goalsScored
		standing.Ranks = ranks

		return standing, nil
	}
	return nil, errors.New("Standing not found")
}

// GetStandings Return a list of all team
func (st *SqliteStoreStanding) GetStandings(league *data.League, season *data.Season) ([]data.Standing, error) {
	var standings []data.Standing
	rows, err := st.database.Query("SELECT league_id, season_year, team_id, points, win, losses, ot, games_played, goals_against, goals_scored, ranks FROM standing WHERE league_id=? AND season_year=?", league.ID, season.Year)
	if err != nil {
		fmt.Printf("GetStandings query Err: %v\n", err)
		return []data.Standing{}, err
	}
	var leagueID string
	var seasonYear int
	var teamID string
	var points int
	var win int
	var losses int
	var ot int
	var gamesPlayed int
	var goalsAgainst int
	var goalsScored int
	var ranks int
	for rows.Next() {
		standing := &data.Standing{}
		err := rows.Scan(&leagueID, &seasonYear, &teamID, &points, &win, &losses, &ot, &gamesPlayed, &goalsAgainst, &goalsScored, &ranks)
		if err != nil {
			fmt.Printf("GetStanding Scan Err: %v\n", err)
			return nil, err
		}
		standingLeague, err := st.store.League().GetLeague(leagueID)
		if err != nil {
			fmt.Printf("GetStanding GetLeague LeagueID:%v League:%v Err: %v\n", leagueID, standingLeague, err)
		}

		standingSeason, err := st.store.Season().GetSeason(seasonYear, league)
		if err != nil {
			fmt.Printf("GetStanding GetSeason year:%v season:%v Err: %v\n", seasonYear, standingSeason, err)
		}
		team, err := st.store.Team().GetTeam(teamID, league)
		if err != nil {
			fmt.Printf("GetStanding GetTeam teamid:%v team:%v Err: %v\n", teamID, team, err)
		}
		standing.League = *league
		standing.Season = *season
		if team != nil {
			standing.Team = *team
		}
		standing.Points = points
		standing.Win = win
		standing.Losses = losses
		standing.OT = ot
		standing.GamesPlayed = gamesPlayed
		standing.GoalsAgainst = goalsAgainst
		standing.GoalsScored = goalsScored
		standing.Ranks = ranks

		standings = append(standings, *standing)
	}
	rows.Close()
	return standings, nil
}
