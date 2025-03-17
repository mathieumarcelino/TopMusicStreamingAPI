package collector

import (
	"github.com/gocolly/colly"
	"topmusicstreaming/utils"
	"topmusicstreaming/models"
)

func Spotify(country string) []models.Info {
	var allInfosSpotify []models.Info
	i := 0

	collectorSpotify := colly.NewCollector(
		colly.AllowedDomains(utils.CollectorDomain),
	)

	collectorSpotify.OnHTML("#spotifydaily", func(element *colly.HTMLElement) {
		element.ForEach("tbody tr", func(_ int, el *colly.HTMLElement) {
			if i < 100 {
				infoTRACKSpotify := utils.TrimStringTrack(el.ChildText(".mp div"))
				infoARTISTSpotify := utils.TrimStringArtist(el.ChildText(".mp div"))
				infoPOSITIONSpotify := i + 1

				infoSpotify := models.Info{infoTRACKSpotify, infoARTISTSpotify, infoPOSITIONSpotify}

				allInfosSpotify = append(allInfosSpotify, infoSpotify)
			}
			i++
		})

	})
	if country == "ww" { country = "global" }
	if country == "uk" { country = "gb" }
	collectorSpotify.Visit(utils.BuildCollectorUrl(utils.Spotify, country))

	return allInfosSpotify
}

func AppleMusic(country string) []models.Info {
	var allInfosAppleMusic []models.Info
	i := 0

	collectorAppleMusic := colly.NewCollector(
		colly.AllowedDomains(utils.CollectorDomain),
	)

	collectorAppleMusic.OnHTML(".sortable", func(element *colly.HTMLElement) {
		element.ForEach("tbody tr", func(_ int, el *colly.HTMLElement) {
			if i < 100 {
				infoTRACKAppleMusic := utils.TrimStringTrack(el.ChildText(".mp div"))
				infoARTISTAppleMusic := utils.TrimStringArtist(el.ChildText(".mp div"))
				infoPOSITIONAppleMusic := i + 1

				infoAppleMusic := models.Info{infoTRACKAppleMusic, infoARTISTAppleMusic, infoPOSITIONAppleMusic}

				allInfosAppleMusic = append(allInfosAppleMusic, infoAppleMusic)
			}
			i++
		})

	})

	switch country {
		case "ww":
			collectorAppleMusic.Visit("https://kworb.net/apple_songs/index.html")
		default:
			collectorAppleMusic.Visit(utils.BuildCollectorUrl(utils.AppleMusic, country))
	}

	return allInfosAppleMusic
}

func Deezer(country string) []models.Info {
	var allInfosDeezer []models.Info
	i := 0

	collectorDeezer := colly.NewCollector(
		colly.AllowedDomains(utils.CollectorDomain),
	)

	collectorDeezer.OnHTML(".sortable", func(element *colly.HTMLElement) {
		element.ForEach("tbody tr", func(_ int, el *colly.HTMLElement) {
			if i < 100 {
				infoTRACKDeezer := utils.TrimStringTrack(el.ChildText(".mp div"))
				infoARTISTDeezer := utils.TrimStringArtist(el.ChildText(".mp div"))
				infoPOSITIONDeezer := i + 1

				infoDeezer := models.Info{infoTRACKDeezer, infoARTISTDeezer, infoPOSITIONDeezer}

				allInfosDeezer = append(allInfosDeezer, infoDeezer)
			}
			i++
		})

	})
	collectorDeezer.Visit(utils.BuildCollectorUrl(utils.Deezer, country))

	return allInfosDeezer
}
