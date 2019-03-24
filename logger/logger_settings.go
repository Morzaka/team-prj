package logger

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
	"team-project/configurations"
)

// Global var for LoadLog function
var Logger *logrus.Logger

//Function for opening and loading log file
func LoadLog(FileName string) error {
	LogFile, err := os.OpenFile(FileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return err
	}
	LogLevel, err := logrus.ParseLevel(configurations.Config.LogLevel)

	Logger = &logrus.Logger{
		Out:   LogFile,
		Level: LogLevel,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}
	return err
}
