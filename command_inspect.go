package main

import (
	"fmt"
)

func commandInspect(config *Config, name string) error {
	if name == "" {
		fmt.Println("Please input pokemon to inspect")
		return nil
	}

	pokemon, exists := config.Pokedex[name]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	return nil
}
