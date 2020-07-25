package cmd

import (
	"fmt"
	"time"

	"nhlpool.com/service/go/nhlpool/client"
)

// ImportNHLGame Import games from nhl website
func ImportNHLGame(year int, user string, password string) {
	nhlclient := client.NewNHLClient(year)
	client := client.NewClient("localhost", 8080)
	err := client.Login(user, password)
	if err == nil {
		teams, _ := nhlclient.GetTeams()
		for _, team := range teams {
			schedule, _ := nhlclient.GetSchedule(team.ID)
			for _, date := range schedule.Dates {
				game := date.Games[0]
				gameInfo := game.Convert()

				err = client.AddGame("nhl", year, gameInfo.Home.ID, gameInfo.Away.ID, gameInfo.Date.Format(time.RFC3339), gameInfo.Type, gameInfo.State, gameInfo.HomeGoal, gameInfo.AwayGoal)
				if err != nil {
					fmt.Printf("Cannot Add game year:%v home:%v: away:%v err:%v\n", year, gameInfo.Home.ID, gameInfo.Away.ID, err)
				}
			}

		}
		client.Logout()
	} else {
		fmt.Printf("Cannot login using credential \n")
	}
}
