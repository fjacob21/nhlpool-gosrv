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
	DB    string      `json:"db"`
}

//LoadConfigs Load config files
func LoadConfigs() Configs {
	data := loadDefaultConfig()
	file, err := ioutil.ReadFile("./etc/configs.json")
	if err != nil {
		file, err = ioutil.ReadFile("configs.json")
	}
	if err != nil {
		return data
	}
	_ = json.Unmarshal([]byte(file), &data)
	return data
}

func loadDefaultConfig() Configs {
	config := Configs{}
	config.Port = 8080
	config.DB = ":memory:"
	return config
}
