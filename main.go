package main

import (
	"fmt"
	"groupie-tracker/API"
	"html/template"
	"log"
	"net/http"
)

var AA []API.Data

func init() {
	AA = API.CollectData()
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("artestPage.html") // Updated path
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, AA) // Updated variable
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}

}

func ArtestPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the artestPage!")
}

func main() {
	printAA()
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/artestPage", ArtestPage)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
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
