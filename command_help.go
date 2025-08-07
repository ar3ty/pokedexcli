package main

import (
	"errors"
	"fmt"
)

func commandHelp(cfg *config, arg string) error {
	if arg != "" {
		return errors.New("command accepts no arguments")
	}

	fmt.Print("\nWelcome to the Pokedex!\nUsage:\n\n")
	for _, command := range getCommandRegistry() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()
	return nil
}
