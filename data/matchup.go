package data

import (
	"time"
)

// Matchup Define the information about round matchup
type Matchup struct {
	ID           string        `json:"id"`
	League       League        `json:"league"`
	Season       Season        `json:"season"`
	Home         Team          `json:"home"`
	Away         Team          `json:"away"`
	Round        int           `json:"round"`
	Start        time.Time     `json:"start"`
	PlayoffGames []Game        `json:"playoffGames"`
	SeasonGames  []Game        `json:"seasonGames"`
	Result       MatchupResult `json:"result"`
}

// CalculateResult Calculate results
func (m *Matchup) CalculateResult() {
	m.Result.HomeWin = 0
	m.Result.AwayWin = 0
	for _, game := range m.PlayoffGames {
		inverse := false
		if game.Home.ID == m.Away.ID {
			inverse = true
		}
		if game.State == GameStateFinished {
			if game.HomeGoal > game.AwayGoal {
				if inverse {
					m.Result.AwayWin++
				} else {
					m.Result.HomeWin++
				}
			}
		}
	}
}
