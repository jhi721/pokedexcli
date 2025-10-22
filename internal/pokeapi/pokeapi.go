package poke_api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Locations struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocations(url string) (Locations, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
	}

	var locations Locations
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		log.Fatal(err)
	}

	return locations, nil
}
