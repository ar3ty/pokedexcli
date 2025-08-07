package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, target string) error {
	if target == "" {
		return errors.New("command requires a target")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(target)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", target)

	chance := rand.Intn(pokemon.BaseExperience)
	if chance < 40 {
		cfg.pokedex[target] = pokemon
		fmt.Printf("%s was caught!\n", target)
	} else {
		fmt.Printf("%s escaped!\n", target)
	}

	return nil
}
