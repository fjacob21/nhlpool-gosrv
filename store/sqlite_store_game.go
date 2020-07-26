package store

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"nhlpool.com/service/go/nhlpool/data"
)

// SqliteStoreGame Is a game data store for that keep it in Sqlite
type SqliteStoreGame struct {
	database *sql.DB
	store    *SqliteStore
}

// NewSqliteStoreGame Create a new sqlite store
func NewSqliteStoreGame(database *sql.DB, store *SqliteStore) *SqliteStoreGame {
	newStore := &SqliteStoreGame{database: database, store: store}
	newStore.createTables()
	return newStore
}

func (st *SqliteStoreGame) createTables() error {
	if !st.store.tableExist("game") {
		err := st.createTable()
		if err != nil {
			return err
		}
	}

	return nil
}

func (st *SqliteStoreGame) createTable() error {
	statement, err := st.database.Prepare(`CREATE TABLE IF NOT EXISTS game
	(league_id TEXT NOT NULL, season_year INTEGER NOT NULL, home TEXT NOT NULL, away TEXT NOT NULL, date TEXT NOT NULL,
		type INTEGER, state INTEGER, home_goal INTEGER, away_goal INTEGER, PRIMARY KEY(league_id, season_year, home, away, date))`)
	if err != nil {
		fmt.Printf("createTable Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		fmt.Printf("createTable Exec Err: %v\n", err)
		return err
	}
	fmt.Printf("createTable game\n")
	return nil
}

func (st *SqliteStoreGame) cleanTable() error {
	fmt.Printf("cleanTable game\n")
	statement, err := st.database.Prepare("DROP TABLE game")
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
func (st *SqliteStoreGame) Clean() error {
	errGame := st.cleanTable()
	errCreate := st.createTables()
	if errGame != nil {
		return errGame
	}
	if errCreate != nil {
		return errCreate
	}
	return nil
}

// AddGame Add a new venue
func (st *SqliteStoreGame) AddGame(game *data.Game) error {
	statement, err := st.database.Prepare(`INSERT INTO game (league_id, season_year, home,
		away, date, type, state, home_goal, away_goal) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		fmt.Printf("AddGame Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(game.League.ID, game.Season.Year, game.Home.ID, game.Away.ID, game.Date.Format(time.RFC3339), game.Type, game.State, game.HomeGoal, game.AwayGoal)
	if err != nil {
		fmt.Printf("AddGame Exec Err: %v\n", err)
		return err
	}
	return nil
}

// UpdateGame Update a venue info
func (st *SqliteStoreGame) UpdateGame(game *data.Game) error {
	statement, err := st.database.Prepare("UPDATE game SET state=?, home_goal=?, away_goal=? WHERE league_id=? AND season_year=? AND home=? AND away=? AND date=?")
	if err != nil {
		fmt.Printf("UpdateGame Prepare Err: %v\n", err)
		return err
	}
	res, err := statement.Exec(game.State, game.HomeGoal, game.AwayGoal, game.League.ID, game.Season.Year, game.Home.ID, game.Away.ID, game.Date.Format(time.RFC3339))
	if err != nil {
		fmt.Printf("UpdateGame Exec Err: %v\n", err)
		return err
	}
	row, err := res.RowsAffected()
	if err != nil {
		fmt.Printf("UpdateGame RowsAffected Err: %v\n", err)
		return err
	}
	if row != 1 {
		fmt.Printf("UpdateGame Do not update any row\n")
		return errors.New("Invalid Game")
	}
	return nil
}

// DeleteGame Delete a venue
func (st *SqliteStoreGame) DeleteGame(game *data.Game) error {
	statement, err := st.database.Prepare("DELETE FROM game WHERE league_id=? AND season_year=? AND home=? AND away=? AND date=?")
	if err != nil {
		fmt.Printf("DeleteGame Prepare Err: %v\n", err)
		return err
	}
	_, err = statement.Exec(game.League.ID, game.Season.Year, game.Home.ID, game.Away.ID, game.Date.Format(time.RFC3339))
	if err != nil {
		fmt.Printf("DeleteGame Exec Err: %v\n", err)
		return err
	}
	return nil
}

// GetGame Get a venue
func (st *SqliteStoreGame) GetGame(league *data.League, season *data.Season, home *data.Team, away *data.Team, date time.Time) (*data.Game, error) {
	row := st.database.QueryRow(`SELECT league_id, season_year, home,
	away, date, type, state, home_goal, away_goal
	FROM game WHERE league_id=? AND season_year=? AND home=? AND away=? AND date=?`, league.ID, season.Year, home.ID, away.ID, date.Format(time.RFC3339))
	var leagueID string
	var seasonYear int
	var homeID string
	var awayID string
	var dateStr string
	var gameType int
	var state int
	var homeGoal int
	var awayGoal int
	if row != nil {
		game := &data.Game{}
		err := row.Scan(&leagueID, &seasonYear, &homeID, &awayID, &dateStr, &gameType, &state, &homeGoal, &awayGoal)
		if err != nil {
			fmt.Printf("GetGame Scan Err: %v\n", err)
			return nil, err
		}
		league, _ := st.store.League().GetLeague(leagueID)
		season, _ := st.store.Season().GetSeason(seasonYear, league)
		home, _ := st.store.Team().GetTeam(homeID, league)
		away, _ := st.store.Team().GetTeam(awayID, league)
		game.League = *league
		game.Season = *season
		game.Home = *home
		game.Away = *away
		game.Date, err = time.Parse(time.RFC3339, dateStr)
		if err != nil {
			fmt.Printf("GetGame Invalid time Err: %v\n", err)

		}
		game.Type = gameType
		game.State = state
		game.HomeGoal = homeGoal
		game.AwayGoal = awayGoal
		return game, nil
	}
	return nil, errors.New("Game not found")
}

// GetGames Return a list of all team
func (st *SqliteStoreGame) GetGames(league *data.League, season *data.Season, home *data.Team, away *data.Team) ([]data.Game, error) {
	var games []data.Game
	rows, err := st.database.Query(`SELECT  league_id, season_year, home,
	away, date, type, state, home_goal, away_goal FROM game WHERE league_id=? AND season_year=? AND home=? AND away=?`, league.ID, season.Year, home.ID, away.ID)
	if err != nil {
		fmt.Printf("GetGames query Err: %v\n", err)
		return []data.Game{}, err
	}
	var leagueID string
	var seasonYear int
	var homeID string
	var awayID string
	var dateStr string
	var gameType int
	var state int
	var homeGoal int
	var awayGoal int
	for rows.Next() {
		game := &data.Game{}
		err := rows.Scan(&leagueID, &seasonYear, &homeID, &awayID, &dateStr, &gameType, &state, &homeGoal, &awayGoal)
		if err != nil {
			fmt.Printf("GetGame Scan Err: %v\n", err)
			return nil, err
		}
		gameLeague, err := st.store.League().GetLeague(leagueID)
		if err != nil {
			fmt.Printf("GetGame GetLeague LeagueID:%v League:%v Err: %v\n", leagueID, gameLeague, err)
		}

		gameSeason, err := st.store.Season().GetSeason(seasonYear, league)
		if err != nil {
			fmt.Printf("GetGame GetSeason year:%v season:%v Err: %v\n", seasonYear, gameSeason, err)
		}
		home, _ := st.store.Team().GetTeam(homeID, league)
		away, _ := st.store.Team().GetTeam(awayID, league)
		game.League = *league
		game.Season = *season
		game.Home = *home
		game.Away = *away
		game.Date, err = time.Parse(time.RFC3339, dateStr)
		if err != nil {
			fmt.Printf("GetGame Invalid time Err: %v\n", err)
		}
		game.Type = gameType
		game.State = state
		game.HomeGoal = homeGoal
		game.AwayGoal = awayGoal

		games = append(games, *game)
	}
	rows.Close()
	return games, nil
}

// GetSeasonGames Return a list of all games
func (st *SqliteStoreGame) GetSeasonGames(league *data.League, season *data.Season, home *data.Team, away *data.Team) ([]data.Game, error) {
	var games []data.Game
	rows, err := st.database.Query(`SELECT  league_id, season_year, home,
	away, date, type, state, home_goal, away_goal FROM game WHERE league_id=? AND season_year=? AND home=? AND away=? AND Type=?`, league.ID, season.Year, home.ID, away.ID, data.GameTypeRegular)
	if err != nil {
		fmt.Printf("GetGames query Err: %v\n", err)
		return []data.Game{}, err
	}
	var leagueID string
	var seasonYear int
	var homeID string
	var awayID string
	var dateStr string
	var gameType int
	var state int
	var homeGoal int
	var awayGoal int
	for rows.Next() {
		game := &data.Game{}
		err := rows.Scan(&leagueID, &seasonYear, &homeID, &awayID, &dateStr, &gameType, &state, &homeGoal, &awayGoal)
		if err != nil {
			fmt.Printf("GetGame Scan Err: %v\n", err)
			return nil, err
		}
		gameLeague, err := st.store.League().GetLeague(leagueID)
		if err != nil {
			fmt.Printf("GetGame GetLeague LeagueID:%v League:%v Err: %v\n", leagueID, gameLeague, err)
		}

		gameSeason, err := st.store.Season().GetSeason(seasonYear, league)
		if err != nil {
			fmt.Printf("GetGame GetSeason year:%v season:%v Err: %v\n", seasonYear, gameSeason, err)
		}
		home, _ := st.store.Team().GetTeam(homeID, league)
		away, _ := st.store.Team().GetTeam(awayID, league)
		game.League = *league
		game.Season = *season
		game.Home = *home
		game.Away = *away
		game.Date, err = time.Parse(time.RFC3339, dateStr)
		if err != nil {
			fmt.Printf("GetGame Invalid time Err: %v\n", err)
		}
		game.Type = gameType
		game.State = state
		game.HomeGoal = homeGoal
		game.AwayGoal = awayGoal

		games = append(games, *game)
	}
	rows.Close()
	return games, nil
}

// GetPlayoffGames Return a list of all games
func (st *SqliteStoreGame) GetPlayoffGames(league *data.League, season *data.Season, home *data.Team, away *data.Team) ([]data.Game, error) {
	var games []data.Game
	rows, err := st.database.Query(`SELECT  league_id, season_year, home,
	away, date, type, state, home_goal, away_goal FROM game WHERE league_id=? AND season_year=? AND home=? AND away=? AND Type=?`, league.ID, season.Year, home.ID, away.ID, data.GameTypePlayoff)
	if err != nil {
		fmt.Printf("GetGames query Err: %v\n", err)
		return []data.Game{}, err
	}
	var leagueID string
	var seasonYear int
	var homeID string
	var awayID string
	var dateStr string
	var gameType int
	var state int
	var homeGoal int
	var awayGoal int
	for rows.Next() {
		game := &data.Game{}
		err := rows.Scan(&leagueID, &seasonYear, &homeID, &awayID, &dateStr, &gameType, &state, &homeGoal, &awayGoal)
		if err != nil {
			fmt.Printf("GetGame Scan Err: %v\n", err)
			return nil, err
		}
		gameLeague, err := st.store.League().GetLeague(leagueID)
		if err != nil {
			fmt.Printf("GetGame GetLeague LeagueID:%v League:%v Err: %v\n", leagueID, gameLeague, err)
		}

		gameSeason, err := st.store.Season().GetSeason(seasonYear, league)
		if err != nil {
			fmt.Printf("GetGame GetSeason year:%v season:%v Err: %v\n", seasonYear, gameSeason, err)
		}
		home, _ := st.store.Team().GetTeam(homeID, league)
		away, _ := st.store.Team().GetTeam(awayID, league)
		game.League = *league
		game.Season = *season
		game.Home = *home
		game.Away = *away
		game.Date, err = time.Parse(time.RFC3339, dateStr)
		if err != nil {
			fmt.Printf("GetGame Invalid time Err: %v\n", err)
		}
		game.Type = gameType
		game.State = state
		game.HomeGoal = homeGoal
		game.AwayGoal = awayGoal

		games = append(games, *game)
	}
	rows.Close()
	return games, nil
}

// GetAllGames Return a list of all games
func (st *SqliteStoreGame) GetAllGames(league *data.League, season *data.Season) ([]data.Game, error) {
	var games []data.Game
	rows, err := st.database.Query(`SELECT  league_id, season_year, home,
	away, date, type, state, home_goal, away_goal FROM game WHERE league_id=? AND season_year=?`, league.ID, season.Year)
	if err != nil {
		fmt.Printf("GetGames query Err: %v\n", err)
		return []data.Game{}, err
	}
	var leagueID string
	var seasonYear int
	var homeID string
	var awayID string
	var dateStr string
	var gameType int
	var state int
	var homeGoal int
	var awayGoal int
	for rows.Next() {
		game := &data.Game{}
		err := rows.Scan(&leagueID, &seasonYear, &homeID, &awayID, &dateStr, &gameType, &state, &homeGoal, &awayGoal)
		if err != nil {
			fmt.Printf("GetGame Scan Err: %v\n", err)
			return nil, err
		}
		gameLeague, err := st.store.League().GetLeague(leagueID)
		if err != nil {
			fmt.Printf("GetGame GetLeague LeagueID:%v League:%v Err: %v\n", leagueID, gameLeague, err)
		}

		gameSeason, err := st.store.Season().GetSeason(seasonYear, league)
		if err != nil {
			fmt.Printf("GetGame GetSeason year:%v season:%v Err: %v\n", seasonYear, gameSeason, err)
		}
		home, _ := st.store.Team().GetTeam(homeID, league)
		away, _ := st.store.Team().GetTeam(awayID, league)
		game.League = *league
		game.Season = *season
		game.Home = *home
		game.Away = *away
		game.Date, err = time.Parse(time.RFC3339, dateStr)
		if err != nil {
			fmt.Printf("GetGame Invalid time Err: %v\n", err)
		}
		game.Type = gameType
		game.State = state
		game.HomeGoal = homeGoal
		game.AwayGoal = awayGoal

		games = append(games, *game)
	}
	rows.Close()
	return games, nil
}
