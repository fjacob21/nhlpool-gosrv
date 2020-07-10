package config

import (
	"encoding/json"
	"io/ioutil"

	"nhlpool.com/service/go/nhlpool/data"
)

// Configs Config file
type Configs struct {
	Admin data.Player `json:"admin"`
	Port  int         `json:"port"`
}

//LoadConfigs Load config files
func LoadConfigs() *Configs {
	file, err := ioutil.ReadFile("./etc/configs.json")
	if err != nil {
		file, err = ioutil.ReadFile("configs.json")
	}
	if err != nil {
		return nil
	}
	data := Configs{}
	_ = json.Unmarshal([]byte(file), &data)
	return &data
}
