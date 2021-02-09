package collectorit

import (
	"topmusicstreaming/utils"

	"github.com/gocolly/colly"
)

func Spotify() [][]string {
	allInfosSpotify := [][]string{}
	i := 0

	collectorSpotify := colly.NewCollector(
		colly.AllowedDomains("spotifycharts.com"),
	)

	collectorSpotify.OnHTML(".chart-table tbody tr", func(element *colly.HTMLElement) {
		if i < 100 {
			infoTRACKSpotify := utils.TrimStringTrack(element.ChildText(".chart-table-track strong"))
			infoARTISTSpotify := utils.TrimStringArtist(element.ChildText(".chart-table-track span"))[3:]
			infoCOVERSpotify := element.ChildAttr(".chart-table-image a img", "src")

			infoSpotify := []string{infoTRACKSpotify, infoARTISTSpotify, infoCOVERSpotify}

			allInfosSpotify = append(allInfosSpotify, infoSpotify)
		}
		i++
	})

	collectorSpotify.Visit("https://spotifycharts.com/regional/it/daily/latest")
	return allInfosSpotify
}

func AppleMusic() [][]string {
	allInfosAppleMusic := [][]string{}
	i := 0

	collectorAppleMusic := colly.NewCollector(
		colly.AllowedDomains("music.apple.com"),
	)

	collectorAppleMusic.OnHTML(".songs-list .track .col-song .col-song__wrapper", func(element *colly.HTMLElement) {
		if i < 100 {
			infoTRACKAppleMusic := utils.TrimStringTrack(element.ChildText(".song-wrapper .song-name-wrapper .song-name"))
			infoARTISTAppleMusic := ""
			infoCOVERAppleMusic := ""
			element.ForEach(".song-wrapper .song-name-wrapper .by-line span", func(_ int, el *colly.HTMLElement) {
				infoARTISTAppleMusic = el.ChildText("a:nth-child(1)")
			})

			element.ForEach(".col-song__index-wrapper .song-index .media-artwork-v2 picture", func(_ int, el *colly.HTMLElement) {
				infoCOVERAppleMusic = utils.TrimStringCoverAppleMusic(el.ChildAttr("source:nth-child(1)", "srcset"))
			})

			infoAppleMusic := []string{infoTRACKAppleMusic, infoARTISTAppleMusic, infoCOVERAppleMusic}

			allInfosAppleMusic = append(allInfosAppleMusic, infoAppleMusic)
		}
		i++
	})

	collectorAppleMusic.Visit("https://music.apple.com/fr/playlist/le-top-100-italie/pl.737e067787df485a8062e2c4927d94db")
	return allInfosAppleMusic
}

func Deezer() [][]string {
	allInfosDeezer := [][]string{}
	i := 0

	collectorDeezer := colly.NewCollector(
		colly.AllowedDomains("kworb.net"),
	)

	collectorDeezer.OnHTML(".sortable", func(element *colly.HTMLElement) {
		element.ForEach("tbody tr", func(_ int, el *colly.HTMLElement) {
			if i < 100 {
				infoTRACKDeezer := utils.TrimStringTrackDeezer(el.ChildText(".mp div"))
				infoARTISTDeezer := utils.TrimStringArtistDeezer(el.ChildText(".mp div"))
				infoCOVERDeezer := ""

				infoDeezer := []string{infoTRACKDeezer, infoARTISTDeezer, infoCOVERDeezer}

				allInfosDeezer = append(allInfosDeezer, infoDeezer)
			}
			i++
		})

	})
	collectorDeezer.Visit("https://kworb.net/charts/deezer/it.html")

	return allInfosDeezer
}
