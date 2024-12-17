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

	// Indicates that the startRoom and endRoom are not yet set.
	startRoom, endRoom := -1, -1

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Check if line defines a room or tunnel.
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
		case strings.HasPrefix(line, "##start"), strings.HasPrefix(line, "##end"):
			continue
		default:
			parts := strings.Split(line, " ")
			if len(parts) == 3 {
				room, err := strconv.Atoi(parts[0])
				if err != nil {
					continue
				}

				// Check if this is start or end room based on previous lines.
				if startRoom == -1 {
					startRoom = room
				} else if endRoom == -1 {
					endRoom = room
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
	// Parse the input into a graph and start/end rooms
	graph, start, end, err := parseGraph(input)
	if err != nil {
		return nil, err
	}

	// Check if start and end are valid
	if start == -1 || end == -1 {
		return nil, errors.New("invalid start or end room")
	}

	visited := make(map[int]bool)
	var paths [][]int

	// Perform depth-first search
	dfs(graph, start, end, visited, []int{}, &paths)

	if len(paths) == 0 {
		return nil, errors.New("no paths found")
	}

	return paths, nil
}

// dfs performs depth-first search.
func dfs(graph Graph, current, end int, visited map[int]bool, path []int, paths *[][]int) {
	// Mark current room as visited.
	visited[current] = true
	path = append(path, current)

	if current == end {
		// Create a copy of the path and add to paths if we reach the end room.
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		*paths = append(*paths, pathCopy)
		visited[current] = false
		return
	}

	// Explore connected rooms.
	for _, neighbor := range graph[current] {
		if !visited[neighbor] {
			dfs(graph, neighbor, end, visited, path, paths)
		}
	}
	visited[current] = false
}
