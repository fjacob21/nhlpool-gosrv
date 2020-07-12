package main

import (
	"github.com/spf13/pflag"
	"nhlpool.com/service/go/nhlpool/cmd"
	"nhlpool.com/service/go/nhlpool/store"
)

var (
	importBackup = pflag.BoolP("import", "i", false, "Import player from backup")
	backup       = pflag.StringP("backup", "b", "backup.json", "Backup file to import")
	user         = pflag.StringP("user", "u", "admin", "User to use for action")
	password     = pflag.StringP("password", "p", "", "Password to use for action")
)

func main() {
	pflag.Parse()

	store.SetStore(store.NewSqliteStore())
	if *importBackup {
		cmd.Import(*backup, *user, *password)
	} else {
		cmd.Service()
	}
}
