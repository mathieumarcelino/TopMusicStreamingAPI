package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/gocolly/colly"
)

type Info struct {
	POSITION           float64 `json:"position"`
	SpotifyPOSITION    int     `json:"spotify"`
	ApplemusicPOSITION int     `json:"applemusic"`
	DeezerPOSITION     int     `json:"deezer"`
	TRACK              string  `json:"track"`
	ARTIST             string  `json:"artist"`
}

func main() {

	// --- Spotify ---

	allInfosSpotify := [][]string{}

	collectorSpotify := colly.NewCollector(
		colly.AllowedDomains("spotifycharts.com"),
	)

	collectorSpotify.OnHTML(".chart-table tbody tr", func(element *colly.HTMLElement) {
		infoPOSITIONSpotify := element.ChildText(".chart-table-position")
		infoTRACKSpotify := trimStringTrack(element.ChildText(".chart-table-track strong"))
		infoARTISTSpotify := trimStringArtist(element.ChildText(".chart-table-track span"))[3:]

		infoSpotify := []string{infoPOSITIONSpotify, infoTRACKSpotify, infoARTISTSpotify}

		allInfosSpotify = append(allInfosSpotify, infoSpotify)
	})

	collectorSpotify.Visit("https://spotifycharts.com/regional/fr/daily/latest")

	// --- Apple Music ---

	allInfosAppleMusic := [][]string{}

	collectorAppleMusic := colly.NewCollector(
		colly.AllowedDomains("music.apple.com"),
	)

	collectorAppleMusic.OnHTML(".songs-list .track .col-song .col-song__wrapper", func(element *colly.HTMLElement) {
		infoPOSITIONAppleMusic := element.ChildText(".rank")
		infoTRACKAppleMusic := trimStringTrack(element.ChildText(".song-wrapper .song-name-wrapper .song-name"))
		infoARTISTAppleMusic := ""
		element.ForEach(".song-wrapper .song-name-wrapper .by-line span", func(_ int, el *colly.HTMLElement) {
			infoARTISTAppleMusic = el.ChildText("a:nth-child(1)")
		})

		infoAppleMusic := []string{infoPOSITIONAppleMusic, infoTRACKAppleMusic, infoARTISTAppleMusic}

		allInfosAppleMusic = append(allInfosAppleMusic, infoAppleMusic)
	})

	collectorAppleMusic.Visit("https://music.apple.com/fr/playlist/le-top-100-france/pl.6e8cfd81d51042648fa36c9df5236b8d")

	// --- Deezer ---

	allInfosDeezer := [][]string{}

	collectorDeezer := colly.NewCollector(
		colly.AllowedDomains("www.chartsmusic.fr"),
	)

	collectorDeezer.OnHTML(".table-hover", func(element *colly.HTMLElement) {
		element.ForEach("tbody tr", func(_ int, el *colly.HTMLElement) {
			infoPOSITIONDeezer := el.ChildText("td:nth-child(1) span")
			infoTRACKDeezer := trimStringTrack(el.ChildText("td:nth-child(3) strong a"))
			infoARTISTDeezer := el.ChildText("td:nth-child(4) a")

			infoDeezer := []string{infoPOSITIONDeezer, infoTRACKDeezer, infoARTISTDeezer}

			allInfosDeezer = append(allInfosDeezer, infoDeezer)
		})

	})

	collectorDeezer.Visit("https://www.chartsmusic.fr/deezer")

	// --- Final Array ---

	allfinalsINFOS := make([]Info, 0)
	alreadyCkeck := []string{}

	for i := 0; i < 100; i++ {

		spotifyPOSITION := float64(i + 1)
		applemusicPOSITION := 150.0
		deezerPOSITION := 150.0

		for j := 0; j < len(allInfosAppleMusic); j++ {
			if strings.ToLower(allInfosSpotify[i][1]) == strings.ToLower(allInfosAppleMusic[j][1]) {
				applemusicPOSITION = float64(j + 1)
			}
		}

		for k := 0; k < len(allInfosDeezer); k++ {
			if strings.ToLower(allInfosSpotify[i][1]) == strings.ToLower(allInfosDeezer[k][1]) {
				deezerPOSITION = float64(k + 1)
			}
		}

		finalPOSITION := (spotifyPOSITION + applemusicPOSITION + deezerPOSITION) / 3.0
		finalTRACK := allInfosSpotify[i][1]
		finalARTIST := allInfosSpotify[i][2]

		finalInfo := Info{
			POSITION:           finalPOSITION,
			SpotifyPOSITION:    int(spotifyPOSITION),
			ApplemusicPOSITION: int(applemusicPOSITION),
			DeezerPOSITION:     int(deezerPOSITION),
			TRACK:              finalTRACK,
			ARTIST:             finalARTIST,
		}

		allfinalsINFOS = append(allfinalsINFOS, finalInfo)
		alreadyCkeck = append(alreadyCkeck, strings.ToLower(allInfosSpotify[i][1]))
		println(strings.ToLower(allInfosSpotify[i][1]))
	}

	for i := 0; i < 100; i++ {

		applemusicPOSITION := float64(i + 1)
		spotifyPOSITION := 150.0
		deezerPOSITION := 150.0

		if stringInSlice(strings.ToLower(allInfosAppleMusic[i][1]), alreadyCkeck) == false {

			for j := 0; j < len(allInfosSpotify); j++ {
				if strings.ToLower(allInfosAppleMusic[i][1]) == strings.ToLower(allInfosSpotify[j][1]) {
					spotifyPOSITION = float64(j + 1)
				}
			}

			for k := 0; k < len(allInfosDeezer); k++ {
				if strings.ToLower(allInfosAppleMusic[i][1]) == strings.ToLower(allInfosDeezer[k][1]) {
					deezerPOSITION = float64(k + 1)
				}
			}

			finalPOSITION := (spotifyPOSITION + applemusicPOSITION + deezerPOSITION) / 3.0
			finalTRACK := allInfosAppleMusic[i][1]
			finalARTIST := allInfosAppleMusic[i][2]

			finalInfo := Info{
				POSITION:           finalPOSITION,
				SpotifyPOSITION:    int(spotifyPOSITION),
				ApplemusicPOSITION: int(applemusicPOSITION),
				DeezerPOSITION:     int(deezerPOSITION),
				TRACK:              finalTRACK,
				ARTIST:             finalARTIST,
			}

			allfinalsINFOS = append(allfinalsINFOS, finalInfo)
			alreadyCkeck = append(alreadyCkeck, strings.ToLower(allInfosAppleMusic[i][1]))

		}

	}

	for i := 0; i < 100; i++ {

		deezerPOSITION := float64(i + 1)
		spotifyPOSITION := 150.0
		applemusicPOSITION := 150.0

		if stringInSlice(strings.ToLower(allInfosDeezer[i][1]), alreadyCkeck) == false {

			for j := 0; j < len(allInfosSpotify); j++ {
				if strings.ToLower(allInfosDeezer[i][1]) == strings.ToLower(allInfosSpotify[j][1]) {
					spotifyPOSITION = float64(j + 1)
				}
			}

			for k := 0; k < len(allInfosAppleMusic); k++ {
				if strings.ToLower(allInfosDeezer[i][1]) == strings.ToLower(allInfosAppleMusic[k][1]) {
					applemusicPOSITION = float64(k + 1)
				}
			}

			finalPOSITION := (spotifyPOSITION + applemusicPOSITION + deezerPOSITION) / 3.0
			finalTRACK := allInfosDeezer[i][1]
			finalARTIST := allInfosDeezer[i][2]

			finalInfo := Info{
				POSITION:           finalPOSITION,
				SpotifyPOSITION:    int(spotifyPOSITION),
				ApplemusicPOSITION: int(applemusicPOSITION),
				DeezerPOSITION:     int(deezerPOSITION),
				TRACK:              finalTRACK,
				ARTIST:             finalARTIST,
			}

			allfinalsINFOS = append(allfinalsINFOS, finalInfo)
			alreadyCkeck = append(alreadyCkeck, strings.ToLower(allInfosDeezer[i][1]))

		}

	}

	encF := json.NewEncoder(os.Stdout)
	encF.SetIndent("", " ")
	encF.Encode(allfinalsINFOS)

	sort.Slice(allfinalsINFOS, func(p, q int) bool {
		return allfinalsINFOS[p].POSITION < allfinalsINFOS[q].POSITION
	})

	writeJSON(allfinalsINFOS, "final.json")

}

func trimStringTrack(s string) string {
	if idx := strings.Index(s, " ("); idx != -1 {
		return s[:idx]
	}
	return s
}

func trimStringArtist(s string) string {
	if idx := strings.Index(s, ","); idx != -1 {
		return s[:idx]
	}
	return s
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func writeJSON(data []Info, file string) {
	dataFile, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Could not create JSON")
	}
	_ = ioutil.WriteFile(file, dataFile, 0666)
}
