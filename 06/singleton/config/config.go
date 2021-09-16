package config

import (
	"fmt"
	"os"

	json "github.com/goccy/go-json"
)

type config struct {
	Database struct {
		Addr     string `json:"addr"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"database"`
}

var conf *config

func GetInstance() *config {
	if conf == nil {
		configResult, err := setupConfig()
		if err != nil {
			panic(err)
		}
		return configResult
	}
	return conf
}

func setupConfig() (conf *config, err error) {
	var configFile *os.File
	if configFile, err = os.Open("./config/config.json"); err != nil {
		return nil, err
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&conf); err != nil {
		return nil, err
	}
	return conf, nil
}

// Demo

func DemoConfig() {
	config := GetInstance()
	fmt.Println(config)
}