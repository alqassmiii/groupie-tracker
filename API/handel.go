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
		http.Error(w, "404 error", http.StatusNotFound)
		return
	}
	data := ArtistData()
	tmpl, err := template.ParseFiles("templates.html")
	if err != nil {
		http.Error(w, "500 swefderver error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "500 eeeserver error", http.StatusInternalServerError)
	}
}
func ArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artistInfo" {
		http.Error(w, "404 error", http.StatusNotFound)
		return
	}
	artistName := r.URL.Query().Get("name")
	if artistName == "" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}
	data := CollectData()
	artistExists := false
	var selectedArtist Data
	for _, artist := range data {
		if artist.A.Name == artistName { //artistName is the one from the url
			selectedArtist = artist
			artistExists = true
			break
		}
	}
	if !artistExists {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("artistPage.html")
	if err != nil {
		http.Error(w, "500 seeerver error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	fmt.Print(selectedArtist.L)
	err = tmpl.Execute(w, selectedArtist)
	if err != nil {
		http.Error(w, "500 ffeserver error", http.StatusInternalServerError)
	}
}
