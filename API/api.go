package API

import (
	"encoding/json"
	"log"
	"net/http"

)





func ArtistData() []Artist {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Fatalf("Error after request: %v", err)
	}
	defer resp.Body.Close()
	var artistInfo []Artist
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&artistInfo)
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
	var locationsResp LocationResp
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&locationsResp)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	// Extract locations from locationsResp
	//this loop add all the locations details into a slice of structs
	//beacuse of json layers we extract the loactions alone
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
	var datesResponse struct {
		Index []struct {
			Dates []string `json:"dates"`
		} `json:"index"`
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&datesResponse)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	// Extract dates from datesResponse
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
	var relationResponse struct {
		Index []struct {
			DatesLocations map[string][]string `json:"datesLocations"`
		} `json:"index"`
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&relationResponse)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}
	// Extract relations from relationResponse
	// this will sperate the dates and locations
	var relationInfo []Relation
	for _, item := range relationResponse.Index {
		relation := Relation{DatesLocations: item.DatesLocations}
		relationInfo = append(relationInfo, relation)
	}
	return relationInfo
}
func CollectData() []Data {
	artests := ArtistData()
	locations := LocationsData()
	dates := DatesData()
	relations := RelationData()
	dataData := make([]Data, len(artests))
	for i := 0; i < len(artests); i++ {
		dataData[i].A = artests[i]
		dataData[i].L = locations[i]
		dataData[i].D = dates[i]
		dataData[i].R = relations[i]
	}
	return dataData
}


