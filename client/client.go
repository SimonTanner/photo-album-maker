package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	apiKey  = ""
	mapsApi = "https://maps.googleapis.com/maps/api/geocode/json?"

	placeApi = `https://maps.googleapis.com/maps/api/place/findplacefromtext/json?input=%s&inputtype=textquery&fields=name,opening_hours,rating&key=%s`

	travelAdvisorApi = `https://travel-advisor.p.rapidapi.com/locations/search?query=%s&limit=1&offset=0&units=km&currency=USD&sort=relevance&lang=en_US`
	rapidApiKey      = ""

	weatherApiDecoded = `https://visual-crossing-weather.p.rapidapi.com/history?startDateTime=%sT08:00:00&aggregateHours=12&location=%s&endDateTime=%sT00:00:00&unitGroup=metric&contentType=json&shortColumnNames=0`
)

type AddressComponent struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

type LocationResult struct {
	AddressComponents []AddressComponent `json:"address_components"`
	FormattedAdress   string             `json:"formatted_address"`
}

func (l *LocationResult) GetCountry() string {
	return l.getLocationByKey("country")
}

func (l *LocationResult) GetCity() string {
	keys := []string{
		"administrative_area_level_1",
		"administrative_area_level_2",
	}

	return l.getValByKeys(keys)
}

func (l *LocationResult) GetArea() string {
	keys := []string{
		"sublocality_level_1",
		"locality",
		"administrative_area_level_3",
	}

	return l.getValByKeys(keys)
}

func (l *LocationResult) getValByKeys(keys []string) string {
	var val string
	for _, key := range keys {
		val = l.getLocationByKey(key)
		if val != "" {
			return val
		}
	}
	return val
}

func (l *LocationResult) getLocationByKey(key string) string {
	for _, add := range l.AddressComponents {
		for _, t := range add.Types {
			if t == key {
				return add.LongName
			}
		}
	}
	return ""
}

// GetLocation gets the location data from an googles geocoding api
func GetLocation(lat, long float64) (LocationResult, error) {
	var locRes LocationResult
	client := http.Client{}
	latLongStr := fmt.Sprintf("latlng=%v,%v", lat, long)
	URL := fmt.Sprintf("%s%s&key=%s", mapsApi, latLongStr, apiKey)

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return locRes, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return locRes, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return locRes, err
	}
	defer resp.Body.Close()

	var jsonData map[string]interface{}
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		return locRes, err
	}

	var results []LocationResult
	r, _ := json.Marshal(jsonData["results"])
	err = json.Unmarshal(r, &results)
	if err != nil {
		return locRes, err
	}
	if len(results) != 0 {
		return results[0], nil
	} else {
		return locRes, errors.New("no addresses returned")
	}
}

// GetPlace gets extra data from the google places api
func GetPlace(address string) error {
	URL := fmt.Sprintf(placeApi, url.QueryEscape(address), apiKey)
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println(string(body))

	return nil
}

// GetPlaceInfo finds information about a city from the tripadvisor api
func GetPlaceInfo(place string) (AreaInfo, error) {
	var areaInfo AreaInfo
	URL := fmt.Sprintf(travelAdvisorApi, url.QueryEscape(place))

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return areaInfo, err
	}

	req.Header.Add("x-rapidapi-key", rapidApiKey)
	req.Header.Add("x-rapidapi-host", "travel-advisor.p.rapidapi.com")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return areaInfo, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &areaInfo)
	if err != nil {
		return areaInfo, err
	}

	return areaInfo, nil
}

func GetWeather(start, end, location string) error {
	URL := fmt.Sprintf(weatherApiDecoded, start, location, end)
	URLenc := url.QueryEscape(URL)

	req, err := http.NewRequest("GET", URLenc, nil)
	if err != nil {
		return err
	}

	req.Header.Add("x-rapidapi-key", rapidApiKey)
	req.Header.Add("x-rapidapi-host", "visual-crossing-weather.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(body))
	return nil
}
