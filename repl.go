package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ar3ty/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
	backward      bool
	pokedex       map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommandRegistry() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Each subsequent call displays the names of next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Each subsequent call displays the names of previous 20 locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Takes <location-area> as argument, shows list of pokemons located there",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Takes <pokemon> as an argument, adds target to a user's pokedex if succeded",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects characteristics of pokemon, if it has been caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Shows all the names of the Pokemon the user has caught",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	formattedText := strings.ToLower(text)
	words := strings.Fields(formattedText)
	return words
}

func repl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		arg := ""
		if len(words) == 2 {
			arg = words[1]
		}

		command, ok := getCommandRegistry()[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command.callback(cfg, arg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}
