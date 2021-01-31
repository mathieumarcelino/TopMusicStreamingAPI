package sorter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"time"
	"topmusicstreaming/utils"

	"github.com/masatana/go-textdistance"
)

type Final struct {
	Header Header  `json:"header"`
	Tracks []Track `json:"tracks"`
}

type Track struct {
	Position  int       `json:"position"`
	Evolution string    `json:"evolution"`
	Track     string    `json:"track"`
	Artist    string    `json:"artist"`
	Cover     string    `json:"cover"`
	Positions Positions `json:"positions"`
}

type TrackBeforeSort struct {
	Position          float64 `json:"position"`
	Evolution         string  `json:"evolution"`
	Track             string  `json:"track"`
	Artist            string  `json:"artist"`
	Cover             string  `json:"cover"`
	Platform1Position int     `json:"p1"`
	Platform2Position int     `json:"p2"`
	Platform3Position int     `json:"p3"`
}

type Positions struct {
	Platform1Position int     `json:"p1"`
	Platform2Position int     `json:"p2"`
	Platform3Position int     `json:"p3"`
	Average           float64 `json:"average"`
}

type Header struct {
	Country string `json:"country"`
	Date    string `json:"date"`
	Time    string `json:"time"`
	Names   Names  `json:"names"`
}

type Names struct {
	Platform1Name string `json:"n1"`
	Platform2Name string `json:"n2"`
	Platform3Name string `json:"n3"`
}

