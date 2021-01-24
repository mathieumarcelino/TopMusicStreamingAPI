package collectorfr

import (
	"topmusicstreaming/utils"

	"github.com/gocolly/colly"
)

func Spotify() [][]string {
	allInfosSpotify := [][]string{}

	collectorSpotify := colly.NewCollector(
		colly.AllowedDomains("spotifycharts.com"),
	)

	collectorSpotify.OnHTML(".chart-table tbody tr", func(element *colly.HTMLElement) {
		infoPOSITIONSpotify := element.ChildText(".chart-table-position")
		infoTRACKSpotify := utils.TrimStringTrack(element.ChildText(".chart-table-track strong"))
		infoARTISTSpotify := utils.TrimStringArtist(element.ChildText(".chart-table-track span"))[3:]

		infoSpotify := []string{infoPOSITIONSpotify, infoTRACKSpotify, infoARTISTSpotify}

		allInfosSpotify = append(allInfosSpotify, infoSpotify)
	})

	collectorSpotify.Visit("https://spotifycharts.com/regional/fr/daily/latest")

	return allInfosSpotify
}

func AppleMusic() [][]string {
	allInfosAppleMusic := [][]string{}

	collectorAppleMusic := colly.NewCollector(
		colly.AllowedDomains("music.apple.com"),
	)

	collectorAppleMusic.OnHTML(".songs-list .track .col-song .col-song__wrapper", func(element *colly.HTMLElement) {
		infoPOSITIONAppleMusic := element.ChildText(".rank")
		infoTRACKAppleMusic := utils.TrimStringTrack(element.ChildText(".song-wrapper .song-name-wrapper .song-name"))
		infoARTISTAppleMusic := ""
		element.ForEach(".song-wrapper .song-name-wrapper .by-line span", func(_ int, el *colly.HTMLElement) {
			infoARTISTAppleMusic = el.ChildText("a:nth-child(1)")
		})

		infoAppleMusic := []string{infoPOSITIONAppleMusic, infoTRACKAppleMusic, infoARTISTAppleMusic}

		allInfosAppleMusic = append(allInfosAppleMusic, infoAppleMusic)
	})

	collectorAppleMusic.Visit("https://music.apple.com/fr/playlist/le-top-100-france/pl.6e8cfd81d51042648fa36c9df5236b8d")

	return allInfosAppleMusic
}

func Deezer() [][]string {

	allInfosDeezer := [][]string{}

	collectorDeezer := colly.NewCollector(
		colly.AllowedDomains("www.chartsmusic.fr"),
	)

	collectorDeezer.OnHTML(".table-hover", func(element *colly.HTMLElement) {
		element.ForEach("tbody tr", func(_ int, el *colly.HTMLElement) {
			infoPOSITIONDeezer := el.ChildText("td:nth-child(1) span")
			infoTRACKDeezer := utils.TrimStringTrack(el.ChildText("td:nth-child(3) strong a"))
			infoARTISTDeezer := el.ChildText("td:nth-child(4) a")

			infoDeezer := []string{infoPOSITIONDeezer, infoTRACKDeezer, infoARTISTDeezer}

			allInfosDeezer = append(allInfosDeezer, infoDeezer)
		})

	})

	collectorDeezer.Visit("https://www.chartsmusic.fr/deezer")

	return allInfosDeezer
}
