package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"team-project/configurations"
	"team-project/database"
	"team-project/logger"
	"team-project/services"

	"github.com/urfave/negroni"
)

func main() {
	//Flags for LoadLog and LoadConfig functions
	port := os.Getenv("PORT")
	configFile := flag.String("config", "./project_config.json", "Configuration file in JSON-format")
	logFile := flag.String("logFile", "project_log_file.log", "Logging out file .log")
	flag.Parse()

	err := configurations.LoadConfig(*configFile)
	if err != nil {
		fmt.Printf("Error while loading configurations, %s \n", err)
		return
	}
	err = logger.LoadLog(*logFile)
	if err != nil {
		fmt.Printf("Error logger not loaded, %s \n", err)
		return
	}
	err = database.PostgresInit()
	if err != nil {
		fmt.Printf("Error while connecting to postgres database, %s \n", err)
		return
	}
	defer database.Db.Close()
	err = database.RedisInit()
	if err != nil {
		fmt.Printf("Error while initializing redis client, %s \n", err)
		return
	}
	//Middleware manager
	middlewareManager := negroni.New()
	middlewareManager.Use(negroni.NewRecovery())
	middlewareManager.UseHandler(services.NewRouter())
	fmt.Println("Starting HTTP listener...")
	//Starting server
	err = http.ListenAndServe(":"+port, middlewareManager)
	if err != nil {
		logger.Logger.Errorf("Error, %s", err)
	}
	fmt.Println("Stop running application")
}
