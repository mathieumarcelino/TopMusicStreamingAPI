package main

//"topmusicstreaming/collector"

import (
	"net/http"
	"topmusicstreaming/api"
	"topmusicstreaming/hub"

	"github.com/robfig/cron/v3"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("/public")))
	http.Handle("/ww/", http.StripPrefix("/ww/", http.FileServer(http.Dir("/public"))))
	http.Handle("/us/", http.StripPrefix("/us/", http.FileServer(http.Dir("/public"))))
	http.Handle("/fr/", http.StripPrefix("/fr/", http.FileServer(http.Dir("/public"))))
	http.Handle("/de/", http.StripPrefix("/de/", http.FileServer(http.Dir("/public"))))
	http.Handle("/es/", http.StripPrefix("/es/", http.FileServer(http.Dir("/public"))))
	http.Handle("/pt/", http.StripPrefix("/pt/", http.FileServer(http.Dir("/public"))))
	http.Handle("/it/", http.StripPrefix("/it/", http.FileServer(http.Dir("/public"))))

	cWW := cron.New()
	cWW.AddFunc("CRON_TZ=Europe/Paris 30 14 * * *", func() { hub.Hub_WW() })
	cWW.Start()

	cUS := cron.New()
	cUS.AddFunc("CRON_TZ=Europe/Paris 30 15 * * *", func() { hub.Hub_US() })
	cUS.Start()

	cFR := cron.New()
	cFR.AddFunc("CRON_TZ=Europe/Paris 30 16 * * *", func() { hub.Hub_FR() })
	cFR.Start()

	cDE := cron.New()
	cDE.AddFunc("CRON_TZ=Europe/Paris 30 17 * * *", func() { hub.Hub_DE() })
	cDE.Start()

	cES := cron.New()
	cES.AddFunc("CRON_TZ=Europe/Paris 30 18 * * *", func() { hub.Hub_ES() })
	cES.Start()

	cPT := cron.New()
	cPT.AddFunc("CRON_TZ=Europe/Paris 30 19 * * *", func() { hub.Hub_PT() })
	cPT.Start()

	cIT := cron.New()
	cIT.AddFunc("CRON_TZ=Europe/Paris 30 20 * * *", func() { hub.Hub_IT() })
	cIT.Start()

	http.HandleFunc("/api", api.Api)
	http.ListenAndServe(":9990", nil)

}
