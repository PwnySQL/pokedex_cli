package main

import (
	"fmt"
	"strings"
)

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Pokedex is an interactive program to query information about Pokemon.")
	fmt.Println()
	fmt.Println("Available commands:")
	commandRegistry := getCommandRegistry()
	for _, cliCmd := range commandRegistry {
		var b strings.Builder
		for idx, arg := range cliCmd.arguments {
			sep := " "
			if idx == len(cliCmd.arguments)-1 {
				sep = ""
			}
			fmt.Fprintf(&b, "<%s>%s", arg, sep)
		}
		sep := " "
		if b.Len() <= 0 {
			sep = ""
		}
		fmt.Printf("%s%s%s: %s\n", cliCmd.name, sep, b.String(), cliCmd.description)
	}
	return nil
}
