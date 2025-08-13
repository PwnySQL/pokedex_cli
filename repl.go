package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)


type config struct {
}


type cliCommand struct {
	name        string
	description string
	callback    func(config) error
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
	}
}


func replLoop() {
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
		var cfg config
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

func commandExit(cfg config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New("os.Exit(0) did not work")
}

func commandHelp(cfg config) error {
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
