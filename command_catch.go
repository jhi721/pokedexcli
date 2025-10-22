package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(config *Config, name string) error {
	if name == "" {
		fmt.Println("Please input pokemon to catch")
		return nil
	}

	_, exists := config.Pokedex[name]
	if exists {
		fmt.Printf("You already have %s\n", name)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := config.ApiClient.GetPokemonByName(name)
	if err != nil {
		return err
	}

	chance := rand.IntN(pokemon.BaseExperience)

	if pokemon.BaseExperience/2 > chance {
		fmt.Printf("%s was caught!\n", name)

		config.Pokedex[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
