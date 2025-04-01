package main

import (
	"bufio"
	"fmt"
	"github.com/hontauadrian/pokedex/internal/pokeapi"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		val, ok := getCommands()[commandName]
		if ok {
			err := val.callback(cfg, args...)
			if err != nil {
				fmt.Print(err)
			}
			continue
		} else {
			fmt.Print("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
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
		"map": {
			name:        "map",
			description: "Display the name of 20 location areas in Pokemon world.",
			callback:    commandMap,
		},
		"mapb": {
			name: "map",
			description: "Map back. It's similar to the map command, however, " +
				"instead of displaying the next 20 locations, it displays the previous 20 locations.",
			callback: commandMapb,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Show a list of all the Pokémon located the entered area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Catching Pokemon adds them to the user's Pokedex.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: " It takes the name of a Pokemon and prints the name, height, weight, stats and type(s) of the Pokemon",
			callback:    commandInspect,
		},
	}
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
	pokedex          map[string]pokeapi.Pokemon
}
