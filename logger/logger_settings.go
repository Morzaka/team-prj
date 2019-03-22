package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// Global var for LoadLog function
var Logger *logrus.Logger

//Function for opening and loading log file
func LoadLog(filePath string) error {
	LogFile, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)

	if err != nil {
		fmt.Printf("Error while opening %s ", filePath)
		return err
	}

	Logger = &logrus.Logger{
		Out:   LogFile,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}
	return nil
}
