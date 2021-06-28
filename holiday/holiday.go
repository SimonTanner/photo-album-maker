package holiday

import (
	"fmt"
	"strings"
	"time"

	"github.com/SimonTanner/photo-album-maker/client"
	"github.com/SimonTanner/photo-album-maker/reader"
)

func (h *Holiday) SetCountry(country string) {
	h.Country = country
}

func (h *Holiday) CreateTitle() {
	cities := strings.Join(h.Cities, ", ")
	h.Title = fmt.Sprintf("%d days in %s, %s", h.Duration, cities, h.Country)
}

func (h *Holiday) GetAbout() error {
	placeData, err := client.GetPlaceInfo(h.Country)
	if err != nil {
		return err
	}

	h.About = placeData.Data[0].ResultObject.GeoDescription
	return nil
}

func (d *Day) setDayTitle() {
	d.Title = fmt.Sprintf("Day %d, %d places visited %d areas: %s", d.DayNum, len(d.Places), len(d.Areas), strings.Join(d.Areas, ", "))
}

func (h *Holiday) SetDayPlaceData(tlocs []reader.TimeLocation) {
	tlocByDay := getPlacesByDays(tlocs)
	var days []Day
	day1 := tlocs[0].DateTime
	cities := make(map[string]bool)
	for date, locs := range tlocByDay {
		var places []Place
		areas := make(map[string]bool)
		for _, loc := range locs {
			locInfo, _ := getLocationData(loc)
			area := locInfo.GetArea()
			areas[area] = true

			city := locInfo.GetCity()
			cities[city] = true

			places = append(places, Place{
				DateTime:  loc.DateTime,
				Lattitude: loc.Lattitude,
				Longitude: loc.Longitude,
				Location:  locInfo,
				Area:      area,
				Title:     locInfo.FormattedAdress,
			})
		}
		// fmt.Println(fmt.Sprintf("%+v", places))

		// fmt.Println(getDayNum(day1, locs[0].DateTime))
		var areaNames []string
		for area := range areas {
			if area != "" {
				areaNames = append(areaNames, area)
			}
		}

		day := Day{
			Day:             date,
			DayNum:          getDayNum(day1, locs[0].DateTime),
			Places:          places,
			NumPlacesVisted: len(places),
			Areas:           areaNames,
		}

		day.setDayTitle()
		fmt.Println(day.Title)
		days = append(days, day)
	}
	// fmt.Println(fmt.Sprintf("%+v", days))
	h.Days = days
	h.Country = days[0].Places[0].Location.GetCountry()
	h.Cities = getCities(cities)
	h.Duration = len(days)
}

func getCities(citiesMap map[string]bool) []string {
	var cities []string
	for city := range citiesMap {
		if city != "" {
			cities = append(cities, city)
		}
	}
	return cities
}

func getDayNum(startDay, currDay time.Time) int {
	return currDay.YearDay() - startDay.YearDay() + 1
}

func getPlacesByDays(tlocs []reader.TimeLocation) map[string][]reader.TimeLocation {
	tlocByDay := make(map[string][]reader.TimeLocation)
	for _, tloc := range tlocs {
		dateStr := tloc.DateTime.Format("2006-01-02")
		if _, ok := tlocByDay[dateStr]; ok == true {
			tlocByDay[dateStr] = append(tlocByDay[dateStr], tloc)
		} else {
			tlocByDay[dateStr] = []reader.TimeLocation{tloc}
		}
	}
	// fmt.Println(tlocByDay)
	return tlocByDay
}

func getLocationData(tloc reader.TimeLocation) (client.LocationResult, error) {
	locationData, err := client.GetLocation(tloc.Lattitude, tloc.Longitude)
	return locationData, err
}
