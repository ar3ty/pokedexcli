package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, name string) error {
	if name == "" {
		return errors.New("pokemon name as an argument is expected")
	}

	pok, ok := cfg.pokedex[name]
	if !ok {
		return errors.New("you have not caught this pokemon")
	}

	fmt.Printf("Name: %s\n", pok.Name)
	fmt.Printf("Heigth: %d\n", pok.Height)
	fmt.Printf("Weight: %d\n", pok.Weight)
	fmt.Printf("Stats:\n")
	for _, value := range pok.Stats {
		fmt.Printf("  -%s: %d\n", value.Stat.Name, value.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, value := range pok.Types {
		fmt.Printf("  - %s\n", value.Type.Name)
	}

	return nil
}
