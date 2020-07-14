package nhl

import "time"

// Game Represent the info about a game
type Game struct {
	GamePk    int        `json:"gamePk"`
	Link      string     `json:"link"`
	GameType  string     `json:"gameType"`
	Season    string     `json:"season"`
	GameDate  *time.Time `json:"gameDate"`
	Status    GameStatus `json:"status"`
	Teams     GameTeams  `json:"teams"`
	Linescore Linescore  `json:"linescore"`
	Venue     Venue      `json:"venue"`
}

// GameStatus Is the game status
type GameStatus struct {
	AbstractGameState string `json:"abstractGameState"`
	CodedGameState    string `json:"codedGameState"`
	DetailedState     string `json:"detailedState"`
	StatusCode        string `json:"statusCode"`
	StartTimeTBD      bool   `json:"startTimeTBD"`
}

// GameTeams Teams matching for a game
type GameTeams struct {
	Away GameTeam `json:"away"`
	Home GameTeam `json:"Home"`
}

// GameTeam Team info for a game
type GameTeam struct {
	LeagueRecord League `json:"leagueRecord"`
	Score        int    `json:"score"`
	Team         Team   `json:"team"`
}

// Linescore Info about stats of the game
type Linescore struct {
	CurrentPeriod              int              `json:"currentPeriod"`
	CurrentPeriodOrdinal       string           `json:"currentPeriodOrdinal"`
	CurrentPeriodTimeRemaining string           `json:"currentPeriodTimeRemaining"`
	Periods                    []Periode        `json:"periods"`
	ShootoutInfo               ShootoutInfo     `json:"shootoutInfo"`
	Teams                      LinescoreTeams   `json:"teams"`
	PowerPlayStrength          string           `json:"powerPlayStrength"`
	HasShootout                bool             `json:"hasShootout"`
	IntermissionInfo           IntermissionInfo `json:"intermissionInfo"`
	PowerPlayInfo              PowerPlayInfo    `json:"powerPlayInfo"`
}

// PowerPlayInfo Power play info
type PowerPlayInfo struct {
	SituationTimeRemaining int  `json:"situationTimeRemaining"`
	SituationTimeElapsed   int  `json:"situationTimeElapsed"`
	InSituation            bool `json:"inSituation"`
}

// IntermissionInfo Intermission info
type IntermissionInfo struct {
	IntermissionTimeRemaining int  `json:"intermissionTimeRemaining"`
	IntermissionTimeElapsed   int  `json:"intermissionTimeElapsed"`
	InIntermission            bool `json:"inIntermission"`
}

// Periode Info about the stats of a periode
type Periode struct {
	PeriodType string            `json:"periodType"`
	StartTime  time.Time         `json:"startTime"`
	EndTime    time.Time         `json:"endTime"`
	Num        int               `json:"num"`
	OrdinalNum string            `json:"ordinalNum"`
	Home       PeriodeTeamResult `json:"home"`
	Away       PeriodeTeamResult `json:"away"`
}

// PeriodeTeamResult Periode stats for a team
type PeriodeTeamResult struct {
	Goals       int    `json:"goals"`
	ShotsOnGoal int    `json:"shotsOnGoal"`
	RinkSide    string `json:"rinkSide"`
}

// ShootoutInfo Info about the shootout
type ShootoutInfo struct {
	Away ShootoutInfoTeamResult `json:"away"`
	Home ShootoutInfoTeamResult `json:"home"`
}

// ShootoutInfoTeamResult Shootout team stats
type ShootoutInfoTeamResult struct {
	Cores    int `json:"scores"`
	Attempts int `json:"attempts"`
}

// LinescoreTeams Teams game stats
type LinescoreTeams struct {
	Away LinescoreTeamResult `json:"away"`
	Home LinescoreTeamResult `json:"home"`
}

// LinescoreTeamResult Team game stats
type LinescoreTeamResult struct {
	Team         Team `json:"team"`
	Goals        int  `json:"goals"`
	ShotsOnGoal  int  `json:"shotsOnGoal"`
	GoaliePulled bool `json:"goaliePulled"`
	NumSkaters   int  `json:"numSkaters"`
	PowerPlay    bool `json:"powerPlay"`
}
