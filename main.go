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

	if config.Env == utils.PROD {
		http.Handle("/", http.FileServer(http.Dir("home/mathieu/topmusicstreaming/build")))
		http.Handle("/us/", http.StripPrefix("/us/", http.FileServer(http.Dir("home/mathieu/topmusicstreaming/build"))))
		http.Handle("/fr/", http.StripPrefix("/fr/", http.FileServer(http.Dir("home/mathieu/topmusicstreaming/build"))))
		http.Handle("/de/", http.StripPrefix("/de/", http.FileServer(http.Dir("home/mathieu/topmusicstreaming/build"))))
		http.Handle("/es/", http.StripPrefix("/es/", http.FileServer(http.Dir("home/mathieu/topmusicstreaming/build"))))
		http.Handle("/pt/", http.StripPrefix("/pt/", http.FileServer(http.Dir("home/mathieu/topmusicstreaming/build"))))
		http.Handle("/it/", http.StripPrefix("/it/", http.FileServer(http.Dir("home/mathieu/topmusicstreaming/build"))))
	}

	http.HandleFunc("/api", api.Api)
	err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
	if err != nil {
		utils.Logger.Panicf(err.Error())
	}
}
