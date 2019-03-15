package main

import (
<<<<<<< HEAD
	"flag"
	"fmt"
	"github.com/urfave/negroni"
        "gitlab.com/golang-lv-388/team-project/services"
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
=======
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"net/http"
	"os"
	"team-project/config"
	"team-project/services"
)

func main() {
	err := config.ReadAndLoad("project_config.json")
	f, err := os.OpenFile(config.Config.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
>>>>>>> routing
	if err != nil {
		logger.LogFatal("Fatal error while reading config, %s", err)
	}
<<<<<<< HEAD
	middlewareManager.Use(negroni.NewRecovery())
        middlewareManager.UseHandler(services.NewRouter())
	logger.LogInfo("Starting HTTP listener...")
	err = http.ListenAndServe(configurations.Config.ListenURL, middlewareManager)
=======
	defer func() {
		err = f.Close()
		if err != nil {
			logrus.Fatal("Error while closing log file", err)
		}
	}()

	logrus.SetOutput(f)
	// setting up web server middlewares
	middlewareManager := negroni.New()
	middlewareManager.Use(negroni.NewRecovery())
	middlewareManager.UseHandler(services.NewRouter())
	logrus.Info("Starting HTTP listening...")
	err = http.ListenAndServe(config.Config.ListenURL, middlewareManager)
>>>>>>> routing
	if err != nil {
		logger.LogError("Error, %s",err)
	}
	logger.LogWarn("Stop running application, %s", err)
}
