package main

import (
	"fmt"
)

func commandExplore(cfg *config) error {
	pokemonResp, err := cfg.pokeapiClient.GetPokemonList(cfg.commandArgument)
	if err != nil {
		return err
	}
	for _, enc := range pokemonResp.PokemonEncounters {
		fmt.Println(enc.Pokemon.Name)
	}
	return nil
}
