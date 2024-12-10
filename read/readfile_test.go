package read

import (
	"os"
	"testing"
)

var TestCases = []struct {
	name     string
	filepath string
	content  []byte
	output   string
}{
	{"Test1", "test1.txt", []byte("Hello World"), "Hello World"},
}

func TestReadFile(t *testing.T) {
	for _, tc := range TestCases {
		err := os.WriteFile(tc.filepath, tc.content, 0o666)
		if err != nil {
			t.Errorf("Error creating file %v", tc.filepath)
		}
		content, _ := ReadFile(tc.filepath)

		if content != tc.output {
			t.Errorf("Expected\ncontent: %v but got\ncontent %v", tc.output, content)
		}
	}
}
