package collector

import (
	"github.com/gocolly/colly"
	"topmusicstreaming/utils"
)

func Spotify(country string) [][]string {
	var allInfosSpotify [][]string
	i := 0

	collectorSpotify := colly.NewCollector(
		colly.AllowedDomains(utils.CollectorDomain),
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
	collectorSpotify.Visit(utils.BuildCollectorUrl(utils.Spotify, country))

	return allInfosSpotify
}

func AppleMusic(country string) [][]string {
	var allInfosAppleMusic [][]string
	i := 0

	collectorAppleMusic := colly.NewCollector(
		colly.AllowedDomains(utils.CollectorDomain),
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
	collectorAppleMusic.Visit(utils.BuildCollectorUrl(utils.AppleMusic, country))

	return allInfosAppleMusic
}

func Deezer(country string) [][]string {
	var allInfosDeezer [][]string
	i := 0

	collectorDeezer := colly.NewCollector(
		colly.AllowedDomains(utils.CollectorDomain),
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
	collectorDeezer.Visit(utils.BuildCollectorUrl(utils.Deezer, country))

	return allInfosDeezer
}
