package main

import (
	"bufio"
	"errors"
	"fmt"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"
)


type config struct {
	Next string
	Prev string
}


type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}


type Location struct {
	Name string
	Url  string
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
	}
}


func replLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	var cfg config
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
		err := cliCmd.callback(&cfg)
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
	if cfg.Next == "" {
		cfg.Next = "https://pokeapi.co/api/v2/location-area/"
	}
	req, err := http.NewRequest("GET", cfg.Next, nil)
	if err != nil {
		return fmt.Errorf("Error while setting up GET request in map: %v", err)
	}
	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error while doing GET request in map: %v", err)
	}
	defer res.Body.Close()

	var answer struct {
		Count    int         `json:count`
		Next     string      `json:next`
		Previous string      `json:previous`
		Results  []Location  `json:results`
	}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&answer); err != nil {
		return fmt.Errorf("Error while decoding json in map: %v", err)
	}
	for _, loc := range answer.Results {
		fmt.Println(loc.Name)
	}
	cfg.Next = answer.Next
	cfg.Prev = answer.Previous

	return nil
}
