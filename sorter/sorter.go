package sorter

import (
	"encoding/json"
	"os"
	"sort"
	"strings"
	"time"
	"fmt"
	"topmusicstreaming/bot"
	"topmusicstreaming/models"
	"topmusicstreaming/utils"
	"unicode/utf8" //nolint
	"io/ioutil"
	"net/http"
	"net/url"

	distancetext "github.com/masatana/go-textdistance"
)

func Sort(array1 [][]string, name1 string, array2 [][]string, name2 string, array3 [][]string, name3 string, country string) {

	path, err := utils.BuildFilePath("json", country, "json")
	if err != nil {
		utils.Logger.Errorf("build file path: %v", err)
	}

	byteValue, _ := os.ReadFile(path)
	var final models.Final
	json.Unmarshal(byteValue, &final)

	paris, _ := time.LoadLocation(utils.EuropeLocation)
	dt := time.Now().In(paris)

	finalNames := models.Names{
		Platform1Name: name1,
		Platform2Name: name2,
		Platform3Name: name3,
	}

	finalHeader := models.Header{
		Country: country,
		Date:    dt.Format("01-02-2006"),
		Time:    dt.Format("15:04"),
		Names:   finalNames,
	}

	finalsTracksBeforeSort := make([]models.TrackBeforeSort, 0)
	var alreadyCheck []string

	for i := 0; i < 100; i++ {

		platform1Position := float64(i + 1)
		platform2Position := 150.0
		platform3Position := 150.0

		for j := 0; j < len(array2); j++ {
			val := distancetext.JaroWinklerDistance(strings.ToLower(array1[i][0]), strings.ToLower(array2[j][0]))
			if val == 1.0 {
				platform2Position = float64(j + 1)
			} else if val >= 0.8 && strings.ToLower(array1[i][1]) == strings.ToLower(array2[j][1]) {
				platform2Position = float64(j + 1)
				alreadyCheck = append(alreadyCheck, strings.ToLower(array2[j][0]))
			}
		}

		for k := 0; k < len(array3); k++ {
			val := distancetext.JaroWinklerDistance(strings.ToLower(array1[i][0]), strings.ToLower(array3[k][0]))
			if val == 1.0 {
				platform3Position = float64(k + 1)
			} else if val >= 0.8 && strings.ToLower(array1[i][1]) == strings.ToLower(array3[k][1]) {
				platform2Position = float64(k + 1)
				alreadyCheck = append(alreadyCheck, strings.ToLower(array3[k][0]))
			}
		}

		finalPositionGlobal := (platform1Position + platform2Position + platform3Position) / 3.0

		if platform2Position == 150.0 {
			platform2Position = 0
		}
		if platform3Position == 150.0 {
			platform3Position = 0
		}

		finalTrackName := array1[i][0]
		finalArtistName := array1[i][1]

		finalTrackBeforeSort := models.TrackBeforeSort{
			Position:          finalPositionGlobal,
			Track:             finalTrackName,
			Artist:            finalArtistName,
			Platform1Position: int(platform1Position),
			Platform2Position: int(platform2Position),
			Platform3Position: int(platform3Position),
		}

		finalsTracksBeforeSort = append(finalsTracksBeforeSort, finalTrackBeforeSort)
		alreadyCheck = append(alreadyCheck, strings.ToLower(array1[i][0]))
	}

	for i := 0; i < 100; i++ {

		platform2Position := float64(i + 1)
		platform1Position := 150.0
		platform3Position := 150.0

		if utils.StringInSlice(strings.ToLower(array2[i][0]), alreadyCheck) == false {

			for j := 0; j < len(array1); j++ {
				val := distancetext.JaroWinklerDistance(strings.ToLower(array2[i][0]), strings.ToLower(array1[j][0]))
				if val == 1.0 {
					platform1Position = float64(j + 1)
				} else if val >= 0.8 && strings.ToLower(array2[i][1]) == strings.ToLower(array1[j][1]) {
					platform2Position = float64(j + 1)
					alreadyCheck = append(alreadyCheck, strings.ToLower(array1[j][0]))
				}
			}

			for k := 0; k < len(array3); k++ {
				val := distancetext.JaroWinklerDistance(strings.ToLower(array2[i][0]), strings.ToLower(array3[k][0]))
				if val == 1.0 {
					platform3Position = float64(k + 1)
				} else if val >= 0.8 && strings.ToLower(array2[i][1]) == strings.ToLower(array3[k][1]) {
					platform2Position = float64(k + 1)
					alreadyCheck = append(alreadyCheck, strings.ToLower(array3[k][0]))
				}
			}

			finalPositionGlobal := (platform1Position + platform2Position + platform3Position) / 3.0

			if platform1Position == 150.0 {
				platform1Position = 0
			}
			if platform3Position == 150.0 {
				platform3Position = 0
			}

			finalTrackName := array2[i][0]
			finalArtistName := array2[i][1]

			finalTrackBeforeSort := models.TrackBeforeSort{
				Position:          finalPositionGlobal,
				Track:             finalTrackName,
				Artist:            finalArtistName,
				Platform1Position: int(platform1Position),
				Platform2Position: int(platform2Position),
				Platform3Position: int(platform3Position),
			}

			finalsTracksBeforeSort = append(finalsTracksBeforeSort, finalTrackBeforeSort)
			alreadyCheck = append(alreadyCheck, strings.ToLower(array2[i][0]))

		}

	}

	for i := 0; i < 100; i++ {

		platform3Position := float64(i + 1)
		platform1Position := 150.0
		platform2Position := 150.0

		if utils.StringInSlice(strings.ToLower(array3[i][0]), alreadyCheck) == false {

			for j := 0; j < len(array1); j++ {
				val := distancetext.JaroWinklerDistance(strings.ToLower(array3[i][0]), strings.ToLower(array1[j][0]))
				if val == 1.0 {
					platform1Position = float64(j + 1)
				} else if val >= 0.8 && strings.ToLower(array3[i][1]) == strings.ToLower(array1[j][1]) {
					platform2Position = float64(j + 1)
					alreadyCheck = append(alreadyCheck, strings.ToLower(array1[j][0]))
				}
			}

			for k := 0; k < len(array2); k++ {
				val := distancetext.JaroWinklerDistance(strings.ToLower(array3[i][0]), strings.ToLower(array2[k][0]))
				if val == 1.0 {
					platform2Position = float64(k + 1)
				} else if val >= 0.8 && strings.ToLower(array3[i][1]) == strings.ToLower(array2[k][1]) {
					platform2Position = float64(k + 1)
					alreadyCheck = append(alreadyCheck, strings.ToLower(array2[k][0]))
				}
			}

			finalPositionGlobal := (platform1Position + platform2Position + platform3Position) / 3.0

			if platform1Position == 150.0 {
				platform1Position = 0
			}
			if platform2Position == 150.0 {
				platform2Position = 0
			}

			finalTrackName := array3[i][0]
			finalArtistName := array3[i][1]

			finalTrackBeforeSort := models.TrackBeforeSort{
				Position:          finalPositionGlobal,
				Track:             finalTrackName,
				Artist:            finalArtistName,
				Platform1Position: int(platform1Position),
				Platform2Position: int(platform2Position),
				Platform3Position: int(platform3Position),
			}

			finalsTracksBeforeSort = append(finalsTracksBeforeSort, finalTrackBeforeSort)
			alreadyCheck = append(alreadyCheck, strings.ToLower(array3[i][0]))

		}

	}

	sort.Slice(finalsTracksBeforeSort, func(p, q int) bool {
		return finalsTracksBeforeSort[p].Position < finalsTracksBeforeSort[q].Position
	})

	finalsTracks := make([]models.Track, 0)

	position := 0
	lastTrackPosition := 0.0
	for i := 0; i < len(finalsTracksBeforeSort); i++ {
		if finalsTracksBeforeSort[i].Position > lastTrackPosition && position < 100 {
			position++
			finalPosition := models.Positions{
				Platform1Position: finalsTracksBeforeSort[i].Platform1Position,
				Platform2Position: finalsTracksBeforeSort[i].Platform2Position,
				Platform3Position: finalsTracksBeforeSort[i].Platform3Position,
				Average:           finalsTracksBeforeSort[i].Position,
			}
			finalTrack := models.Track{
				Position:  position,
				Evolution: checkEvolution(final, finalsTracksBeforeSort[i].Track, position),
				Track:     finalsTracksBeforeSort[i].Track,
				Artist:    finalsTracksBeforeSort[i].Artist,
				Cover:	   getCover(finalsTracksBeforeSort[i].Artist, finalsTracksBeforeSort[i].Track),
				Positions: finalPosition,
			}
			finalsTracks = append(finalsTracks, finalTrack)
			lastTrackPosition = finalsTracksBeforeSort[i].Position
		} else if finalsTracksBeforeSort[i].Position == lastTrackPosition {
			finalPosition := models.Positions{
				Platform1Position: finalsTracksBeforeSort[i].Platform1Position,
				Platform2Position: finalsTracksBeforeSort[i].Platform2Position,
				Platform3Position: finalsTracksBeforeSort[i].Platform3Position,
				Average:           finalsTracksBeforeSort[i].Position,
			}
			finalTrack := models.Track{
				Position:  position,
				Evolution: checkEvolution(final, finalsTracksBeforeSort[i].Track, position),
				Track:     finalsTracksBeforeSort[i].Track,
				Artist:    finalsTracksBeforeSort[i].Artist,
				Cover:	   getCover(finalsTracksBeforeSort[i].Artist, finalsTracksBeforeSort[i].Track),
				Positions: finalPosition,
			}
			finalsTracks = append(finalsTracks, finalTrack)
		} else {
			break
		}
	}

	if utils.LoadConfig().Env == utils.PROD && os.Getenv("TwitterBotEnabled") == "true"{
		tweet := bot.TweetHeader(country)
		tweet += "\n\n"
		tweet += bot.TweetPosition(finalsTracks[0].Position) + bot.TweetEvolution(finalsTracks[0].Evolution) + " " + finalsTracks[0].Artist + " - " + finalsTracks[0].Track + "\n"
		tweet += bot.TweetPosition(finalsTracks[1].Position) + bot.TweetEvolution(finalsTracks[1].Evolution) + " " + finalsTracks[1].Artist + " - " + finalsTracks[1].Track + "\n"
		tweet += bot.TweetPosition(finalsTracks[2].Position) + bot.TweetEvolution(finalsTracks[2].Evolution) + " " + finalsTracks[2].Artist + " - " + finalsTracks[2].Track + "\n"
		tweet += bot.TweetPosition(finalsTracks[3].Position) + bot.TweetEvolution(finalsTracks[3].Evolution) + " " + finalsTracks[3].Artist + " - " + finalsTracks[3].Track + "\n"
		tweet += bot.TweetPosition(finalsTracks[4].Position) + bot.TweetEvolution(finalsTracks[4].Evolution) + " " + finalsTracks[4].Artist + " - " + finalsTracks[4].Track + "\n"
		tweet += "\n"
		tweet += utils.AppDomainBaseUri + "/" + country
		tweet += "\n\n"
		tweet += bot.TweetHashtag(finalsTracks[0].Artist, finalsTracks[1].Artist, finalsTracks[2].Artist, finalsTracks[3].Artist, finalsTracks[4].Artist, utf8.RuneCountInString(tweet))
		bot.Tweet(tweet)
	}

	finalJson := models.Final{
		Header: finalHeader,
		Tracks: finalsTracks,
	}

	if err = utils.WriteJSON(finalJson, path); err != nil {
		utils.Logger.Errorf("write json: %v", err)
	}
}

func checkEvolution(final models.Final, name string, position int) (response string) {
	evolution := "N"
	for i := 0; i < len(final.Tracks); i++ {
		if strings.ToLower(name) == strings.ToLower(final.Tracks[i].Track) {
			if position > final.Tracks[i].Position {
				evolution = "-"
				break
			} else if position < final.Tracks[i].Position {
				evolution = "+"
				break
			} else if position == final.Tracks[i].Position {
				evolution = "="
				break
			}
		}
	}
	return evolution
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

func getCover(artist string, track string) (response string) {

	if os.Getenv("GetCover") == "false" {
		return ""
	}

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

