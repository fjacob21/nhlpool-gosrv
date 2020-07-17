package main

import (
	"github.com/spf13/pflag"
	"nhlpool.com/service/go/nhlpool/cmd"
	"nhlpool.com/service/go/nhlpool/store"
)

var (
	importBackup = pflag.BoolP("import", "i", false, "Import player from backup")
	command      = pflag.StringP("cmd", "c", "", "Command to execute")
	backup       = pflag.StringP("backup", "b", "backup.json", "Backup file to import")
	user         = pflag.StringP("user", "u", "admin", "User to use for action")
	password     = pflag.StringP("password", "p", "", "Password to use for action")
	league       = pflag.StringP("league", "l", "nhl", "League to use for action")
	year         = pflag.IntP("year", "y", 2019, "Yeat\r to use for action")
	data         = pflag.StringP("data", "d", "", "Data param for command")
)

func main() {
	pflag.Parse()

	store.SetStore(store.NewSqliteStore())
	if *importBackup {
		cmd.Import(*backup, *user, *password)
	} else if *command != "" {
		if *command == "addleague" {
			cmd.AddLeague(*data, *user, *password)
		} else if *command == "importnhlteams" {
			cmd.ImportNHLTeam(*user, *password)
		} else if *command == "addseason" {
			cmd.AddSeason(*league, *year, *user, *password)
		}

	} else {
		cmd.Service()
	}
}
