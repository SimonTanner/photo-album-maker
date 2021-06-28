package reader

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type TimeLocation struct {
	DateTime  time.Time
	Lattitude float64
	Longitude float64
}

func ReadCSV(f string) ([]TimeLocation, error) {
	var timeLocations []TimeLocation
	csvFile, err := os.Open(f)
	if err != nil {
		return timeLocations, err
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	data, err := reader.ReadAll()
	if err != nil {
		return timeLocations, err
	}

	for idx, tloc := range data {
		tl, err := convertTimeLocation(tloc)
		if err != nil {
			return timeLocations, fmt.Errorf("error converting csv data at line %d: %w", idx, err)
		}
		timeLocations = append(timeLocations, tl)
	}

	sort.SliceStable(timeLocations, func(i, j int) bool {
		return timeLocations[j].DateTime.After(timeLocations[i].DateTime)
	})

	return timeLocations, nil
}

func convertTimeLocation(tloc []string) (TimeLocation, error) {
	var tl TimeLocation
	dt, tErr := parseTime(tloc[0])
	if tErr != nil {
		return tl, tErr
	}

	lat, lErr := strconv.ParseFloat(tloc[1], 64)
	if lErr != nil {
		return tl, lErr
	}

	long, loErr := strconv.ParseFloat(tloc[2], 64)
	if loErr != nil {
		return tl, loErr
	}

	tl = TimeLocation{
		DateTime:  dt,
		Lattitude: lat,
		Longitude: long,
	}

	return tl, nil
}

func parseTime(dtStr string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	if strings.Contains(dtStr, "T") && strings.Contains(dtStr, "Z") {
		layout = "2006-01-02T15:04:05Z"
	}

	return time.Parse(layout, dtStr)
}
