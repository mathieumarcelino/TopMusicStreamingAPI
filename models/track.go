package models

type Final struct {
	Header Header  `json:"header"`
	Tracks []Track `json:"tracks"`
}

type Track struct {
	Position  int       `json:"position"`
	Evolution string    `json:"evolution"`
	Track     string    `json:"track"`
	Artist    string    `json:"artist"`
	Cover    string    	`json:"cover"`
	Positions Positions `json:"positions"`
}

type TrackBeforeSort struct {
	Position          float64 `json:"position"`
	Evolution         string  `json:"evolution"`
	Track             string  `json:"track"`
	Artist            string  `json:"artist"`
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