func Sorter(array1 [][]string, name1 string, array2 [][]string, name2 string, array3 [][]string, name3 string, country string) {

	jsonFile, err := os.Open("json/" + country + ".json")
	// jsonFile, err := os.Open("root/go/go-web/json/" + country + ".json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var final Final
	json.Unmarshal(byteValue, &final)

	paris, _ := time.LoadLocation("Europe/Paris")
	dt := time.Now().In(paris)

	finalNames := Names{
		Platform1Name: name1,
		Platform2Name: name2,
		Platform3Name: name3,
	}

	finalHeader := Header{
		Country: country,
		Date:    dt.Format("01-02-2006"),
		Time:    dt.Format("15:04"),
		Names:   finalNames,
	}

	finalsTracksBeforeSort := make([]TrackBeforeSort, 0)
	alreadyCkeck := []string{}

	for i := 0; i < 100; i++ {

		platform1Position := float64(i + 1)
		platform2Position := 150.0
		platform3Position := 150.0

		for j := 0; j < len(array2); j++ {
			val := textdistance.JaroWinklerDistance(strings.ToLower(array1[i][0]), strings.ToLower(array2[j][0]))
			if val == 1.0 {
				platform2Position = float64(j + 1)
			} else if val >= 0.8 && strings.ToLower(array1[i][1]) == strings.ToLower(array2[j][1]) {
				platform2Position = float64(j + 1)
				alreadyCkeck = append(alreadyCkeck, strings.ToLower(array2[j][0]))
			}
		}

		for k := 0; k < len(array3); k++ {
			val := textdistance.JaroWinklerDistance(strings.ToLower(array1[i][0]), strings.ToLower(array3[k][0]))
			if val == 1.0 {
				platform3Position = float64(k + 1)
			} else if val >= 0.8 && strings.ToLower(array1[i][1]) == strings.ToLower(array3[k][1]) {
				platform2Position = float64(k + 1)
				alreadyCkeck = append(alreadyCkeck, strings.ToLower(array3[k][0]))
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
		finalCoverUrl := array1[i][2]

		finalTrackBeforeSort := TrackBeforeSort{
			Position:          finalPositionGlobal,
			Track:             finalTrackName,
			Artist:            finalArtistName,
			Cover:             finalCoverUrl,
			Platform1Position: int(platform1Position),
			Platform2Position: int(platform2Position),
			Platform3Position: int(platform3Position),
		}

		finalsTracksBeforeSort = append(finalsTracksBeforeSort, finalTrackBeforeSort)
		alreadyCkeck = append(alreadyCkeck, strings.ToLower(array1[i][0]))
	}

	for i := 0; i < 100; i++ {

		platform2Position := float64(i + 1)
		platform1Position := 150.0
		platform3Position := 150.0

		if utils.StringInSlice(strings.ToLower(array2[i][0]), alreadyCkeck) == false {

			for j := 0; j < len(array1); j++ {
				val := textdistance.JaroWinklerDistance(strings.ToLower(array2[i][0]), strings.ToLower(array1[j][0]))
				if val == 1.0 {
					platform1Position = float64(j + 1)
				} else if val >= 0.8 && strings.ToLower(array2[i][1]) == strings.ToLower(array1[j][1]) {
					platform2Position = float64(j + 1)
					alreadyCkeck = append(alreadyCkeck, strings.ToLower(array1[j][0]))
				}
			}

			for k := 0; k < len(array3); k++ {
				val := textdistance.JaroWinklerDistance(strings.ToLower(array2[i][0]), strings.ToLower(array3[k][0]))
				if val == 1.0 {
					platform3Position = float64(k + 1)
				} else if val >= 0.8 && strings.ToLower(array2[i][1]) == strings.ToLower(array3[k][1]) {
					platform2Position = float64(k + 1)
					alreadyCkeck = append(alreadyCkeck, strings.ToLower(array3[k][0]))
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
			finalCoverUrl := array2[i][2]

			finalTrackBeforeSort := TrackBeforeSort{
				Position:          finalPositionGlobal,
				Track:             finalTrackName,
				Artist:            finalArtistName,
				Cover:             finalCoverUrl,
				Platform1Position: int(platform1Position),
				Platform2Position: int(platform2Position),
				Platform3Position: int(platform3Position),
			}

			finalsTracksBeforeSort = append(finalsTracksBeforeSort, finalTrackBeforeSort)
			alreadyCkeck = append(alreadyCkeck, strings.ToLower(array2[i][0]))

		}

	}

	for i := 0; i < 100; i++ {

		platform3Position := float64(i + 1)
		platform1Position := 150.0
		platform2Position := 150.0

		if utils.StringInSlice(strings.ToLower(array3[i][0]), alreadyCkeck) == false {

			for j := 0; j < len(array1); j++ {
				val := textdistance.JaroWinklerDistance(strings.ToLower(array3[i][0]), strings.ToLower(array1[j][0]))
				if val == 1.0 {
					platform1Position = float64(j + 1)
				} else if val >= 0.8 && strings.ToLower(array3[i][1]) == strings.ToLower(array1[j][1]) {
					platform2Position = float64(j + 1)
					alreadyCkeck = append(alreadyCkeck, strings.ToLower(array1[j][0]))
				}
			}

			for k := 0; k < len(array2); k++ {
				val := textdistance.JaroWinklerDistance(strings.ToLower(array3[i][0]), strings.ToLower(array2[k][0]))
				if val == 1.0 {
					platform2Position = float64(k + 1)
				} else if val >= 0.8 && strings.ToLower(array3[i][1]) == strings.ToLower(array2[k][1]) {
					platform2Position = float64(k + 1)
					alreadyCkeck = append(alreadyCkeck, strings.ToLower(array2[k][0]))
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
			finalCoverUrl := array3[i][2]

			finalTrackBeforeSort := TrackBeforeSort{
				Position:          finalPositionGlobal,
				Track:             finalTrackName,
				Artist:            finalArtistName,
				Cover:             finalCoverUrl,
				Platform1Position: int(platform1Position),
				Platform2Position: int(platform2Position),
				Platform3Position: int(platform3Position),
			}

			finalsTracksBeforeSort = append(finalsTracksBeforeSort, finalTrackBeforeSort)
			alreadyCkeck = append(alreadyCkeck, strings.ToLower(array3[i][0]))

		}

	}

	sort.Slice(finalsTracksBeforeSort, func(p, q int) bool {
		return finalsTracksBeforeSort[p].Position < finalsTracksBeforeSort[q].Position
	})

	finalsTracks := make([]Track, 0)

	position := 0
	lastTrackPosition := 0.0
	for i := 0; i < len(finalsTracksBeforeSort); i++ {
		if finalsTracksBeforeSort[i].Position > lastTrackPosition && position < 100 {
			position++
			finalPosition := Positions{
				Platform1Position: finalsTracksBeforeSort[i].Platform1Position,
				Platform2Position: finalsTracksBeforeSort[i].Platform2Position,
				Platform3Position: finalsTracksBeforeSort[i].Platform3Position,
				Average:           finalsTracksBeforeSort[i].Position,
			}
			finalTrack := Track{
				Position:  position,
				Evolution: CheckEvolution(final, finalsTracksBeforeSort[i].Track, position),
				Track:     finalsTracksBeforeSort[i].Track,
				Artist:    finalsTracksBeforeSort[i].Artist,
				Cover:     finalsTracksBeforeSort[i].Cover,
				Positions: finalPosition,
			}
			finalsTracks = append(finalsTracks, finalTrack)
			lastTrackPosition = finalsTracksBeforeSort[i].Position
		} else if finalsTracksBeforeSort[i].Position == lastTrackPosition {
			finalPosition := Positions{
				Platform1Position: finalsTracksBeforeSort[i].Platform1Position,
				Platform2Position: finalsTracksBeforeSort[i].Platform2Position,
				Platform3Position: finalsTracksBeforeSort[i].Platform3Position,
				Average:           finalsTracksBeforeSort[i].Position,
			}
			finalTrack := Track{
				Position:  position,
				Evolution: CheckEvolution(final, finalsTracksBeforeSort[i].Track, position),
				Track:     finalsTracksBeforeSort[i].Track,
				Artist:    finalsTracksBeforeSort[i].Artist,
				Cover:     finalsTracksBeforeSort[i].Cover,
				Positions: finalPosition,
			}
			finalsTracks = append(finalsTracks, finalTrack)
		} else {
			break
		}
	}

	finalJson := Final{
		Header: finalHeader,
		Tracks: finalsTracks,
	}

	WriteJSON(finalJson, "json/"+country+".json")
	// WriteJSON(finalJson, "root/go/go-web/json/"+country+".json")
}

func WriteJSON(data Final, file string) {
	dataFile, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Could not create JSON")
	}
	_ = ioutil.WriteFile(file, dataFile, 0666)
}

func CheckEvolution(final Final, name string, position int) (response string) {
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
