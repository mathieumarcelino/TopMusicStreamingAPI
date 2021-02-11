package api

import (
	"fmt"
	"io/ioutil"
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
		jsonFile, err := os.Open("json/" + country + ".json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		w.Write([]byte(byteValue))
	}
}
