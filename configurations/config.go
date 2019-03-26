package configurations

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

// Config is variable for configurations
var Config Configuration

// Configuration is a singleton object for application configurations
type Configuration struct {
	ListenURL   string `json:"ListenURL"`
	LogFilePath string `json:"LogFilePath"`
	RedisAddr   string `json:"RedisAddr"`
	LogLevel    string `json:"LogLevel"`
}

// LoadConfig loads configurations once
func LoadConfig(filePath string) error {
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	err = json.NewDecoder(bytes.NewBuffer(contents)).Decode(&Config)
	if err != nil{
		return err
	}

	return err
}
