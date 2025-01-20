package colony

import (
	"errors"
	"fmt"
	"strings"
)

// Graph represents the ant colony.
type Graph map[any][]any

// parseGraph converts input into an adjacency list representation.
func parseGraph(input string) (Graph, any, any, error) {
	lines := strings.Split(input, "\n")
	graph := make(Graph)
	var startRoom, endRoom any
	startRoom, endRoom = -1, -1
	isStart, isEnd := false, false

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			if line == "##start" {
				isStart = true
				isEnd = false
				continue
			}
			if line == "##end" {
				isEnd = true
				isStart = false
				continue
			}
			if strings.HasPrefix(line, "#") {
				continue
			}
		}

		switch {
		case strings.Contains(line, "-"):
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, -1, -1, fmt.Errorf("invalid tunnel format: %s", line)
			}

			from := parts[0]
			to := parts[1]

			graph[from] = append(graph[from], to)
			graph[to] = append(graph[to], from)

		default:
			parts := strings.Split(line, " ")
			if len(parts) == 3 {
				room := parts[0]

				if isStart {
					startRoom = room
					isStart = false
				} else if isEnd {
					endRoom = room
					isEnd = false
				}
			}
		}
	}

	if startRoom == -1 || endRoom == -1 {
		return nil, -1, -1, errors.New("start or end room not found")
	}

	return graph, startRoom, endRoom, nil
}

// Route finds all routes between start and end.
func Route(input string) ([][]any, error) {
	graph, start, end, err := parseGraph(input)
	if err != nil {
		return nil, err
	}

	if start == -1 || end == -1 {
		return nil, errors.New("invalid start or end room")
	}

	var paths [][]any
	visited := make(map[any]bool)
	currentPath := []any{start}
	visited[start] = true

	FindPaths(graph, start, end, visited, currentPath, &paths)

	if len(paths) == 0 {
		return nil, errors.New("no paths found")
	}

	return FilterOptimalPaths(paths), nil
}

func FindPaths(graph Graph, current, end any, visited map[any]bool, currentPath []any, paths *[][]any) {
	if current == end {
		pathCopy := make([]any, len(currentPath))
		copy(pathCopy, currentPath)
		*paths = append(*paths, pathCopy)
		return
	}

	for _, next := range graph[current] {
		if !visited[next] {
			visited[next] = true
			currentPath = append(currentPath, next)

			FindPaths(graph, next, end, visited, currentPath, paths)

			currentPath = currentPath[:len(currentPath)-1]
			visited[next] = false
		}
	}
}

func FilterOptimalPaths(paths [][]any) [][]any {
	if len(paths) == 0 {
		return paths
	}

	// Create a map to store unique paths
	uniquePaths := make(map[string][]any)

	// Filter paths that don't lead to cycles or unnecessary detours
	for _, path := range paths {
		key := fmt.Sprintf("%v", path)
		isOptimal := true

		// Check if this path is optimal
		for _, otherPath := range paths {
			if len(otherPath) < len(path) {
				// Check if this is a subpath
				isSubpath := true
				for i := range otherPath {
					if i >= len(path) || path[i] != otherPath[i] {
						isSubpath = false
						break
					}
				}
				if isSubpath {
					isOptimal = false
					break
				}
			}
		}

		if isOptimal {
			uniquePaths[key] = path
		}
	}

	// Convert back to slice
	result := make([][]any, 0, len(uniquePaths))
	for _, path := range uniquePaths {
		result = append(result, path)
	}

	return result
}
