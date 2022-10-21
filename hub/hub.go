package hub

import (
	"topmusicstreaming/collector"
	"topmusicstreaming/sorter"
	"topmusicstreaming/utils"
)

func Launch(country string) {
	allInfosSpotify := collector.Spotify(country)
	allInfosAppleMusic := collector.AppleMusic(country)
	allInfosDeezer := collector.Deezer(country)

	sorter.Sort(allInfosSpotify, utils.Spotify, allInfosAppleMusic, utils.AppleMusic, allInfosDeezer, utils.Deezer, country)
}

func LaunchAll() {
	Launch(utils.US)
	Launch(utils.FR)
	Launch(utils.DE)
	Launch(utils.ES)
	Launch(utils.PT)
	Launch(utils.IT)
}
