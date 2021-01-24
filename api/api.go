package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Api(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		country := r.URL.Query().Get("country")
		jsonFile, err := os.Open("root/go/go-web/json/" + country + ".json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		w.Write([]byte(byteValue))
	}
}
