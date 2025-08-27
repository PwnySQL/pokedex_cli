package main

import (
	"github.com/PwnySQL/pokedex_cli/internal/pokeapi"
)

type Pokemon struct {
	name           string
	baseExperience int
}

type Pokedex map[string]Pokemon

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsUrl *string
	prevLocationsUrl *string
	commandArgument  *string
	pokedex          Pokedex
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
	arguments   []string
}
