package collectorde

import (
	"topmusicstreaming/utils"

	"github.com/gocolly/colly"
)

func Spotify() [][]string {
	allInfosSpotify := [][]string{}
	i := 0

	collectorSpotify := colly.NewCollector(
		colly.AllowedDomains("kworb.net"),
	)

	collectorSpotify.OnHTML("#spotifydaily", func(element *colly.HTMLElement) {
		element.ForEach("tbody tr", func(_ int, el *colly.HTMLElement) {
			if i < 100 {
				infoTRACKSpotify := utils.TrimStringTrack(el.ChildText(".mp div"))
				infoARTISTSpotify := utils.TrimStringArtist(el.ChildText(".mp div"))
				infoCOVERSpotify := ""

				infoSpotify := []string{infoTRACKSpotify, infoARTISTSpotify, infoCOVERSpotify}

				allInfosSpotify = append(allInfosSpotify, infoSpotify)
			}
			i++
		})

	})
	collectorSpotify.Visit("https://kworb.net/spotify/country/de_daily.html")

	return allInfosSpotify
}

func AppleMusic() [][]string {
	allInfosAppleMusic := [][]string{}
	i := 0

	collectorAppleMusic := colly.NewCollector(
		colly.AllowedDomains("kworb.net"),
	)

	collectorAppleMusic.OnHTML(".sortable", func(element *colly.HTMLElement) {
		element.ForEach("tbody tr", func(_ int, el *colly.HTMLElement) {
			if i < 100 {
				infoTRACKAppleMusic := utils.TrimStringTrack(el.ChildText(".mp div"))
				infoARTISTAppleMusic := utils.TrimStringArtist(el.ChildText(".mp div"))
				infoCOVERAppleMusic := ""

				infoAppleMusic := []string{infoTRACKAppleMusic, infoARTISTAppleMusic, infoCOVERAppleMusic}

				allInfosAppleMusic = append(allInfosAppleMusic, infoAppleMusic)
			}
			i++
		})

	})
	collectorAppleMusic.Visit("https://kworb.net/charts/apple_s/de.html")

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
				infoTRACKDeezer := utils.TrimStringTrack(el.ChildText(".mp div"))
				infoARTISTDeezer := utils.TrimStringArtist(el.ChildText(".mp div"))
				infoCOVERDeezer := ""

				infoDeezer := []string{infoTRACKDeezer, infoARTISTDeezer, infoCOVERDeezer}

				allInfosDeezer = append(allInfosDeezer, infoDeezer)
			}
			i++
		})

	})
	collectorDeezer.Visit("https://kworb.net/charts/deezer/de.html")

	return allInfosDeezer
}
