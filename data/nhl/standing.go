package nhl

// Standing Represent the standing for a year
type Standing struct {
	Conferences []StandingRecord `json:"records"`
}

// StandingRecord Define the info for a standing record
type StandingRecord struct {
	StandingsType string       `json:"standingsType"`
	League        League       `json:"league"`
	Division      Division     `json:"division"`
	Conference    Conference   `json:"conference"`
	Season        string       `json:"season"`
	TeamRecords   []TeamRecord `json:"teamRecords"`
}

// LeagueStandingRecord Represent the league standing for a team
type LeagueStandingRecord struct {
	Wins   int    `json:"wins"`
	Losses int    `json:"losses"`
	OT     int    `json:"ot"`
	Type   string `json:"type"`
}

// StreakRecord Info about the recent streak for a team
type StreakRecord struct {
	StreakType   string `json:"streakType"`
	StreakNumber int    `json:"streakNumber"`
	StreakCode   string `json:"streakCode"`
}

// TeamRecord Info about the standing record for a team
type TeamRecord struct {
	Team               Team                 `json:"team"`
	LeagueRecord       LeagueStandingRecord `json:"leagueRecord"`
	Type               string               `json:"type"`
	RegulationWins     int                  `json:"regulationWins"`
	GoalsAgainst       int                  `json:"goalsAgainst"`
	GoalsScored        int                  `json:"goalsScored"`
	Points             int                  `json:"points"`
	DivisionRank       string               `json:"divisionRank"`
	DivisionL10Rank    string               `json:"divisionL10Rank"`
	DivisionRoadRank   string               `json:"divisionRoadRank"`
	DivisionHomeRank   string               `json:"divisionHomeRank"`
	ConferenceRank     string               `json:"conferenceRank"`
	ConferenceL10Rank  string               `json:"conferenceL10Rank"`
	ConferenceRoadRank string               `json:"conferenceRoadRank"`
	ConferenceHomeRank string               `json:"conferenceHomeRank"`
	LeagueRank         string               `json:"leagueRank"`
	LeagueL10Rank      string               `json:"leagueL10Rank"`
	LeagueRoadRank     string               `json:"leagueRoadRank"`
	LeagueHomeRank     string               `json:"leagueHomeRank"`
	WildCardRank       string               `json:"wildCardRank"`
	Row                int                  `json:"row"`
	GamesPlayed        int                  `json:"gamesPlayed"`
	Streak             StreakRecord         `json:"streak"`
	PointsPercentage   float64              `json:"pointsPercentage"`
	PpDivisionRank     string               `json:"ppDivisionRank"`
	PpConferenceRank   string               `json:"ppConferenceRank"`
	PpLeagueRank       string               `json:"ppLeagueRank"`
	LastUpdated        string               `json:"lastUpdated"`
}
