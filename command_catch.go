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
		stats := struct {
			hp              int
			attack          int
			defense         int
			special_attack  int
			special_defense int
			speed           int
		}{}
		for _, stat := range pokemonResp.Stats {
			switch stat.Stat.Name {
			case "hp":
				stats.hp = stat.BaseStat
			case "attack":
				stats.attack = stat.BaseStat
			case "defense":
				stats.defense = stat.BaseStat
			case "special-attack":
				stats.special_attack = stat.BaseStat
			case "special-defense":
				stats.special_defense = stat.BaseStat
			case "speed":
				stats.speed = stat.BaseStat
			}
		}
		types := []string{}
		for _, t := range pokemonResp.Types {
			types = append(types, t.Type.Name)
		}
		cfg.pokedex[pokemonResp.Name] = Pokemon{name: pokemonResp.Name, height: pokemonResp.Height, weight: pokemonResp.Weight, stats: stats, types: types}
	} else {
		fmt.Printf("You failed to catch %s. Try again!\n", pokemonResp.Name)
	}
	fmt.Printf("Your pokedex: %v\n", cfg.pokedex)
	return nil
}
