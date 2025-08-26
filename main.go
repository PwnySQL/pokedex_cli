package main

import (
	"time"

	"github.com/PwnySQL/pokedex_cli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       make(Pokedex),
	}
	replLoop(cfg)
}
