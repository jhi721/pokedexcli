package main

import "fmt"

func commandHelp(config *Config, arg string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for _, val := range getCommands() {
		fmt.Printf("%s: %s\n", val.name, val.description)
	}

	return nil
}
