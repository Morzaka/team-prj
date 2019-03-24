package logger

import "testing"

func TestLoadLog(t *testing.T) {
	FileName := "project_log_file.log"
	err := LoadLog(FileName)
	if err != nil{
		t.Errorf("Opening logging file was failed, %s", err)
	}
}