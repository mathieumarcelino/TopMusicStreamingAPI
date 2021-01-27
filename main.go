package main

//"topmusicstreaming/collector"

import (
	"net/http"
	"topmusicstreaming/api"
	hubfr "topmusicstreaming/hub"
)

func main() {

	hubfr.Hub_FR()

	http.HandleFunc("/api", api.Api)
	http.ListenAndServe(":9990", nil)

}
