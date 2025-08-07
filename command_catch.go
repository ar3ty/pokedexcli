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

	if _, ok := cfg.pokedex[target]; ok {
		return errors.New("you already have one in your pokedex")
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
		fmt.Println("You may now inspect it with the inspect command")
	} else {
		fmt.Printf("%s escaped!\n", target)
	}

	return nil
}
