package cmd

import (
	"fmt"

	"nhlpool.com/service/go/nhlpool/client"
)

// AddSeason Add a new league
func AddSeason(leagueID string, year int, user string, password string) {
	client := client.NewClient("localhost", 8080)
	err := client.Login(user, password)
	if err == nil {
		client.AddSeason(leagueID, year)
		client.Logout()
	} else {
		fmt.Printf("Cannot login using credential \n")
	}
}
