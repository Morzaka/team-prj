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
	PgHost      string `json:"PgHost"`
	PgPort      string `json:"PgPort"`
	PgUser      string `json:"PgUser"`
	PgPassword  string `json:"PgPassword"`
	PgName      string `json:"PgName"`
	JSONApi     string `json:"JSONApi"`
}

// LoadConfig loads configurations once
func LoadConfig(filePath string) error {
	contents, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	return json.NewDecoder(bytes.NewBuffer(contents)).Decode(&Config)
}
