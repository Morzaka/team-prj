package logger

import (
	"team-project/configurations"
	"testing"
)

//TestLoadLogWithWrongParseLevel will test our LoadLog function with wrong ParseLevel
func TestLoadLogWithWrongParseLevel(t *testing.T) {
	LogFileName := "project_log_file.log"
	err := LoadLog(LogFileName)
	if err == nil {
		t.Errorf("Expected error, got %s", err)
	}
}

//TestLoadLog will test our LoadLog function
func TestLoadLog(t *testing.T) {
	configurations.Config = configurations.Configuration{
		LogLevel: "error",
	}
	LogFileName := "project_log_file.log"
	err := LoadLog(LogFileName)
	if err != nil {
		t.Errorf("Expected empty error, got  %s", err)
	}
}
