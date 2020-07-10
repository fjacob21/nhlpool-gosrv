package main

import (
	"nhlpool.com/service/go/nhlpool/cmd"

	"github.com/spf13/pflag"
)

var (
	importBackup = pflag.BoolP("import", "i", false, "Import player from backup")
	backup       = pflag.StringP("backup", "b", "backup.json", "Backup file to import")
	user         = pflag.StringP("user", "u", "admin", "User to use for action")
	password     = pflag.StringP("password", "p", "", "Password to use for action")
)

func main() {
	pflag.Parse()

	if *importBackup {
		cmd.Import(*backup, *user, *password)
	} else {
		cmd.Service()
	}
}
