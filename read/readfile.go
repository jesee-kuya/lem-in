package read

import (
	"bufio"
	"os"
	"strings"
)

// ReadFile reads a file and returns the contents of the file 
func ReadFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "Error opening file", err
	}
	defer file.Close()

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
