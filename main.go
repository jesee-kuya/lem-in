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

	routes, err := colony.Route(input)
	if err != nil {
		fmt.Println("ERROR: invalid data format,", err)
		return
	}

	routes = colony.Clash(routes)

	// Parse number of ants
	lines := strings.Split(input, "\n")
	numberOfAnts := ParseAnts(lines)
	paths := colony.Path(routes, numberOfAnts)
	if len(paths) == 0 {
		return
	}
	fmt.Println(input)
	fmt.Println()

	for _, path := range paths {
		fmt.Println(strings.Join(path, " "))
	}
}
