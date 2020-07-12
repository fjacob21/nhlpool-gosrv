package cmd

import (
	"encoding/json"
	"fmt"

	"nhlpool.com/service/go/nhlpool/client"
	"nhlpool.com/service/go/nhlpool/data"
)

// AddLeague Add a new league
func AddLeague(cmdData string, user string, password string) {
	client := client.NewClient("localhost", 8080)
	err := client.Login(user, password)
	if err == nil {
		var request data.AddLeagueRequest
		json.Unmarshal([]byte(cmdData), &request)
		client.AddLeague(request.ID, request.Name, request.Description, request.Website)
		client.Logout()
	} else {
		fmt.Printf("Cannot login using credential \n")
	}
}
