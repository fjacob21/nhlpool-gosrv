package nhl

import (
	"fmt"

	"nhlpool.com/service/go/nhlpool/data"
)

// GetTeamReply Is the struct returned by the get team request
type GetTeamReply struct {
	Teams []Team `json:"teams"`
}

// League Info about the league
type League struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

// Division Info about a division
type Division struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	NameShort    string `json:"nameShort"`
	Link         string `json:"link"`
	Abbreviation string `json:"abbreviation"`
}

// Conference Info about a conference
type Conference struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

// Franchise Info about a franchise
type Franchise struct {
	ID   int    `json:"franchiseId"`
	Name string `json:"teamName"`
	Link string `json:"link"`
}

// Team Info about a team
type Team struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	Link            string     `json:"link"`
	Venue           Venue      `json:"venue"`
	Abbreviation    string     `json:"abbreviation"`
	TeamName        string     `json:"teamName"`
	LocationName    string     `json:"locationName"`
	FirstYearOfPlay string     `json:"firstYearOfPlay"`
	Division        Division   `json:"division"`
	Conference      Conference `json:"conference"`
	Franchise       Franchise  `json:"franchise"`
	ShortName       string     `json:"shortName"`
	OfficialSiteURL string     `json:"officialSiteUrl"`
	FranchiseID     int        `json:"franchiseId"`
	Active          bool       `json:"active"`
}

// Convert Convert to data team
func (t *Team) Convert() *data.Team {
	team := &data.Team{}
	team.ID = fmt.Sprintf("%d", t.ID)
	team.Abbreviation = t.Abbreviation
	team.Name = t.TeamName
	team.Fullname = t.Name
	if t.Venue.City != nil {
		team.City = *t.Venue.City
	}
	team.Active = t.Active
	team.CreationYear = t.FirstYearOfPlay
	team.Website = t.OfficialSiteURL
	team.Venue = t.Venue.Convert()
	team.Conference = t.Conference.Convert()
	team.Division = t.Division.Convert()
	return team
}

// Convert Convert to data Conference
func (c *Conference) Convert() *data.Conference {
	conference := &data.Conference{}
	conference.ID = fmt.Sprintf("%d", c.ID)
	conference.Name = c.Name
	return conference
}

// Convert Convert to data division
func (c *Division) Convert() *data.Division {
	division := &data.Division{}
	division.ID = fmt.Sprintf("%d", c.ID)
	division.Name = c.Name
	return division
}
