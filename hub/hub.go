package hub

import (
	"topmusicstreaming/collector"
	"topmusicstreaming/sorter"
	"topmusicstreaming/utils"
	"topmusicstreaming/models"
)

func Launch(country string) {
	allInfosSpotify := collector.Spotify(country)
	allInfosAppleMusic := collector.AppleMusic(country)
	allInfosDeezer := collector.Deezer(country)

	allPlateform := []models.Plateform{
		{Name: utils.Spotify, Penality: utils.SpotifyPenalty, Data: allInfosSpotify},
		{Name: utils.AppleMusic, Penality: utils.AppleMusicPenalty, Data: allInfosAppleMusic},
		{Name: utils.Deezer, Penality: utils.DeezerPenalty, Data: allInfosDeezer},
	}

	sorter.Sort(allPlateform, country)
}

func LaunchAll() {
	Launch(utils.WW)
	Launch(utils.US)
	Launch(utils.FR)
	Launch(utils.DE)
	Launch(utils.ES)
	Launch(utils.PT)
	Launch(utils.IT)
}
