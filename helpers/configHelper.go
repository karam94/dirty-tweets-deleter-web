package helpers

import (
	"dirty-tweets-deleter-web/models"
	"encoding/json"
	"fmt"
	"os"
)

type ConfigHelper struct {}

func (helper ConfigHelper) LoadConfiguration() models.Config {
	var config models.Config

	configFile, err := os.Open("config.json")
	defer configFile.Close()

	if err != nil {
		fmt.Println(err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

