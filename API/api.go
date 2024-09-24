package API

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Artist, Location, Dates, Relation, Data structs need to be defined here

func ArtistData() []Artist {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatalf("Error after request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var artistInfo []Artist
	err = json.Unmarshal(body, &artistInfo)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	return artistInfo
}

func LocationsData() []Location {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		log.Fatalf("Error after request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var locationsResp LocationResp
	err = json.Unmarshal(body, &locationsResp)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	var locations []Location
	for _, item := range locationsResp.Index {
		location := Location{locations: item.Locations}
		locations = append(locations, location)
	}
	return locations
}

func DatesData() []Dates {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		log.Fatalf("Error after request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var datesResponse struct {
		Index []struct {
			Dates []string `json:"dates"`
		} `json:"index"`
	}
	err = json.Unmarshal(body, &datesResponse)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	var datesInfo []Dates
	for _, item := range datesResponse.Index {
		dates := Dates{Dates: item.Dates}
		datesInfo = append(datesInfo, dates)
	}
	return datesInfo
}

func RelationData() []Relation {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		log.Fatalf("Error after request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var relationResponse struct {
		Index []struct {
			DatesLocations map[string][]string `json:"datesLocations"`
		} `json:"index"`
	}
	err = json.Unmarshal(body, &relationResponse)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	var relationInfo []Relation
	for _, item := range relationResponse.Index {
		relation := Relation{DatesLocations: item.DatesLocations}
		relationInfo = append(relationInfo, relation)
	}
	return relationInfo
}

func CollectData() []Data {
	artists := ArtistData()
	locations := LocationsData()
	dates := DatesData()
	relations := RelationData()

	dataData := make([]Data, len(artists))
	for i := 0; i < len(artists); i++ {
		dataData[i].A = artists[i]
		dataData[i].L = locations[i]
		dataData[i].D = dates[i]
		dataData[i].R = relations[i]
	}
	return dataData
}
