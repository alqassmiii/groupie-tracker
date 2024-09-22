package main


import (
	"fmt"
	server "groupie-tracker/server"
	"log"
	"net/http"
)

func main() {
	http.Handle("/ast/", http.StripPrefix("/ast/", http.FileServer(http.Dir("./ast/"))))
	http.HandleFunc("/", server.HomePage)
	http.HandleFunc("/artestPage", server.ArtestPage)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

