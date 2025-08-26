package main

type Pokemon struct {
	name           string
	baseExperience int
}

type Pokedex map[string]Pokemon
