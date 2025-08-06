package main

import (
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommandRegistry() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
