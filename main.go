package main

import (
	"flag"
	"fmt"
	"net/http"
	"team-project/configurations"
	"team-project/logger"
)


func main(){
	configFile := flag.String("config", "./project_config.json", "Configuration file in JSON-format")
	flag.Parse()

	if len(*configFile) > 0 {
		fmt.Println("Config file was read")
	}

	err := configurations.LoadConfig(*configFile)
	if err != nil {
		logger.LogFatal("Fatal error while reading config, %s", err)
	}

	logger.LogInfo("Starting HTTP listener...")
	err = http.ListenAndServe(configurations.Config.ListenURL, nil)
	if err != nil {
		logger.LogError("Error, %s",err)
	}
	logger.LogWarn("Stop running application, %s", err)
}


