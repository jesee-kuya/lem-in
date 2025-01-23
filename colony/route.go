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

var Res = [][]string{{"L1-t L2-h L3-0"}, {"L1-E L2-A L3-o L4-t L5-h L6-0"}, {"L1-a L2-c L3-n L4-E L5-A L6-o L7-t L8-h L9-0"}, {"L1-m L2-k L3-e L4-a L5-c L6-n L7-E L8-A L9-o L10-t"}, {"L1-end L2-end L3-end L4-m L5-k L6-e L7-a L8-c L9-n L10-E"}, {"L4-end L5-end L6-end L7-m L8-k L9-e L10-a"}, {"L7-end L8-end L9-end L10-m"}, {"L10-end"}}

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
