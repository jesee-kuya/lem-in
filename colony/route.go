package colony

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Graph represents the ant colony.
type Graph map[int][]int

// parseGraph converts input into an adjacency list representation.
func parseGraph(input string) (Graph, int, int, error) {
	lines := strings.Split(input, "\n")
	graph := make(Graph)
	startRoom, endRoom := -1, -1
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

			from, err1 := strconv.Atoi(parts[0])
			to, err2 := strconv.Atoi(parts[1])

			if err1 != nil || err2 != nil {
				return nil, -1, -1, fmt.Errorf("invalid room numbers in tunnel: %s", line)
			}

			graph[from] = append(graph[from], to)
			graph[to] = append(graph[to], from)

		default:
			parts := strings.Split(line, " ")
			if len(parts) == 3 {
				room, err := strconv.Atoi(parts[0])
				if err != nil {
					continue
				}

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
func Route(input string) ([][]int, error) {
	graph, start, end, err := parseGraph(input)
	if err != nil {
		return nil, err
	}

	if start == -1 || end == -1 {
		return nil, errors.New("invalid start or end room")
	}

	var paths [][]int
	visited := make(map[int]bool)
	currentPath := []int{start}
	visited[start] = true

	findPaths(graph, start, end, visited, currentPath, &paths)

	if len(paths) == 0 {
		return nil, errors.New("no paths found")
	}

	return filterOptimalPaths(paths), nil
}

func findPaths(graph Graph, current, end int, visited map[int]bool, currentPath []int, paths *[][]int) {
	if current == end {
		pathCopy := make([]int, len(currentPath))
		copy(pathCopy, currentPath)
		*paths = append(*paths, pathCopy)
		return
	}

	for _, next := range graph[current] {
		if !visited[next] {
			visited[next] = true
			currentPath = append(currentPath, next)

			findPaths(graph, next, end, visited, currentPath, paths)

			currentPath = currentPath[:len(currentPath)-1]
			visited[next] = false
		}
	}
}

func filterOptimalPaths(paths [][]int) [][]int {
	if len(paths) == 0 {
		return paths
	}

	// Create a map to store unique paths
	uniquePaths := make(map[string][]int)

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
	result := make([][]int, 0, len(uniquePaths))
	for _, path := range uniquePaths {
		result = append(result, path)
	}

	return result
}
