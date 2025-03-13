package sorter

import (
	"encoding/json"
	"os"
	"sort"
	"strings"
	"time"
	"fmt"
	"topmusicstreaming/models"
	"topmusicstreaming/utils"
	"io/ioutil"
	"net/http"
	"net/url"

	distancetext "github.com/masatana/go-textdistance"
)

func Sort(array []models.Plateform, country string) {

	paris, _ := time.LoadLocation(utils.EuropeLocation)
	dt := time.Now().In(paris)

	header := models.Header{
		Country: country,
		Date:    dt.Format("01-02-2006"),
		Time:    dt.Format("15:04"),
	}

	var tracks []models.Track
	var alreadyCheck []string

	for _, platformCheck := range array {

		nameCheck := platformCheck.Name

		for _, infoCheck := range platformCheck.Data {
			
			track := infoCheck.Track
			artist := infoCheck.Artist
			penality := 0

			trackCheck := strings.ToLower(infoCheck.Track)
			artistCheck := strings.ToLower(infoCheck.Artist)

			if utils.StringInSlice(trackCheck, alreadyCheck) == false {

				var positions []models.Position

				positions = append(positions, models.Position{platformCheck.Name, infoCheck.Position})
				
				for _, platformToCheck := range array {		

					nameToCheck := platformToCheck.Name
					penalityToCheck := platformToCheck.Penality

					if(nameCheck != nameToCheck){

						found := false

						for _, infoToCheck := range platformToCheck.Data {

							if found == false {

								trackToCheck := strings.ToLower(infoToCheck.Track)
								artistToCheck := strings.ToLower(infoToCheck.Artist)

								if utils.StringInSlice(trackToCheck, alreadyCheck) == false {

									distanceTrack := distancetext.JaroWinklerDistance(trackToCheck, trackCheck)
									distanceArtist := distancetext.JaroWinklerDistance(artistToCheck, artistCheck)

									if distanceTrack == 1.0 && distanceArtist >= 0.8{
										found = true
										positions = append(positions, models.Position{platformToCheck.Name, infoToCheck.Position})
									} else if distanceTrack >= 0.8 && distanceArtist >= 0.8 {
										found = true
										positions = append(positions, models.Position{platformToCheck.Name, infoToCheck.Position})
										alreadyCheck = append(alreadyCheck, trackToCheck)
									}

								}

							}

						}

						if found == false {penality += penalityToCheck}
					}
				}

				alreadyCheck = append(alreadyCheck, trackCheck)
				
				track := models.Track{
					Position:	0,
					Evolution:	"",
					Track:		track,
					Artist:		artist,
					Cover:		"",
					Average:	AveragePosition(positions, penality),
					Positions:	positions,
				}

				tracks = append(tracks, track)
			}

		}

	}

	path, err := utils.BuildFilePath("json", country, "json")
	if err != nil {
		utils.Logger.Errorf("build file path: %v", err)
	}

	byteValue, _ := os.ReadFile(path)
	var pastFinal models.Final
	json.Unmarshal(byteValue, &pastFinal)

	tracks = SortTracks(tracks)
	tracks = CompareTracksEvolution(pastFinal.Tracks, tracks)
	tracks = GetCover(pastFinal.Tracks, tracks)

	final := models.Final{
		Header: header,
		Tracks: tracks,
	}

	if err = utils.WriteJSON(final, path); err != nil {
		utils.Logger.Errorf("write json: %v", err)
	}
}

func AveragePosition(positions []models.Position, penality int) float64 {
	if len(positions) == 0 {
		return 0
	}

	total := 0
	for _, pos := range positions {
		total += pos.Position
	}

	return (float64(total) / float64(len(positions))) + float64(penality)
}

func SortTracks(tracks []models.Track) []models.Track {
	sort.Slice(tracks, func(i, j int) bool {
		return tracks[i].Average < tracks[j].Average
	})

	currentPosition := 1
	for i := range tracks {
		if i > 0 && tracks[i].Average != tracks[i-1].Average {
			currentPosition = i + 1
		}
		tracks[i].Position = currentPosition
	}

	return tracks
}

func CompareTracksEvolution(pastTracks, tracks []models.Track) []models.Track {
	lastPositions := make(map[string]int)
	for _, track := range pastTracks {
		lastPositions[track.Track] = track.Position
	}

	for i := range tracks {
		track := &tracks[i]
		lastPosition, exists := lastPositions[track.Track]

		if !exists {
			track.Evolution = "N" // Nouveau morceau
		} else if track.Position < lastPosition {
			track.Evolution = "+" // Position améliorée
		} else if track.Position > lastPosition {
			track.Evolution = "-" // Position baissée
		} else {
			track.Evolution = "=" // Position inchangée
		}
	}

	return tracks
}

func GetCover(pastTracks, tracks []models.Track) []models.Track {
	if os.Getenv("GetCover") == "false" {
		return tracks
	}

	for i := range tracks {
		for _, lastTrack := range pastTracks {
			if lastTrack.Track == tracks[i].Track && lastTrack.Artist == tracks[i].Artist && lastTrack.Cover != "" {
				tracks[i].Cover = lastTrack.Cover
				break
			}
		}

		if tracks[i].Cover == "" {
			tracks[i].Cover = fetchCoverFromAPI(tracks[i].Artist, tracks[i].Track)
		}
	}

	return tracks
}

type LastFmResponse struct {
	Track struct {
		Album struct {
			Image []struct {
				Size string `json:"size"`
				URL  string `json:"#text"`
			} `json:"image"`
		} `json:"album"`
	} `json:"track"`
}

func fetchCoverFromAPI(artist, track string) string {
	baseURL := "https://ws.audioscrobbler.com/2.0/"
	params := url.Values{}
	params.Set("method", "track.getInfo")
	params.Set("api_key", os.Getenv("LastFmAPIKey"))
	params.Set("artist", artist)
	params.Set("track", track)
	params.Set("format", "json")

	requestURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(requestURL)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	var result LastFmResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return ""
	}

	for _, img := range result.Track.Album.Image {
		if img.Size == "large" {
			return img.URL
		}
	}

	return ""
}