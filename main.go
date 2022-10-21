package main

import (
	"fmt"
	"net/http"
	"topmusicstreaming/api"
	"topmusicstreaming/cron"
	"topmusicstreaming/hub"
	"topmusicstreaming/utils"
)

func main() {

	config := utils.LoadConfig()

	utils.Logger.Infof("Running %s on %s", config.AppName, config.Env)

	if config.Env == utils.PROD {
		cron.Start()
	} else {
		hub.LaunchAll()
	}

	utils.Logger.Infof("Listening on port %d", config.Port)

	http.HandleFunc("/api", api.Api)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
	if err != nil {
		utils.Logger.Panicf(err.Error())
	}
}


