package main

import "errors"

func commandCatch(cfg *config, target string) error {
	if target == "" {
		return errors.New("command requires a target")
	}

	//pokemon, err := cfg.pokeapiClient.

	return nil
}
