package api

import (
	"net/http"
	"os"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Api(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	switch r.Method {
	case http.MethodGet:
		country := r.URL.Query().Get("country")
		byteValue, _ := os.ReadFile("json/" + country + ".json")
		w.Write(byteValue)
	}
}
