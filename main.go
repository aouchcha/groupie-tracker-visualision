package main

import (
	"fmt"
	"net/http"

	f "groupie-tracker-visualition/functions"
)

func main() {
	f.GetArtistData()
	f.GetRelationData()
	f.GetLocationData()
	f.GetDatesData()
	http.HandleFunc("/styles/", f.ServeStyle)
	http.HandleFunc("/", f.FirstPage)
	http.HandleFunc("/artist", f.OtherPages)
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
