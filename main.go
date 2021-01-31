package main

//"topmusicstreaming/collector"

import (
	"net/http"
	"topmusicstreaming/api"
	hubfr "topmusicstreaming/hub"

	"github.com/robfig/cron/v3"
)

func main() {
	// http.Handle("/", http.FileServer(http.Dir("root/go/go-web/public")))
	http.Handle("/", http.FileServer(http.Dir("/public")))

	hubfr.Hub_FR()

	http.HandleFunc("/jsonfrance", func(w http.ResponseWriter, r *http.Request) {
		hubfr.Hub_FR()
	})

	c := cron.New()
	c.AddFunc("CRON_TZ=Europe/Paris 0 16 * * *", func() { hubfr.Hub_FR() })
	c.Start()

	http.HandleFunc("/api", api.Api)
	http.ListenAndServe(":9990", nil)

}
