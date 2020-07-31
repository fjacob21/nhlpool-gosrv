package controller

import (
	"fmt"
	"log"

	"nhlpool.com/service/go/nhlpool/store"

	"nhlpool.com/service/go/nhlpool/data"
)

// GetTeams Process the get players request
func GetTeams(leagueID string) data.GetTeamsReply {
	log.Println("Get Teams")
	var reply data.GetTeamsReply
	reply.Result.Code = data.SUCCESS
	reply.Teams, _ = store.GetStore().Team().GetTeams(getLeague(leagueID))
	return reply
}

// AddTeam Process the add team request
func AddTeam(leagueID string, request data.AddTeamRequest) data.AddTeamReply {
	var reply data.AddTeamReply
	log.Println("Add Team", request)
	league := getLeague(leagueID)
	venue, _ := store.GetStore().Venue().GetVenue(request.Venue.ID, league)
	if venue == nil {
		venue = &data.Venue{ID: request.Venue.ID, League: *league, City: request.Venue.City, Name: request.Venue.Name, Timezone: request.Venue.Timezone, Address: request.Venue.Address}
		fmt.Printf("Add venue %v\n", venue)
		store.GetStore().Venue().AddVenue(venue)
	} else {
		venue.City = request.Venue.City
		venue.Name = request.Venue.Name
		venue.Timezone = request.Venue.Timezone
		venue.Address = request.Venue.Address
		store.GetStore().Venue().UpdateVenue(venue)
	}
	conference, _ := store.GetStore().Conference().GetConference(request.Conference.ID, league)
	if conference == nil {
		conference = &data.Conference{ID: request.Conference.ID, League: *league, Name: request.Conference.Name}
		fmt.Printf("Add conference %v\n", conference)
		store.GetStore().Conference().AddConference(conference)
	}
	division, _ := store.GetStore().Division().GetDivision(request.Division.ID, league)
	if division == nil {
		division = &data.Division{ID: request.Division.ID, League: *league, Name: request.Division.Name}
		fmt.Printf("Add division %v\n", division)
		store.GetStore().Division().AddDivision(division)
	}
	team := &data.Team{ID: request.ID, League: *getLeague(leagueID), Abbreviation: request.Abbreviation, Name: request.Name, Fullname: request.Fullname, City: request.City, Active: request.Active, CreationYear: request.CreationYear, Website: request.Website, Venue: venue, Conference: conference, Division: division}

	err := store.GetStore().Team().AddTeam(team)
	if err != nil {
		reply.Result.Code = data.EXISTS
		return reply
	}

	reply.Result.Code = data.SUCCESS
	reply.Team = *team
	return reply
}
