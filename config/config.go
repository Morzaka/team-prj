package config

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"

)


// Config is variable for config, log for display logs
var Config	Configuration


// Configuration is a singleton object for application config
type Configuration struct {
	ListenURL   string `json:"ListenURL"`
	LogFilePath string `json:"LogFilePath"`
}

// Load loads config once
func ReadAndLoad(FilePath string) error{
	contents, err := ioutil.ReadFile(FilePath)
	if err == nil {
		reader := bytes.NewBuffer(contents)
		err = json.NewDecoder(reader).Decode(&Config)
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"function": "ReadAndLoad()",
			"action": "Reading config from JSON file",
			"result": "File was not read",
		}).Fatal("Error while reading file")
	} else {
		logrus.WithFields(logrus.Fields{
			"function": "ReadAndLoad()",
			"action": "Reading from JSON file",
			"result": "File was read successfully!",
		}).Info("Configurations was loaded")
	}

	return err
}

