package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config, arg string) error {
	if arg != "" {
		return errors.New("command accepts no arguments")
	}

	cfg.pokeapiClient.Cache.StopChannel <- 1
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
