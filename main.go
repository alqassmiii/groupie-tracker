package main

import (
	"fmt"
	API "groupie-tracker/API"
	"log"
	"net/http"
)



func main() {

	http.NewServeMux()

	http.Handle("/Static/", http.StripPrefix("/Static/", http.FileServer(http.Dir("./Static/")))) // Serve static CSS files
	http.HandleFunc("/", API.MainPage)
	http.HandleFunc("/artistInfo", API.ArtistPage)
	
	fmt.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

 