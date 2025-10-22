package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jhi721/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type Config struct {
	Next      string
	Previous  string
	ApiClient pokeapi.PokeApiClient
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Iterate over locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Iterate over locations",
			callback:    commandMapb,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := Config{
		Next:      "",
		Previous:  "",
		ApiClient: pokeapi.NewPokeApiClient(),
	}

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()

		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			continue
		}

		commandName := input[0]

		commands := getCommands()

		command, ok := commands[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(&config)
		if err != nil {
			fmt.Println(err)
		}
	}
}
