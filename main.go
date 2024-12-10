package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Usage: go run main.go [input_file]")
		os.Exit(1)
	}
}
