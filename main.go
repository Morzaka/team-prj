package main

import (
	"flag"
	"fmt"
	"github.com/urfave/negroni"
	"net/http"
	"team-project/configurations"
	"team-project/logger"
	"team-project/services"
)

func main() {

	//Flags for LoadLog and LoadConfig functions
	configFile := flag.String("config", "./project_fig.json", "Configuration file in JSON-format")
	logFile := flag.String("logFile", "project_log_file.log", "Logging out file .log")
	flag.Parse()
	// Invoke LoadLog functions and checking err
	err := logger.LoadLog(*logFile)
	if err != nil {
		fmt.Printf("Logger not loaded,", err)
	}
	//Invoke LoadConfig functions
	err = configurations.LoadConfig(*configFile)
	if err != nil {
		logger.Logger.Errorf("Error while reading config, %s", err)
	}
	//Middleware manager
	middlewareManager := negroni.New()
	middlewareManager.Use(negroni.NewRecovery())
	middlewareManager.UseHandler(services.NewRouter())
	logger.Logger.Infof("Starting HTTP listener...")
	//Starting server
	err = http.ListenAndServe(configurations.Config.ListenURL, middlewareManager)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	logger.Logger.Infof("Stop running application, %s", err)
}
