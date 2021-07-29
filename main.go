package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/SimonTanner/photo-album-maker/holiday"
	"github.com/SimonTanner/photo-album-maker/reader"
)

func main() {
	data, err := reader.ReadCSV("./reader/test_data/1.csv")
	if err != nil {
		log.Fatal(err)
	}

	hol := holiday.Holiday{}

	hol.SetDayPlaceData(data)
	hol.GetAbout()
	hol.CreateTitle()
	hol.GetAreaInfo()

	js, err := json.Marshal(hol)
	if err != nil {
		log.Fatal(err)
	}

	path := fmt.Sprintf("%s-%s.json", hol.Cities[0], hol.Country)
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	ioutil.WriteFile(path, js, 0644)
}
