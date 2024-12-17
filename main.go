package main

import (
	"fmt"
	"os"

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

	// Parse graph
	graph, start, end, err := colony.ParseGraph(input)
	if err != nil {
		fmt.Println("Error parsing graph:", err)
		return
	}

	// Find all routes
	routes, err := colony.Route(graph, start, end)
	if err != nil {
		fmt.Println("Error finding paths:", err)
		return
	}

	fmt.Println("Routes:")
	for _, route := range routes {
		fmt.Println(route)
	}
}
