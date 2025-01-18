package colony

func Clash(paths [][]int) [][]int {
	if len(paths) <= 1 {
		return paths
	}

	sortedPaths := filterAndSortPaths(paths)
	result := [][]int{sortedPaths[0]}

	// Add each remaining path if it doesn't create conflicts.
	for i := 0; i < len(sortedPaths); i++ {
		candidatePath := sortedPaths[i]

		// Skip paths that are longer than the shortest path
		if len(candidatePath) > len(sortedPaths[0])*2 {
			continue
		}

		// Check if this path can work with existing paths
		if isCompatiblePath(result, candidatePath) {
			result = append(result, candidatePath)
		}
	}
	return result
}

// isCompatiblePath checks if a new path can work with existing paths.
func isCompatiblePath(existingPaths [][]int, newPath []int) bool {
	existingNodes := make(map[int]bool)
	for _, path := range existingPaths {
		for i := 1; i < len(path)-1; i++ {
			existingNodes[path[i]] = true
		}
	}

	// Count shared nodes between new path and existing paths.
	sharedNodes := 0
	for i := 1; i < len(newPath)-1; i++ {
		if existingNodes[newPath[i]] {
			sharedNodes++
		}
	}

	// Calculate path overlap ratio by excluding start and end nodes.
	pathLength := len(newPath) - 2 
	if pathLength == 0 {
		return false
	}

	overlapRatio := float64(sharedNodes) / float64(pathLength)

	// If the overlap ratio is too high, paths are incompatible.
	return overlapRatio <= 0.5
}

func filterAndSortPaths(paths [][]int) [][]int {
	if len(paths) == 0 {
		return paths
	}

	filtered := make([][]int, 0)

	for _, path := range paths {
		if isValidPath(path) {
			filtered = append(filtered, path)
		}
	}

	// Sort paths by length.
	for i := 0; i < len(filtered)-1; i++ {
		for j := i + 1; j < len(filtered); j++ {
			if len(filtered[i]) > len(filtered[j]) {
				filtered[i], filtered[j] = filtered[j], filtered[i]
			}
		}
	}
	return filtered
}

func isValidPath(path []int) bool {
	if len(path) < 2 {
		return false
	}

	// Check for duplicates
	visited := make(map[int]bool)

	for _, node := range path {
		if visited[node] {
			return false
		}
		visited[node] = true
	}
	return true
}

// containsNode checks if a node exists in a path.
func containsNode(path []int, node int) bool {
	for _, n := range path {
		if n == node {
			return true
		}
	}
	return false
}
