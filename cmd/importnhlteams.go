package cmd

import (
	"fmt"
	"time"

	"nhlpool.com/service/go/nhlpool/client"
)

// ImportNHLTeam players from backup
func ImportNHLTeam(user string, password string) {
	now := time.Now()
	nhlclient := client.NewNHLClient(2019)
	client := client.NewClient("localhost", 8080)
	err := client.Login(user, password)
	if err == nil {
		teams, _ := nhlclient.GetTeams()
		fmt.Printf("Import teams %v %v\n", now.Year(), teams)
		for teamID, team := range teams {
			teamInfo := team.Convert()
			fmt.Printf("Add team %v\n", teamInfo)
			err = client.AddTeam("nhl", teamInfo.ID, teamInfo.Abbreviation, teamInfo.Name, teamInfo.Fullname, teamInfo.City, teamInfo.Active, teamInfo.CreationYear, teamInfo.Website, teamInfo.Venue.ID, teamInfo.Venue.City, teamInfo.Venue.Name, teamInfo.Venue.Timezone, teamInfo.Venue.Address, teamInfo.Conference.ID, teamInfo.Conference.Name, teamInfo.Division.ID, teamInfo.Division.Name)
			if err != nil {
				fmt.Printf("Cannot Add team %v: %v err:%v\n", teamID, teamInfo.Name, err)
			}
		}
		client.Logout()
	} else {
		fmt.Printf("Cannot login using credential \n")
	}
}
