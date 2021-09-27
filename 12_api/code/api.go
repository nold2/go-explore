package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// BaseURL is the base endpoint for the star wars API
const BaseURL = "https://swapi.dev/api/"
const PeopleURL = BaseURL + "people/"
const PlanetURL = BaseURL + "planets/"

type AllPeople struct {
	People []Person `json:"results"`
}

type Person struct {
	Name         string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld    Planet
}

type Planet struct {
	Name       string `json:"name"`
	Population string `json:"population"`
	Terrain    string `json:"terrain"`
}

func (p *Person) getHomeWorld() {
	res, err := http.Get(p.HomeworldURL)
	if err != nil {
		log.Print("Error fetching homeworld", err)
	}

	var data []byte
	if data, err = ioutil.ReadAll(res.Body); err != nil {
		log.Print("Error reading response body", err)
	}

	json.Unmarshal(data, &p.Homeworld)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(PeopleURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request star wars people")
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to parse request request body")
	}

	var people AllPeople
	if err := json.Unmarshal(data, &people); err != nil {
		fmt.Println("Error Parsing json", err)
	}

	for _, p := range people.People {
		p.getHomeWorld()
		fmt.Println(p)
	}
}

func main() {
	http.HandleFunc("/people", getPeople)
	fmt.Println("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
