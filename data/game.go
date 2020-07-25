package data

import (
	"time"
)

// GameStateScheduled Game state scheduled
const GameStateScheduled = 1

// GameStateInProgress Game state in progress
const GameStateInProgress = 2

// GameStateFinished Game state finished
const GameStateFinished = 3

// GameTypeRegular Game type regular season
const GameTypeRegular = 1

// GameTypePlayoff Game type playoff
const GameTypePlayoff = 2

// GameTypePreparation Game type preparation
const GameTypePreparation = 3

// Game Define the information about a game
type Game struct {
	League   League    `json:"league"`
	Season   Season    `json:"season"`
	Home     Team      `json:"home"`
	Away     Team      `json:"away"`
	Date     time.Time `json:"date"`
	Type     int       `json:"type"`
	State    int       `json:"state"`
	HomeGoal int       `json:"home_goal"`
	AwayGoal int       `json:"away_goal"`
}
