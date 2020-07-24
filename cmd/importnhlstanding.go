package cmd

import (
	"fmt"

	"nhlpool.com/service/go/nhlpool/client"
)

// ImportNHLStanding players from backup
func ImportNHLStanding(year int, user string, password string) {
	nhlclient := client.NewNHLClient(year)
	client := client.NewClient("localhost", 8080)
	err := client.Login(user, password)
	if err == nil {
		standings, _ := nhlclient.GetStandings()

		for _, divison := range standings.Division {
			for _, team := range divison.TeamRecords {
				standingInfo := team.Convert()
				fmt.Printf("Add Standing %v\n", standingInfo)
				err = client.AddStanding("nhl", year, standingInfo.Team.ID, standingInfo.Points, standingInfo.Win, standingInfo.Losses, standingInfo.OT, standingInfo.GamesPlayed, standingInfo.GoalsAgainst, standingInfo.GoalsScored, standingInfo.Ranks)
				if err != nil {
					fmt.Printf("Cannot Add standing %v: %v err:%v\n", year, team.Team.ID, err)
				}
			}

		}
		client.Logout()
	} else {
		fmt.Printf("Cannot login using credential \n")
	}
}
