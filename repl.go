package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type config struct {
	Next     string
	Previous string
	Backward bool
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
	}
}

func cleanInput(text string) []string {
	formattedText := strings.ToLower(text)
	words := strings.Fields(formattedText)
	return words
}

func repl() {
	reader := bufio.NewScanner(os.Stdin)
	cfg := config{
		Next:     "https://pokeapi.co/api/v2/location-area",
		Previous: "https://pokeapi.co/api/v2/location-area",
		Backward: false,
	}

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		command, ok := getCommandRegistry()[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command.callback(&cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}
