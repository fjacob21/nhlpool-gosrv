package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"nhlpool.com/service/go/nhlpool/data"
)

// Client Object to call the service
type Client struct {
	server    string
	port      int
	url       string
	sessionID string
	user      string
}

// NewClient create a new client
func NewClient(server string, port int) *Client {
	url := fmt.Sprintf("http://%v:%v", server, port)
	return &Client{server: server, port: port, url: url, sessionID: ""}
}

// Login using the specified credential
func (c *Client) Login(user string, password string) error {
	body := &data.LoginRequest{
		Password: password,
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	url := fmt.Sprintf("%v/player/%v/login/", c.url, user)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var reply data.LoginReply
	bodyReply, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bodyReply, &reply)
	if reply.Result.Code != 0 {
		return errors.New("Cannot login")
	}
	c.sessionID = reply.SessionID
	c.user = user
	return nil
}

// Logout current session
func (c *Client) Logout() error {
	if c.sessionID == "" {
		return nil
	}
	body := &data.LogoutRequest{
		SessionID: c.sessionID,
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	url := fmt.Sprintf("%v/player/%v/logout/", c.url, c.user)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var reply data.LogoutReply
	bodyReply, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bodyReply, &reply)
	if reply.Result.Code != 0 {
		return errors.New("Cannot login")
	}
	c.sessionID = ""
	c.user = ""
	return nil
}

// Import a player from backup
func (c *Client) Import(id string, player data.BackupPlayer) error {
	if c.sessionID == "" {
		return errors.New("Need to be logged")
	}
	body := data.ImportPlayerRequest{SessionID: c.sessionID}
	body.Player.Admin = player.Admin
	body.Player.Email = player.Email
	body.Player.ID = id
	body.Player.Name = player.Name
	body.Player.Password = player.Psw
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	url := fmt.Sprintf("%v/player/import/", c.url)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	}
	defer res.Body.Close()
	var importReply data.ImportPlayerReply
	bodyReply, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bodyReply, &importReply)
	if importReply.Result.Code != 0 {
		return errors.New("Cannot Import")
	}
	return nil
}

// AddLeague add a league
func (c *Client) AddLeague(id string, name string, description string, website string) error {
	if c.sessionID == "" {
		return errors.New("Need to be logged")
	}
	body := data.AddLeagueRequest{}
	body.ID = id
	body.Name = name
	body.Description = description
	body.Website = website
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	url := fmt.Sprintf("%v/league/", c.url)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	}
	defer res.Body.Close()
	var addLeagueReply data.AddLeagueReply
	bodyReply, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bodyReply, &addLeagueReply)
	if addLeagueReply.Result.Code != 0 {
		return errors.New("Cannot Add")
	}
	return nil
}

// AddTeam add a team
func (c *Client) AddTeam(
	leagueID string,
	id string,
	abbreviation string,
	name string,
	fullname string,
	city string,
	active bool,
	creationYear string,
	website string,
	venueID string,
	venueCity string,
	venueName string,
	venueTimezone string,
	venueAddress string,
) error {
	if c.sessionID == "" {
		return errors.New("Need to be logged")
	}
	body := data.AddTeamRequest{}
	body.ID = id
	body.Abbreviation = abbreviation
	body.Name = name
	body.Fullname = fullname
	body.City = city
	body.Active = active
	body.CreationYear = creationYear
	body.Website = website
	body.Venue.ID = venueID
	body.Venue.City = venueCity
	body.Venue.Name = venueName
	body.Venue.Timezone = venueTimezone
	body.Venue.Address = venueAddress

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	url := fmt.Sprintf("%v/league/%v/team/", c.url, leagueID)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	}
	defer res.Body.Close()
	var addTeamReply data.AddTeamReply
	bodyReply, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bodyReply, &addTeamReply)
	if addTeamReply.Result.Code != 0 {
		return errors.New("Cannot Add")
	}
	return nil
}

// AddSeason add a team
func (c *Client) AddSeason(
	leagueID string,
	year int,
) error {
	if c.sessionID == "" {
		return errors.New("Need to be logged")
	}
	body := data.AddSeasonRequest{}
	body.Year = year

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	url := fmt.Sprintf("%v/league/%v/season/", c.url, leagueID)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	}
	defer res.Body.Close()
	var addSeasonReply data.AddSeasonReply
	bodyReply, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bodyReply, &addSeasonReply)
	if addSeasonReply.Result.Code != 0 {
		return errors.New("Cannot Add")
	}
	return nil
}

// AddStanding add a standing
func (c *Client) AddStanding(
	leagueID string,
	year int,
	teamID string,
	points int,
	win int,
	losses int,
	ot int,
	gamesPlayed int,
	goalsAgainst int,
	goalsScored int,
	ranks int,
) error {
	if c.sessionID == "" {
		return errors.New("Need to be logged")
	}
	body := data.AddStandingRequest{}
	body.TeamID = teamID
	body.Points = points
	body.Win = win
	body.Losses = losses
	body.OT = ot
	body.GamesPlayed = gamesPlayed
	body.GoalsAgainst = goalsAgainst
	body.GoalsScored = goalsScored
	body.Ranks = ranks

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	url := fmt.Sprintf("%v/league/%v/season/%v/standing/", c.url, leagueID, year)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	}
	defer res.Body.Close()
	var addStandingReply data.AddStandingReply
	bodyReply, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bodyReply, &addStandingReply)
	if addStandingReply.Result.Code != 0 {
		return errors.New("Cannot Add")
	}
	return nil
}

// AddGame add a game
func (c *Client) AddGame(
	leagueID string,
	year int,
	home string,
	away string,
	date string,
	gameType int,
	state int,
	homeGoal int,
	awayGoal int,
) error {
	if c.sessionID == "" {
		return errors.New("Need to be logged")
	}
	body := data.AddGameRequest{}
	body.HomeID = home
	body.AwayID = away
	body.Date = date
	body.Type = gameType
	body.State = state
	body.HomeGoal = homeGoal
	body.AwayGoal = awayGoal

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)

	url := fmt.Sprintf("%v/league/%v/season/%v/game/", c.url, leagueID, year)
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	}
	defer res.Body.Close()
	var addStandingReply data.AddStandingReply
	bodyReply, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(bodyReply, &addStandingReply)
	if addStandingReply.Result.Code != 0 {
		return errors.New("Cannot Add")
	}
	return nil
}
