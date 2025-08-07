package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, area string) error {
	if area == "" {
		return errors.New("argument is expected")
	}

	pocs, err := cfg.pokeapiClient.GetPokemonList(area)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found pokemon:")

	for _, encounter := range pocs.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
