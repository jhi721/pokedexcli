package main

import "fmt"

func commandExplore(config *Config, city string) error {
	if city == "" {
		fmt.Println("Please input area to explore")
		return nil
	}

	fmt.Printf("Exploring %s\n", city)

	locationArea, err := config.ApiClient.GetLocationAreaByName(city)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
