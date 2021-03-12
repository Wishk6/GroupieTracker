package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

//Artists is...
type Artists struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate string   `json:"creationsdates"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

//Locations is...
type Locations struct {
	ID        int    `json:"id"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
}

//Dates is..
type Dates struct {
	ID    int    `json:"id"`
	Dates string `json:"dates"`
}

//Relation is..
type Relation struct {
	ID             int `json:"id"`
	DatesLocations int `json:"datesLocations"`
}

/*--------------------------MAIN----------------------------*/

func main() {
	http.HandleFunc("/yo", GetArtistsData)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.ListenAndServe(":8080", nil)
}

/*------------------------ARTISTS----------------------------*/

//GetArtistsData func
func GetArtistsData(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/1")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var Data Artists
	json.Unmarshal(responseData, &Data)
	t, err := template.ParseFiles("static/index.html")
	t.Execute(w, Data)
}

/*-----------------------LOCATIONS--------------------------*/

//GetLocationsData func
func GetLocationsData(id string) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + id)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var Data Locations
	json.Unmarshal(responseData, &Data)

}

/*--------------------------DATES---------------------------*/

//GetDatesData func
func GetDatesData(id string) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + id)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var Data Artists
	json.Unmarshal(responseData, &Data)

}

/*--------------------------RELATIONS---------------------------*/

//GetRelationData func
func GetRelationData(id string) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/relation/" + id)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var Data Relation
	json.Unmarshal(responseData, &Data)
}
