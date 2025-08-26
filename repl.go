package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/PwnySQL/pokedex_cli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsUrl *string
	prevLocationsUrl *string
	commandArgument  *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
	arguments   []string
}

func getCommandRegistry() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Program",
			callback:    commandExit,
			arguments:   []string{},
		},
		"help": {
			name:        "help",
			description: "Display the help message",
			callback:    commandHelp,
			arguments:   []string{},
		},
		"map": {
			name:        "map",
			description: "Show current locations and go to next area",
			callback:    commandMap,
			arguments:   []string{},
		},
		"mapb": {
			name:        "mapb",
			description: "Show current locations and go to previous area",
			callback:    commandMapb,
			arguments:   []string{},
		},
		"explore": {
			name:        "explore",
			description: "Show pokemons in the passed area",
			callback:    commandExplore,
			arguments:   []string{"area"},
		},
		"catch": {
			name:        "catch",
			description: "Catch the pokemon given by name",
			callback:    commandCatch,
			arguments:   []string{"name"},
		},
	}
}

func replLoop(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if hasToken := scanner.Scan(); !hasToken {
			break
		}
		userInput := scanner.Text()
		words := cleanInput(userInput)
		if len(words) <= 0 {
			continue
		}
		cliCmd, ok := getCommandRegistry()[words[0]]
		if !ok {
			fmt.Printf("Unknown command: %s\n", words[0])
			continue
		}
		cfg.commandArgument = nil
		if len(words) > 1 {
			cfg.commandArgument = &words[1]
		}
		err := cliCmd.callback(cfg)
		if err != nil {
			fmt.Printf("Error while executing '%s': %s: \n", cliCmd.name, err.Error())
			continue
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(strings.ToLower(text))
	return words
}
