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
	allInfosYouTube := collectorfr.YouTube()
	allInfosAppleMusic := collectorfr.AppleMusic()
	allInfosDeezer := collectorfr.Deezer()

	sorter.Sorter(allInfosYouTube, "youtube", allInfosAppleMusic, "applemusic", allInfosDeezer, "deezer", "fr")
}

func Hub_US() {
	allInfosYouTube := collectorus.YouTube()
	allInfosAppleMusic := collectorus.AppleMusic()
	allInfosDeezer := collectorus.Deezer()

	sorter.Sorter(allInfosYouTube, "youtube", allInfosAppleMusic, "applemusic", allInfosDeezer, "deezer", "us")
}

func Hub_DE() {
	allInfosYouTube := collectorde.YouTube()
	allInfosAppleMusic := collectorde.AppleMusic()
	allInfosDeezer := collectorde.Deezer()

	sorter.Sorter(allInfosYouTube, "youtube", allInfosAppleMusic, "applemusic", allInfosDeezer, "deezer", "de")
}

func Hub_ES() {
	allInfosYouTube := collectores.YouTube()
	allInfosAppleMusic := collectores.AppleMusic()
	allInfosDeezer := collectores.Deezer()

	sorter.Sorter(allInfosYouTube, "youtube", allInfosAppleMusic, "applemusic", allInfosDeezer, "deezer", "es")
}

func Hub_PT() {
	allInfosYouTube := collectorpt.YouTube()
	allInfosAppleMusic := collectorpt.AppleMusic()
	allInfosDeezer := collectorpt.Deezer()

	sorter.Sorter(allInfosYouTube, "youtube", allInfosAppleMusic, "applemusic", allInfosDeezer, "deezer", "pt")
}

func Hub_IT() {
	allInfosYouTube := collectorit.YouTube()
	allInfosAppleMusic := collectorit.AppleMusic()
	allInfosDeezer := collectorit.Deezer()

	sorter.Sorter(allInfosYouTube, "youtube", allInfosAppleMusic, "applemusic", allInfosDeezer, "deezer", "it")
}
