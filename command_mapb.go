package main

import (
	"fmt"

	"github.com/jhi721/pokedexcli/internal/pokeapi"
)

func commandMapb(config *Config) error {
	locationsUrl := "https://pokeapi.co/api/v2/location-area"
	if config.Previous != "" {
		locationsUrl = config.Previous
	}

	locations, err := pokeapi.GetLocations(locationsUrl)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	config.Next = locations.Next
	config.Previous = locations.Previous

	return nil
}
