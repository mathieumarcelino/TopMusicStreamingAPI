package sorter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"topmusicstreaming/utils"
)

type Infos struct {
	Tracks []Info `json:"tracks"`
}

type Info struct {
	POSITION           float64 `json:"position"`
	SpotifyPOSITION    int     `json:"spotify"`
	ApplemusicPOSITION int     `json:"applemusic"`
	DeezerPOSITION     int     `json:"deezer"`
	TRACK              string  `json:"track"`
	ARTIST             string  `json:"artist"`
}

func Sorter(array1 [][]string, array2 [][]string, array3 [][]string, country string) {

	allfinalsINFOS := make([]Info, 0)
	allfinalsINFOSJson := Infos{allfinalsINFOS}
	alreadyCkeck := []string{}

	for i := 0; i < 100; i++ {

		spotifyPOSITION := float64(i + 1)
		applemusicPOSITION := 150.0
		deezerPOSITION := 150.0

		for j := 0; j < len(array2); j++ {
			if strings.ToLower(array1[i][1]) == strings.ToLower(array2[j][1]) {
				applemusicPOSITION = float64(j + 1)
			}
		}

		for k := 0; k < len(array3); k++ {
			if strings.ToLower(array1[i][1]) == strings.ToLower(array3[k][1]) {
				deezerPOSITION = float64(k + 1)
			}
		}

		finalPOSITION := (spotifyPOSITION + applemusicPOSITION + deezerPOSITION) / 3.0
		finalTRACK := array1[i][1]
		finalARTIST := array1[i][2]

		finalInfo := Info{
			POSITION:           finalPOSITION,
			SpotifyPOSITION:    int(spotifyPOSITION),
			ApplemusicPOSITION: int(applemusicPOSITION),
			DeezerPOSITION:     int(deezerPOSITION),
			TRACK:              finalTRACK,
			ARTIST:             finalARTIST,
		}

		allfinalsINFOSJson.Tracks = append(allfinalsINFOSJson.Tracks, finalInfo)
		alreadyCkeck = append(alreadyCkeck, strings.ToLower(array1[i][1]))
		println(strings.ToLower(array1[i][1]))
	}

	for i := 0; i < 100; i++ {

		applemusicPOSITION := float64(i + 1)
		spotifyPOSITION := 150.0
		deezerPOSITION := 150.0

		if utils.StringInSlice(strings.ToLower(array2[i][1]), alreadyCkeck) == false {

			for j := 0; j < len(array1); j++ {
				if strings.ToLower(array2[i][1]) == strings.ToLower(array1[j][1]) {
					spotifyPOSITION = float64(j + 1)
				}
			}

			for k := 0; k < len(array3); k++ {
				if strings.ToLower(array2[i][1]) == strings.ToLower(array3[k][1]) {
					deezerPOSITION = float64(k + 1)
				}
			}

			finalPOSITION := (spotifyPOSITION + applemusicPOSITION + deezerPOSITION) / 3.0
			finalTRACK := array2[i][1]
			finalARTIST := array2[i][2]

			finalInfo := Info{
				POSITION:           finalPOSITION,
				SpotifyPOSITION:    int(spotifyPOSITION),
				ApplemusicPOSITION: int(applemusicPOSITION),
				DeezerPOSITION:     int(deezerPOSITION),
				TRACK:              finalTRACK,
				ARTIST:             finalARTIST,
			}

			allfinalsINFOSJson.Tracks = append(allfinalsINFOSJson.Tracks, finalInfo)
			alreadyCkeck = append(alreadyCkeck, strings.ToLower(array2[i][1]))

		}

	}

	for i := 0; i < 100; i++ {

		deezerPOSITION := float64(i + 1)
		spotifyPOSITION := 150.0
		applemusicPOSITION := 150.0

		if utils.StringInSlice(strings.ToLower(array3[i][1]), alreadyCkeck) == false {

			for j := 0; j < len(array1); j++ {
				if strings.ToLower(array3[i][1]) == strings.ToLower(array1[j][1]) {
					spotifyPOSITION = float64(j + 1)
				}
			}

			for k := 0; k < len(array2); k++ {
				if strings.ToLower(array3[i][1]) == strings.ToLower(array2[k][1]) {
					applemusicPOSITION = float64(k + 1)
				}
			}

			finalPOSITION := (spotifyPOSITION + applemusicPOSITION + deezerPOSITION) / 3.0
			finalTRACK := array3[i][1]
			finalARTIST := array3[i][2]

			finalInfo := Info{
				POSITION:           finalPOSITION,
				SpotifyPOSITION:    int(spotifyPOSITION),
				ApplemusicPOSITION: int(applemusicPOSITION),
				DeezerPOSITION:     int(deezerPOSITION),
				TRACK:              finalTRACK,
				ARTIST:             finalARTIST,
			}

			allfinalsINFOSJson.Tracks = append(allfinalsINFOSJson.Tracks, finalInfo)
			alreadyCkeck = append(alreadyCkeck, strings.ToLower(array3[i][1]))

		}

	}

	encF := json.NewEncoder(os.Stdout)
	encF.SetIndent("", " ")
	encF.Encode(allfinalsINFOSJson)

	sort.Slice(allfinalsINFOSJson.Tracks, func(p, q int) bool {
		return allfinalsINFOSJson.Tracks[p].POSITION < allfinalsINFOSJson.Tracks[q].POSITION
	})

	WriteJSON(allfinalsINFOSJson, "root/go/go-web/json/"+country+".json")
}

func WriteJSON(data Infos, file string) {
	dataFile, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Could not create JSON")
	}
	_ = ioutil.WriteFile(file, dataFile, 0666)
}
