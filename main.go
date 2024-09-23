package main

/*mport (
	"fmt"
	server "groupie-tracker/server"

	"encoding/json"
	"html/template"
	"image"
	"log"
	"net/http"
)

type Artest struct {
	// Define the fields of the Artest struct
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}


type Groupie struct {
	Name       string      `json:"name"`
	Age        int         `json:"age"`
	Image      string      `json:"image"`
	YearActive int         `json:"yearActive"`
	FirstAlbum string      `json:"firstAlbum"`
	Members    string      `json:"members"`
	Location   image.Point `json:"location"`
}



/*
func ArtestPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("artestPage.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, artest)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func main() {
	http.Handle("/ast/", http.StripPrefix("/ast/", http.FileServer(http.Dir("./ast/"))))
	http.HandleFunc("/", server.HomePage)
	http.HandleFunc("/artestPage", ArtestPage)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}




func getGroupie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groupies)
}



func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("artestPage.html")
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

// Removed duplicate ArtestPage function
*/
