package cmd

import (
	"fmt"

	"nhlpool.com/service/go/nhlpool/client"

	"nhlpool.com/service/go/nhlpool/data"
)

// Import players from backup
func Import(backupFile string, user string, password string) {
	backup := data.LoadBackup(backupFile)
	client := client.NewClient("localhost", 8080)
	err := client.Login(user, password)
	if err == nil {
		for pid, player := range backup.Players["1"] {
			err = client.Import(pid, player)
			if err != nil {
				fmt.Printf("Cannot Import player %v: %v err:%v\n", pid, player.Name, err)
			}
		}
		client.Logout()
	} else {
		fmt.Printf("Cannot login using credential \n")
	}
}
