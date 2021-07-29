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
	if len(d.Areas) > 3 {
		d.Title = fmt.Sprintf("Day %d, %d photos in %d areas: %s. What a busy day!", d.DayNum, len(d.Photos), len(d.Areas), strings.Join(d.Areas, ", "))
	} else if len(d.Areas) > 1 && len(d.Areas) <= 3 {
		d.Title = fmt.Sprintf("Day %d, %d photos in %d areas: %s", d.DayNum, len(d.Photos), len(d.Areas), strings.Join(d.Areas, ", "))
	} else {
		d.Title = fmt.Sprintf("Day %d, %d photos in %s", d.DayNum, len(d.Photos), d.Areas[0])
	}
}

func (h *Holiday) GetAreaInfo() {
	var updatedDays []Day
	for _, day := range h.Days {
		updatedDays = append(updatedDays, day.getAreaInfoByDay(h.Country))
	}
	h.Days = updatedDays
}

func (d *Day) getAreaInfoByDay(country string) Day {
	var aboutAreas []AboutArea
	for _, area := range d.Areas {
		queryStr := fmt.Sprintf("%s, %s", area, country)
		areaInf, err := client.GetPlaceInfo(queryStr)
		if err != nil {
			continue
		}
		about := AboutArea{
			Description: areaInf.Data[0].ResultObject.GeoDescription,
			Attractions: areaInf.Data[0].ResultObject.CategoryCounts.Attractions,
		}
		aboutAreas = append(aboutAreas, about)
	}
	d.AboutAreas = aboutAreas

	return *d
}

func (h *Holiday) SetDayPlaceData(tlocs []reader.TimeLocation) {
	tlocByDay := getPlacesByDays(tlocs)
	var days []Day
	day1 := tlocs[0].DateTime
	cities := make(map[string]bool)
	for date, locs := range tlocByDay {
		var photos []Photo
		areas := make(map[string]bool)
		for _, loc := range locs {
			locInfo, _ := getLocationData(loc)
			area := locInfo.GetArea()
			areas[area] = true

			city := locInfo.GetCity()
			cities[city] = true

			photos = append(photos, Photo{
				DateTime:  loc.DateTime,
				Lattitude: loc.Lattitude,
				Longitude: loc.Longitude,
				Location:  locInfo,
				Area:      area,
				Title:     locInfo.FormattedAdress,
			})
		}

		var areaNames []string
		for area := range areas {
			if area != "" {
				areaNames = append(areaNames, area)
			}
		}

		day := Day{
			Day:            date,
			DayNum:         getDayNum(day1, locs[0].DateTime),
			Photos:         photos,
			NumPhotosTaken: len(photos),
			Areas:          areaNames,
		}

		day.setDayTitle()
		fmt.Println(day.Title)
		days = append(days, day)
	}
	h.Days = days
	h.Country = days[0].Photos[0].Location.GetCountry()
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
	return tlocByDay
}

func getLocationData(tloc reader.TimeLocation) (client.LocationResult, error) {
	locationData, err := client.GetLocation(tloc.Lattitude, tloc.Longitude)
	return locationData, err
}
