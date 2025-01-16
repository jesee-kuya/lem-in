package main

import (
	"fmt"
	"os"
	"strings"

	"lem-in/colony"
	"lem-in/read"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [INPUT-FILE]")
		return
	}

	// Read input file
	input, err := read.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(input)

	// Parse number of ants
	lines := strings.Split(input, "\n")
	numberOfAnts := colony.ParseAnts(lines)

	if numberOfAnts == 0 {
		fmt.Println("Invalid number of ants")
		return
	}

	// Find all routes
	routes, err := colony.Route(input)
	if err != nil {
		fmt.Println("Error finding paths:", err)
		return
	}
	fmt.Println()
	colony.Path(routes, numberOfAnts)
}
