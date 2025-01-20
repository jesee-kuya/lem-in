package main

import (
	"fmt"
	"os"
	"strings"

	"lem-in/colony"
	"lem-in/read"
)

func ParseAnts(lines []string) int {
	if len(lines) == 0 {
		return 0
	}

	var ants int
	_, err := fmt.Sscanf(lines[0], "%d", &ants)
	if err != nil {
		fmt.Println("Error parsing number of ants:", err)
		return 0
	}
	return ants
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [INPUT-FILE]")
		return
	}

	input, err := read.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(input)

	routes, err := colony.Route(input)
	if err != nil {
		fmt.Println("Error finding paths:", err)
		return
	}
	fmt.Println()

	routes = colony.Clash(routes)

	// Parse number of ants
	lines := strings.Split(input, "\n")
	numberOfAnts := ParseAnts(lines)
	paths := colony.Path(routes, numberOfAnts)

	for i := 0; i < len(paths); i++ {
		fmt.Println(paths[i])
	}
}
