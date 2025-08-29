package main

import (
	"fmt"
)

func commandInspect(cfg *config) error {
	pokemon, ok := cfg.pokedex[*cfg.commandArgument]
	if !ok {
		fmt.Println("You have not caught this pokemon yet!")
		return nil
	}
	fmt.Println(pokemon.String())

	return nil
}
