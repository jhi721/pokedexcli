package main

import (
	"fmt"
)

func commandPokedex(config *Config, _ string) error {
	if len(config.Pokedex) == 0 {
		fmt.Println("No catched pokemons found")
		return nil
	}

	fmt.Println("Your Pokedex:")

	for key := range config.Pokedex {
		fmt.Printf(" - %s\n", key)
	}

	return nil
}
