package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	cfg.pokeapiClient.Cache.StopChannel <- 1
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
