package holiday

import (
	"time"

	"github.com/SimonTanner/photo-album-maker/client"
)

type Photo struct {
	Location  client.LocationResult `json:"-"`
	DateTime  time.Time             `json:"date_time"`
	Lattitude float64
	Longitude float64
	Weather   string
	Title     string `json:"title"`
	Area      string `json:"area"`
}

type AboutArea struct {
	Description string
	Attractions struct {
		Activities  string `json:"activities"`
		Attractions string `json:"attractions"`
		Nightlife   string `json:"nightlife"`
		Shopping    string `json:"shopping"`
		Total       string `json:"total"`
	} `json:"attractions"`
}

type Day struct {
	Day            string      `json:"day"`
	DayNum         int         `json:"day_num"`
	NumPhotosTaken int         `json:"no_photos_taken"`
	Photos         []Photo     `json:"photos"`
	Title          string      `json:"title"`
	Areas          []string    `json:"areas"`
	AboutAreas     []AboutArea `json:"about_areas"`
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
