package main

import (
	"fmt"
	"os"

	"lem-in/read"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: go run main.go [input_file]")
		os.Exit(1)
	}
	// process the file contents
	file, err := read.ReadFile(args[0])
	if err != nil {
		fmt.Println("ERROR", err)
	}
	fmt.Println(file)
}
