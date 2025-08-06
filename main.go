package main

import (
	"time"

	"github.com/ar3ty/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: client,
		backward:      false,
	}
	repl(cfg)
}
