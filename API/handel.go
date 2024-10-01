package API

import (
	"fmt"
	"net/http"
	"strings"
	"text/template"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path != "/" {
		w.WriteHeader(http.StatusNotFound)
		renderErrorPage(w, "404Error.html", "404 - Not Found")
		return
	}
	data := ArtistData()
	tmpl, err := template.ParseFiles("Static/HomePage.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		renderErrorPage(w, "500Error.html", "500 - Internal Server Error")
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err = tmpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		renderErrorPage(w, "500Error.html", "500 - Internal Server Error")
	}
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artistInfo" {
		w.WriteHeader(http.StatusNotFound)
		renderErrorPage(w, "404Error.html", "404 - Not Found")
		return
	}
	artistName := r.URL.Query().Get("name")
	if artistName == "" {
		w.WriteHeader(http.StatusBadRequest)
		renderErrorPage(w, "400Error.html", "400 - Bad Request")
		return
	}
	data := CollectData()
	artistExists := false
	var selectedArtist Data
	for _, artist := range data {
		if artist.A.Name == artistName {
			selectedArtist = artist
			artistExists = true
			break
		}
	}
	if !artistExists {
		w.WriteHeader(http.StatusNotFound)
		renderErrorPage(w, "404Error.html", "404 - Not Found")
		return
	}
	tmpl, err := template.ParseFiles("Static/ArtistPage.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		renderErrorPage(w, "500Error.html", "500 - Internal Server Error")
		return
	}
	w.Header().Set("Content-Type", "text/html")
	fmt.Print(selectedArtist.L)
	err = tmpl.Execute(w, selectedArtist)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		renderErrorPage(w, "500Error.html", "500 - Internal Server Error")
	}
}


