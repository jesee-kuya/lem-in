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

// containsNode checks if a node exists in a path.
func containsNode(path []int, node int) bool {
	for _, n := range path {
		if n == node {
			return true
		}
	}
	return false
}
