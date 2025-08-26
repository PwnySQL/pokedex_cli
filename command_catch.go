package main

import (
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(cfg *config) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", *cfg.commandArgument)
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(cfg.commandArgument)
	if err != nil {
		return err
	}
	fmt.Printf("Base of %s is: %d\n", pokemonResp.Name, pokemonResp.BaseExperience)
	// The maximum base experience is 608 of pokemon blissey, id 242.
	// The idea is that one can catch blissey in 1 of 10 tries.
	chance := rand.Intn(int(math.Trunc(1.1 * (608.0 + 1.0))))
	if chance > pokemonResp.BaseExperience {
		fmt.Printf("You caught %s\n", pokemonResp.Name)
		cfg.pokedex[pokemonResp.Name] = Pokemon{name: pokemonResp.Name, baseExperience: pokemonResp.BaseExperience}
	} else {
		fmt.Printf("You failed to catch %s. Try again!\n", pokemonResp.Name)
	}
	fmt.Printf("Your pokedex: %v\n", cfg.pokedex)
	return nil
}
