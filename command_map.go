package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	var directionURL *string

	if cfg.backward {
		directionURL = cfg.previous
	} else {
		directionURL = cfg.next
	}

	locs, err := cfg.pokeapiClient.GetLocationList(directionURL)
	if err != nil {
		return err
	}

	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}

	cfg.previous = locs.Previous
	cfg.next = locs.Next

	cfg.backward = false

	return nil
}

func commandMapB(cfg *config) error {
	if cfg.previous == nil {
		return errors.New("you're on the first page")
	}

	cfg.backward = true
	return commandMap(cfg)
}
