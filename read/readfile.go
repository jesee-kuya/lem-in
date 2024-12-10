package read

import (
	"bufio"
	"os"
	"strings"
)

func ReadFile(filePath string) (string, error) {
	// Open file in the filepath
	file, err := os.Open(filePath)
	if err != nil {
		return "Error opening file", err
	}
	defer file.Close()

	// read file contents
	scanner := bufio.NewScanner(file)
	var contents []string
	for scanner.Scan() {
		contents = append(contents, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "Error reading file", err
	}

	return strings.Join(contents, "\n"), nil
}
