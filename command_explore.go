package main

import (
	"fmt"
)

func commandExplore(cfg *config) error {
	fmt.Printf("You asked to explore: %s\n", *cfg.commandArgument)
	return nil
}
