package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"
	"log"
	"os"
)

//LogError for display error massage
func LogError(massage string, args ... interface{}){
	filename := "project_log_file.log"
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		fmt.Printf("Error while opening %s ", filename)
	}

	loggerErr := &logrus.Logger{
		Out:   f,
		Level: logrus.ErrorLevel,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}
	loggerErr.Errorf(massage, args...)
	defer func() {
		err = f.Close()
		if err != nil {
			log.Fatalf("error while close log file: %s", err)
		}
	}()
}

//LogWarn for display warn massage
func LogWarn(massage string, args ... interface{}){
	filename := "project_log_file.log"
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		fmt.Printf("Error while opening %s ", filename)
	}

	loggerWarn := &logrus.Logger{
		Out:   f,
		Level: logrus.WarnLevel,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}
	loggerWarn.Warnf(massage, args...)
	defer func() {
		err = f.Close()
		if err != nil {
			log.Fatalf("error while close log file: %s", err)
		}
	}()
}

//LogInfo for display info massage
func LogInfo(massage string){
	filename := "project_log_file.log"
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		fmt.Printf("Error while opening %s ", filename)
	}

	loggerInfo := &logrus.Logger{
		Out:   f,
		Level: logrus.InfoLevel,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}
	loggerInfo.Infof(massage)
	defer func() {
		err = f.Close()
		if err != nil {
			log.Fatalf("error while close log file: %s", err)
		}
	}()
}

//LogDebug for display debug massage
func LogDebug(massage string, args ... interface{}){
	filename := "project_log_file.log"
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		fmt.Printf("Error while opening %s ", filename)
	}

	loggerDebug := &logrus.Logger{
		Out:   f,
		Level: logrus.DebugLevel,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}
	loggerDebug.Debugf(massage, args...)
	defer func() {
		err = f.Close()
		if err != nil {
			log.Fatalf("error while close log file: %s", err)
		}
	}()
}

//LogFatal for display fatal error
func LogFatal(massage string, args ... interface{}){
	filename := "project_log_file.log"
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		fmt.Printf("Error while opening %s ", filename)
	}

	loggerErr := &logrus.Logger{
		Out:   f,
		Level: logrus.FatalLevel,
		Formatter: &prefixed.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
		},
	}
	loggerErr.Fatalf(massage, args...)
	defer func() {
		err = f.Close()
		if err != nil {
			log.Fatalf("error while close log file: %s", err)
		}
	}()
}


