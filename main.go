package main

import (
	"fmt"
	API "groupie-tracker/API"
	"log"
	"net/http"
)

var AA []API.Data

func init() {
	AA = API.CollectData()
}
//tgyhujik
func main() {
	
	mux := http.NewServeMux()
	mux.Handle("/Static/", http.StripPrefix("/Static/", http.FileServer(http.Dir("./Static/")))) // Serve static CSS files
	mux.HandleFunc("/", API.MainPage)
	mux.HandleFunc("/artistInfo", API.ArtistPage)
	
	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func printAA() {
	for _, data := range AA {
		fmt.Printf("Artist: %+v\n", data.A)
		fmt.Printf("Relation: %+v\n", data.R)
		fmt.Printf("Location: %+v\n", data.L)
		fmt.Printf("Dates: %+v\n", data.D)
		fmt.Println("---")
	}
}
