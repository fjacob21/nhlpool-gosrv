package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// BackupPlayer Represent a player in a backup file
type BackupPlayer struct {
	Admin     bool   `json:"admin"`
	Email     string `json:"email"`
	LastLogin string `json:"last_login"`
	Name      string `json:"name"`
	Psw       string `json:"psw"`
}

// Backup Represent a backup
type Backup struct {
	Players map[string]map[string]BackupPlayer `json:"players"`
}

// LoadBackup Load a backup
func LoadBackup(path string) *Backup {
	var backup Backup
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(byteValue, &backup)
	if err != nil {
		return nil
	}
	return &backup
}
