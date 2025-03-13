package models

type Info struct {
	Track string
	Artist string
	Position int
}

type Plateform struct {
	Name string
	Penality int
	Data []Info
}

type Final = struct {
	Header	Header	`json:"header"`
	Tracks	[]Track	`json:"tracks"`
}

type Header struct {
	Country	string	`json:"country"`
	Date	string 	`json:"date"`
	Time	string 	`json:"time"`
}

type Track struct {
	Position	int		`json:"position"`
	Evolution	string		`json:"evolution"`
	Track		string 		`json:"track"`
	Artist		string		`json:"artist"`
	Cover		string		`json:"cover"`
	Average		float64		`json:"average"`
	Positions	[]Position	`json:"positions"`
}

type Position struct {
	Platform	string	`json:"platform"`
	Position	int		`json:"position"`
}