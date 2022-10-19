package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"net/http"
	"topmusicstreaming/api"
	"topmusicstreaming/hub"
	"topmusicstreaming/utils"

)

func main() {

	config := utils.LoadConfig()

	cUS := cron.New()
	cUS.AddFunc(setCron(config.Env, 15), func() { hub.Hub_US() })
	cUS.Start()

	cFR := cron.New()
	cFR.AddFunc(setCron(config.Env, 16), func() { hub.Hub_FR() })
	cFR.Start()

	cDE := cron.New()
	cDE.AddFunc(setCron(config.Env, 17), func() { hub.Hub_DE() })
	cDE.Start()

	cES := cron.New()
	cES.AddFunc(setCron(config.Env, 18), func() { hub.Hub_ES() })
	cES.Start()

	cPT := cron.New()
	cPT.AddFunc(setCron(config.Env, 19), func() { hub.Hub_PT() })
	cPT.Start()

	cIT := cron.New()
	cIT.AddFunc(setCron(config.Env, 20), func() { hub.Hub_IT() })
	cIT.Start()

	http.HandleFunc("/api", api.Api)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Port), nil)
}

func setCron(env string, hour int) string {
	if env != "prod" {
		return ""
	}

	return fmt.Sprintf("CRON_TZ=Europe/Paris 30 %d * * *", hour)
}
