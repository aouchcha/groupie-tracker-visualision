package main

import (
	"fmt"
	"net/http"

	f "groupie-tracker/functions"
)

// Define a struct that matches the JSON structure

func main() {
	f.GetArtistData()
	f.GetRelationData()
	f.GetLocationData()
	f.GetDatesData()
	http.HandleFunc("/styles/", f.ServeStyle)
	http.HandleFunc("/", f.FirstPage)
	http.HandleFunc("/artist", f.OtherPages)
	fmt.Println("http://localhost:7878")
	http.ListenAndServe(":7878", nil)
}
