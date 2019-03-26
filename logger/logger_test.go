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
		t.Errorf("Testing logging file was failed, %s", err)
	}
}

//TestLoadLog will test our LoadLog function
func TestLoadLog(t *testing.T) {
	err := configurations.LoadConfig("../project_config.json")
	if err != nil {
		t.Errorf("Testing logging file was failed, LogLevel didn't read %s", err)
	}
	LogFileName := "project_log_file.log"
	err = LoadLog(LogFileName)
	if err != nil {
		t.Errorf("Testing logging file was failed, %s", err)
	}
}
