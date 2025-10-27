package main

import (
	"fmt"
)

func commandPokedex(cfg *config) error {
	fmt.Println("Your Pokedex:")
	for _, p := range cfg.pokedex {
		fmt.Printf("  - %s\n", p.name)
	}
	return nil
}
