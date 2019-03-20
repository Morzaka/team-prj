package configurations

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"team-project/logger"
)

// Config is variable for configurations, log for display logs
var Config Configuration

// Configuration is a singleton object for application configurations
type Configuration struct {

	ListenURL   string `json:"ListenURL"`
	LogFilePath string `json:"LogFilePath"`
	RedisAddr   string `json:"localhost:6379"`
}

// Load loads configurations once
func LoadConfig(filePath string) error {
	contents, err := ioutil.ReadFile(filePath)
	if err == nil {
		reader := bytes.NewBuffer(contents)
		err = json.NewDecoder(reader).Decode(&Config)
	}
	if err != nil {
		logger.LogError("Configuration file was not read, %s ", err)
	} else {
		logger.LogInfo("Configuration file was read and loaded")
	}

	return err
}
