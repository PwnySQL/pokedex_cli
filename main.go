package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func cleanInput(text string) []string {
	splitStrings := slices.DeleteFunc(strings.Split(strings.TrimSpace(strings.ToLower(text)), " "), func (s string) bool { return s == "" })
	return splitStrings
}

func main () {
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
		fmt.Printf("Your command was: %s\n", words[0])
	}
}
