package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"team-project/config"
	 "github.com/urfave/negroni"
        "./services"
)


func main(){

	err := config.ReadAndLoad("project_config.json")
	f, err := os.OpenFile(config.Config.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		logrus.Fatal("Error while opening log file", err)
	}
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
	if err != nil {
		logrus.Info(err)
	}
	logrus.Info("Stop running server: ", err)
}


