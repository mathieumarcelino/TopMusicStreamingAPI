package sorter

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"time"
	"topmusicstreaming/utils"
)

type Final struct {
	Header Header  `json:"header"`
	Tracks []Track `json:"tracks"`
}

type Track struct {
	Position  float64   `json:"position"`
	Track     string    `json:"track"`
	Artist    string    `json:"artist"`
	Cover     string    `json:"cover"`
	Positions Positions `json:"positions"`
}

type Positions struct {
	Platform1Position int `json:"1"`
	Platform2Position int `json:"2"`
	Platform3Position int `json:"3"`
}

type Header struct {
	Country string `json:"country"`
	Date    string `json:"date"`
	Time    string `json:"time"`
	Names   Names  `json:"names"`
}

type Names struct {
	Platform1Name string `json:"1"`
	Platform2Name string `json:"2"`
	Platform3Name string `json:"3"`
}

func Sorter(array1 [][]string, name1 string, array2 [][]string, name2 string, array3 [][]string, name3 string, country string) {

	dt := time.Now()

	finalNames := Names{
		Platform1Name: name1,
		Platform2Name: name2,
		Platform3Name: name3,
	}

	finalHeader := Header{
		Country: country,
		Date:    dt.Format("01-02-2006"),
		Time:    dt.Format("15:04:05"),
		Names:   finalNames,
	}

	finalsTracks := make([]Track, 0)
	alreadyCkeck := []string{}

	for i := 0; i < 100; i++ {

		platform1Position := float64(i + 1)
		platform2Position := 150.0
		platform3Position := 150.0

		for j := 0; j < len(array2); j++ {
			if strings.ToLower(array1[i][1]) == strings.ToLower(array2[j][1]) {
				platform2Position = float64(j + 1)
			}
		}

		for k := 0; k < len(array3); k++ {
			if strings.ToLower(array1[i][1]) == strings.ToLower(array3[k][1]) {
				platform3Position = float64(k + 1)
			}
		}

		finalPositionGlobal := (platform1Position + platform2Position + platform3Position) / 3.0

		if platform2Position == 150.0 {
			platform2Position = 0
		}
		if platform3Position == 150.0 {
			platform3Position = 0
		}

		finalTrackName := array1[i][1]
		finalArtistName := array1[i][2]
		finalCoverUrl := array1[i][3]

		finalPosition := Positions{
			Platform1Position: int(platform1Position),
			Platform2Position: int(platform2Position),
			Platform3Position: int(platform3Position),
		}

		finalTrack := Track{
			Position:  finalPositionGlobal,
			Track:     finalTrackName,
			Artist:    finalArtistName,
			Cover:     finalCoverUrl,
			Positions: finalPosition,
		}

		finalsTracks = append(finalsTracks, finalTrack)
		alreadyCkeck = append(alreadyCkeck, strings.ToLower(array1[i][1]))
	}

	for i := 0; i < 100; i++ {

		platform2Position := float64(i + 1)
		platform1Position := 150.0
		platform3Position := 150.0

		if utils.StringInSlice(strings.ToLower(array2[i][1]), alreadyCkeck) == false {

			for j := 0; j < len(array1); j++ {
				if strings.ToLower(array2[i][1]) == strings.ToLower(array1[j][1]) {
					platform1Position = float64(j + 1)
				}
			}

			for k := 0; k < len(array3); k++ {
				if strings.ToLower(array2[i][1]) == strings.ToLower(array3[k][1]) {
					platform3Position = float64(k + 1)
				}
			}

			finalPositionGlobal := (platform1Position + platform2Position + platform3Position) / 3.0

			if platform1Position == 150.0 {
				platform1Position = 0
			}
			if platform3Position == 150.0 {
				platform3Position = 0
			}

			finalTrackName := array2[i][1]
			finalArtistName := array2[i][2]
			finalCoverUrl := array2[i][3]

			finalPosition := Positions{
				Platform1Position: int(platform1Position),
				Platform2Position: int(platform2Position),
				Platform3Position: int(platform3Position),
			}

			finalTrack := Track{
				Position:  finalPositionGlobal,
				Track:     finalTrackName,
				Artist:    finalArtistName,
				Cover:     finalCoverUrl,
				Positions: finalPosition,
			}

			finalsTracks = append(finalsTracks, finalTrack)
			alreadyCkeck = append(alreadyCkeck, strings.ToLower(array2[i][1]))

		}

	}

	for i := 0; i < 100; i++ {

		platform3Position := float64(i + 1)
		platform1Position := 150.0
		platform2Position := 150.0

		if utils.StringInSlice(strings.ToLower(array3[i][1]), alreadyCkeck) == false {

			for j := 0; j < len(array1); j++ {
				if strings.ToLower(array3[i][1]) == strings.ToLower(array1[j][1]) {
					platform1Position = float64(j + 1)
				}
			}

			for k := 0; k < len(array2); k++ {
				if strings.ToLower(array3[i][1]) == strings.ToLower(array2[k][1]) {
					platform2Position = float64(k + 1)
				}
			}

			finalPositionGlobal := (platform1Position + platform2Position + platform3Position) / 3.0

			if platform1Position == 150.0 {
				platform1Position = 0
			}
			if platform2Position == 150.0 {
				platform2Position = 0
			}

			finalTrackName := array3[i][1]
			finalArtistName := array3[i][2]
			finalCoverUrl := array3[i][3]

			finalPosition := Positions{
				Platform1Position: int(platform1Position),
				Platform2Position: int(platform2Position),
				Platform3Position: int(platform3Position),
			}

			finalTrack := Track{
				Position:  finalPositionGlobal,
				Track:     finalTrackName,
				Artist:    finalArtistName,
				Cover:     finalCoverUrl,
				Positions: finalPosition,
			}

			finalsTracks = append(finalsTracks, finalTrack)
			alreadyCkeck = append(alreadyCkeck, strings.ToLower(array3[i][1]))

		}

	}

	// encF := json.NewEncoder(os.Stdout)
	// encF.SetIndent("", " ")
	// encF.Encode(allfinalsINFOSJson)

	sort.Slice(finalsTracks, func(p, q int) bool {
		return finalsTracks[p].Position < finalsTracks[q].Position
	})

	finalJson := Final{
		Header: finalHeader,
		Tracks: finalsTracks,
	}

	WriteJSON(finalJson, "json/"+country+".json")
	// WriteJSON(allfinalsINFOSJson, "root/go/go-web/json/"+country+".json")
}

func WriteJSON(data Final, file string) {
	dataFile, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Could not create JSON")
	}
	_ = ioutil.WriteFile(file, dataFile, 0666)
}
