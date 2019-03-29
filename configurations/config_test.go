package configurations

import (
	"testing"
)

//TestLoadConfig function for testing LoadConfig
func TestLoadConfig(t *testing.T) {
	FilePath := "../project_config.json"
	err := LoadConfig(FilePath)
	if err != nil {
		t.Errorf("Reading configuration failed: %s\n", err)
	}
}

//TestLoadConfig function with wrong FilePath
func TestLoadConfigWithWrongPath(t *testing.T) {
	FilePath := ""
	err := LoadConfig(FilePath)
	if err == nil {
		t.Errorf("Error while reading config from wrong path: %s\n", err)
	}
}
