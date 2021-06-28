package holiday

import (
	"time"

	"github.com/SimonTanner/photo-album-maker/client"
)

type Place struct {
	Location  client.LocationResult `json:"-"`
	DateTime  time.Time             `json:"date_time"`
	Lattitude float64
	Longitude float64
	Weather   string
	Title     string `json:"title"`
	Area      string `json:"area"`
}

type Day struct {
	Day             string   `json:"day"`
	DayNum          int      `json:"day_num"`
	NumPlacesVisted int      `json:"no_places_visited"`
	Places          []Place  `json:"places"`
	Title           string   `json:"title"`
	Areas           []string `json:"areas"`
}

type Holiday struct {
	Duration int      `json:"duration"`
	StartDay string   `json:"start_day"`
	EndDay   string   `json:"end_day"`
	Season   string   `json:"season"`
	Days     []Day    `json:"days"`
	Title    string   `json:"title"`
	About    string   `json:"about"`
	Cities   []string `json:"cities"`
	Country  string   `json:"country"`
}
