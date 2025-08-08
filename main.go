package main

import (
	"fmt"
	"slices"
	"strings"
)

func cleanInput(text string) []string {
	splitStrings := slices.DeleteFunc(strings.Split(strings.TrimSpace(strings.ToLower(text)), " "), func (s string) bool { return s == "" })
	return splitStrings
}

func main () {
	fmt.Println("Hello, World!")
	fmt.Println(cleanInput("Hello, World!"))
}
