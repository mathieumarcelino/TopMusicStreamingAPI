package collectores

import (
	"topmusicstreaming/utils"

	"github.com/gocolly/colly"
)

func YouTube() [][]string {
	allInfosYouTube := [][]string{}
	i := 0

	collectorYouTube := colly.NewCollector(
		colly.AllowedDomains("kworb.net"),
	)

	collectorYouTube.OnHTML("#weeklytable", func(element *colly.HTMLElement) {
		element.ForEach("tbody tr", func(_ int, el *colly.HTMLElement) {
			if i < 100 {
				infoTRACKYouTube := utils.TrimStringTrack(el.ChildText(".mp div"))
				infoARTISTYouTube := utils.TrimStringArtist(el.ChildText(".mp div"))
				infoCOVERYouTube := ""

				infoYouTube := []string{infoTRACKYouTube, infoARTISTYouTube, infoCOVERYouTube}

				allInfosYouTube = append(allInfosYouTube, infoYouTube)
			}
			i++
		})

	})
	collectorYouTube.Visit("https://kworb.net/youtube/insights/es.html")

	return allInfosYouTube
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
	collectorAppleMusic.Visit("https://kworb.net/charts/apple_s/es.html")

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
	collectorDeezer.Visit("https://kworb.net/charts/deezer/es.html")

	return allInfosDeezer
}