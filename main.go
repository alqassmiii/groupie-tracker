package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("artestPage.html") // Updated path
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
	return
}

func ArtestPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the artestPage!")
}

func main() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/artestPage", ArtestPage)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}