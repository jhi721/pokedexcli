package main

import (
	"fmt"
)

func commandMap(config *Config) error {
	locationsUrl := "https://pokeapi.co/api/v2/location-area"
	if config.Next != "" {
		locationsUrl = config.Next
	}

	locations, err := config.ApiClient.Get(locationsUrl)
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

func commandMapb(config *Config) error {
	locationsUrl := "https://pokeapi.co/api/v2/location-area"
	if config.Previous != "" {
		locationsUrl = config.Previous
	}

	locations, err := config.ApiClient.Get(locationsUrl)
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
