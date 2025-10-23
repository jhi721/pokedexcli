package main

import "fmt"

const locationsUrl = "https://pokeapi.co/api/v2/location-area"

func mapHandler(config *Config, url string) error {
	locations, err := config.ApiClient.GetLocations(url)
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

func commandMap(config *Config, arg string) error {
	url := locationsUrl
	if config.Next != "" {
		url = config.Next
	}

	if err := mapHandler(config, url); err != nil {
		return err
	}

	return nil
}

func commandMapb(config *Config, arg string) error {
	url := locationsUrl
	if config.Previous != "" {
		url = config.Previous
	}

	if err := mapHandler(config, url); err != nil {
		return err
	}

	return nil
}
