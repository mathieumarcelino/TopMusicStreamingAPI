package hub

import (
	"topmusicstreaming/collector"
	"topmusicstreaming/sorter"
	"topmusicstreaming/utils"
)

func Hub(country string) {
	allInfosSpotify := collector.Spotify(country)
	allInfosAppleMusic := collector.AppleMusic(country)
	allInfosDeezer := collector.Deezer(country)

	sorter.Sorter(allInfosSpotify, utils.Spotify, allInfosAppleMusic, utils.AppleMusic, allInfosDeezer, utils.Deezer, country)
}


func LaunchAll() {
	Hub(utils.US)
	Hub(utils.FR)
	Hub(utils.DE)
	Hub(utils.ES)
	Hub(utils.PT)
	Hub(utils.IT)
}
