package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, name string) error {
	if name != "" {
		return errors.New("command accepts no arguments")
	}

	fmt.Println("Your Pokedex:")
	for key := range cfg.pokedex {
		fmt.Printf(" - %s\n", key)
	}

	return nil
}
