package hub

import (
	collectorde "topmusicstreaming/collector/de"
	collectores "topmusicstreaming/collector/es"
	collectorfr "topmusicstreaming/collector/fr"
	collectorit "topmusicstreaming/collector/it"
	collectorpt "topmusicstreaming/collector/pt"
	collectorus "topmusicstreaming/collector/us"
	"topmusicstreaming/sorter"
)

func Hub_FR() {
	allInfosSpotify := collectorfr.Spotify()
	allInfosAppleMusic := collectorfr.AppleMusic()
	allInfosDeezer := collectorfr.Deezer()

	sorter.Sorter(allInfosSpotify, "spotify", allInfosAppleMusic, "applemusic", allInfosDeezer, "deezer", "fr")
}

func Hub_US() {
	allInfosSpotify := collectorus.Spotify()
	allInfosAppleMusic := collectorus.AppleMusic()
	allInfosDeezer := collectorus.Deezer()

	sorter.Sorter(allInfosSpotify, "spotify", allInfosAppleMusic, "applemusic", allInfosDeezer, "deezer", "us")
}

func Hub_DE() {
	allInfosSpotify := collectorde.Spotify()
	allInfosAppleMusic := collectorde.AppleMusic()
	allInfosDeezer := collectorde.Deezer()

	sorter.Sorter(allInfosSpotify, "spotify", allInfosAppleMusic, "applemusic", allInfosDeezer, "deezer", "de")
}

func Hub_ES() {
	allInfosSpotify := collectores.Spotify()
	allInfosAppleMusic := collectores.AppleMusic()
	allInfosDeezer := collectores.Deezer()

	sorter.Sorter(allInfosSpotify, "spotify", allInfosAppleMusic, "applemusic", allInfosDeezer, "deezer", "es")
}

func Hub_PT() {
	allInfosSpotify := collectorpt.Spotify()
	allInfosAppleMusic := collectorpt.AppleMusic()
	allInfosDeezer := collectorpt.Deezer()

	sorter.Sorter(allInfosSpotify, "spotify", allInfosAppleMusic, "applemusic", allInfosDeezer, "deezer", "pt")
}

func Hub_IT() {
	allInfosSpotify := collectorit.Spotify()
	allInfosAppleMusic := collectorit.AppleMusic()
	allInfosDeezer := collectorit.Deezer()

	sorter.Sorter(allInfosSpotify, "spotify", allInfosAppleMusic, "applemusic", allInfosDeezer, "deezer", "it")
}
