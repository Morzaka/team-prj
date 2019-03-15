package main

import (
	"flag"
	"fmt"
	"github.com/urfave/negroni"
        "team-project/services"
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
	middlewareManager.Use(negroni.NewRecovery())
        middlewareManager.UseHandler(services.NewRouter())
	logger.LogInfo("Starting HTTP listener...")
	err = http.ListenAndServe(configurations.Config.ListenURL, middlewareManager)
	if err != nil {
		logger.LogError("Error, %s",err)
	}
	logger.LogWarn("Stop running application, %s", err)
}
