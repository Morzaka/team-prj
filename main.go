package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"team-project/config"
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

	logrus.Info("Starting HTTP listening...")
	err = http.ListenAndServe(config.Config.ListenURL, nil)
	if err != nil {
		logrus.Info(err)
	}
	logrus.Info("Stop running server: ", err)
}


<<<<<<< HEAD
	log.Println("Shutting down")
	os.Exit(0)
}
=======
>>>>>>> 59a40cf4e7a9f76479a9ede21ae7051927a08090
