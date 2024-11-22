package main

var zoltanRegistry = map[string]Zoltan{
	"Interspározoltán": {"Interspározoltán", "/res/intersparozoltan.jpeg", []Location{
		{46.9610563, 18.9169661},
	}},
	"Iskolázoltán": {"Iskolázoltán", "/res/iskolazoltan.jpeg", []Location{
		{46.9594517, 18.9342486},
	}},
	"Tescozoltán": {"Tescozoltán", "/res/tescozoltan.jpeg", []Location{
		{46.9778892, 18.9197099},
	}},
	"Kukázoltán": {"Kukázoltán", "/res/kukazoltan.jpeg", []Location{
		{46.9583949, 18.9345873},
	}},
	"MÁVozoltán": {"MÁVozoltán", "/res/mavozoltan.jpeg", []Location{
		{46.9606042, 18.9134288},
	}},
	"VolánBuszozoltán": {"VolánBuszozoltán", "/res/volanbuszozoltan.jpeg", []Location{
		{46.9606042, 18.9134288},
	}},
	"Vonatozoltán": {"Vonatozoltán", "/res/vonatozoltan.jpeg", []Location{
		{46.9606042, 18.9134288},
	}},
	"Túrázoltán": {"Túrázoltán", "/res/turazoltan.jpeg", []Location{
		{0, 0},
	}},
}

type User struct {
	UserID   int    `json:"userID"`
	Picture  string `json:"picture"`
	Username string `json:"username"`
	Friends  []int  `json:"friends"`
	Level    int64  `json:"level"`
}

type Zoltan struct {
	Name      string     `json:"name"`
	URL       string     `json:"url"`
	Locations []Location `json:"locations"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}