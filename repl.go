package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/PwnySQL/pokedex_cli/internal/pokeapi"
)


type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsUrl *string
	prevLocationsUrl *string
}


type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}


func getCommandRegistry() map[string] cliCommand {
	return map[string]cliCommand {
		"exit": {
			name: "exit",
			description: "Exit the Program",
			callback: commandExit,
		},
		"help": {
			name: "help",
			description: "Display the help message",
			callback: commandHelp,
		},
		"map": {
			name: "map",
			description: "Show current locations and go to next area",
			callback: commandMap,
		},
			"mapb": {
			name: "mapb",
			description: "Show current locations and go to previous area",
			callback: commandMapb,
		},
	}
}


func replLoop(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for ;; {
		fmt.Print("Pokedex > ")
		if hasToken := scanner.Scan(); !hasToken {
			break;
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

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("os.Exit(0) did not work")
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Pokedex is an interactive program to query information about Pokemon.")
	fmt.Println()
	fmt.Println("Available commands:")
	commandRegistry := getCommandRegistry()
	for _, cliCmd := range commandRegistry {
		fmt.Printf("%s: %s\n", cliCmd.name, cliCmd.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	locationResp, err := cfg.pokeapiClient.GetLocationList(cfg.nextLocationsUrl)
	if err != nil {
		return err
	}
	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	cfg.nextLocationsUrl = locationResp.Next
	cfg.prevLocationsUrl = locationResp.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsUrl == nil {
		return errors.New("you're on the first page")
	}
	locationResp, err := cfg.pokeapiClient.GetLocationList(cfg.prevLocationsUrl)
	if err != nil {
		return err
	}
	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	cfg.nextLocationsUrl = locationResp.Next
	cfg.prevLocationsUrl = locationResp.Previous

	return nil
}
