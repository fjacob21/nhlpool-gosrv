package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	nhl "nhlpool.com/service/go/nhlpool/data/nhl"
)

// NHLClient Object to call the service
type NHLClient struct {
	year int
}

// NewNHLClient create a new client
func NewNHLClient(year int) *NHLClient {
	return &NHLClient{year: year}
}

// GetTeam Get the info about a team
func (c *NHLClient) GetTeam(id int) (*nhl.Team, error) {
	url := fmt.Sprintf("https://statsapi.web.nhl.com/api/v1/teams/%v", id)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var reply nhl.GetTeamReply
	bodyReply, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("GetTeam ReadAll Err: %v\n", err)
		return nil, err
	}
	err = json.Unmarshal(bodyReply, &reply)
	if err != nil {
		fmt.Printf("GetTeam Unmarshal Err: %v\n", err)
	}
	return &reply.Teams[0], nil
}

// GetStandings Get the season standings and results
func (c *NHLClient) GetStandings() (*nhl.Standing, error) {
	ystr := fmt.Sprintf("%v%v", c.year, c.year+1)
	url := "https://statsapi.web.nhl.com/api/v1/standings?season=" + ystr
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var reply nhl.Standing
	bodyReply, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("GetStandings ReadAll Err: %v\n", err)
		return nil, err
	}
	err = json.Unmarshal(bodyReply, &reply)
	if err != nil {
		fmt.Printf("GetStandings Unmarshal Err: %v\n", err)
	}
	return &reply, nil
}

// GetTeams Get all the teams from the season standing
func (c *NHLClient) GetTeams() (map[int]*nhl.Team, error) {
	standings, _ := c.GetStandings()
	teams := make(map[int]*nhl.Team)
	for _, record := range standings.Conferences {
		for _, team := range record.TeamRecords {
			teamInfo, _ := c.GetTeam(team.Team.ID)
			teams[team.Team.ID] = teamInfo
		}
	}
	return teams, nil
}

// GetSchedule Get the season schedule and result for a team
func (c *NHLClient) GetSchedule(team int) (*nhl.Schedule, error) {
	url := fmt.Sprintf("https://statsapi.web.nhl.com/api/v1/schedule?startDate=%v-10-01&endDate=%v-06-29&expand=schedule.teams,schedule.linescore,schedule.broadcasts,schedule.ticket,schedule.game.content.media.epg&leaderCategories=&site=en_nhlCA&teamId=%v", c.year, c.year, team)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var reply nhl.Schedule
	bodyReply, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("GetSchedule ReadAll Err: %v\n", err)
		return nil, err
	}
	err = json.Unmarshal(bodyReply, &reply)
	if err != nil {
		fmt.Printf("GetSchedule Unmarshal Err: %v\n", err)
	}
	return &reply, nil
}
