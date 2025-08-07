package main

import (
	"time"

	"github.com/ar3ty/pokedexcli/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(10*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: client,
		backward:      false,
	}
	repl(cfg)
}
