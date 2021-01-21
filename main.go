package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Info struct {
	POSITION int    `json:"position"`
	TRACK    string `json:"track"`
	ARTIST   string `json:"artist"`
}

func main() {

	// --- Spotify ---

	allInfosSpotify := make([]Info, 0)

	collectorSpotify := colly.NewCollector(
		colly.AllowedDomains("spotifycharts.com"),
	)

	collectorSpotify.OnHTML(".chart-table tbody tr", func(element *colly.HTMLElement) {
		infoPOSITIONSpotify, err := strconv.Atoi(element.ChildText(".chart-table-position"))
		if err == nil {
			fmt.Println("Could not connvert into int")
		}
		infoTRACKSpotify := trimStringTrack(element.ChildText(".chart-table-track strong"))
		infoARTISTSpotify := trimStringArtist(element.ChildText(".chart-table-track span"))[3:]

		infoSpotify := Info{
			POSITION: infoPOSITIONSpotify,
			TRACK:    infoTRACKSpotify,
			ARTIST:   infoARTISTSpotify,
		}

		allInfosSpotify = append(allInfosSpotify, infoSpotify)
	})

	collectorSpotify.Visit("https://spotifycharts.com/regional/fr/daily/latest")

	encS := json.NewEncoder(os.Stdout)
	encS.SetIndent("", " ")
	encS.Encode(allInfosSpotify)

	writeJSON(allInfosSpotify, "spotify.json")

	// --- Apple Music ---

	allInfosAppleMusic := make([]Info, 0)

	collectorAppleMusic := colly.NewCollector(
		colly.AllowedDomains("music.apple.com"),
	)

	collectorAppleMusic.OnHTML(".songs-list .track .col-song .col-song__wrapper", func(element *colly.HTMLElement) {
		infoPOSITIONAppleMusic, err := strconv.Atoi(element.ChildText(".rank"))
		if err == nil {
			fmt.Println("Could not connvert into int")
		}
		infoTRACKAppleMusic := trimStringTrack(element.ChildText(".song-wrapper .song-name-wrapper .song-name"))
		infoARTISTAppleMusic := element.ChildText(".song-wrapper .song-name-wrapper .by-line span a")

		infoAppleMusic := Info{
			POSITION: infoPOSITIONAppleMusic,
			TRACK:    infoTRACKAppleMusic,
			ARTIST:   infoARTISTAppleMusic,
		}

		allInfosAppleMusic = append(allInfosAppleMusic, infoAppleMusic)
	})

	collectorAppleMusic.Visit("https://music.apple.com/fr/playlist/le-top-100-france/pl.6e8cfd81d51042648fa36c9df5236b8d")

	encAM := json.NewEncoder(os.Stdout)
	encAM.SetIndent("", " ")
	encAM.Encode(allInfosAppleMusic)

	writeJSON(allInfosAppleMusic, "applemusic.json")

	// --- Deezer ---

	allInfosDeezer := make([]Info, 0)

	collectorDeezer := colly.NewCollector(
		colly.AllowedDomains("www.chartsmusic.fr"),
	)

	collectorDeezer.OnHTML(".table-hover", func(element *colly.HTMLElement) {

		element.ForEach("tbody tr", func(_ int, el *colly.HTMLElement) {
			infoPOSITIONDeezer, err := strconv.Atoi(el.ChildText("td:nth-child(1) span"))
			if err == nil {
				fmt.Println("Could not connvert into int")
			}

			infoTRACKDeezer := trimStringTrack(el.ChildText("td:nth-child(3) strong a"))
			infoARTISTDeezer := el.ChildText("td:nth-child(4) a")

			infoDeezer := Info{
				POSITION: infoPOSITIONDeezer,
				TRACK:    infoTRACKDeezer,
				ARTIST:   infoARTISTDeezer,
			}

			allInfosDeezer = append(allInfosDeezer, infoDeezer)
		})

	})

	collectorDeezer.Visit("https://www.chartsmusic.fr/deezer")

	encD := json.NewEncoder(os.Stdout)
	encD.SetIndent("", " ")
	encD.Encode(allInfosDeezer)

	writeJSON(allInfosDeezer, "deezer.json")

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

func writeJSON(data []Info, file string) {
	dataFile, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Could not create JSON")
	}
	_ = ioutil.WriteFile(file, dataFile, 0666)
}
