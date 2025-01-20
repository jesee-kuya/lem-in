package colony

func Clash(paths [][]int) [][]int {
	if len(paths) <= 1 {
		return paths
	}

	sortedPaths := filterAndSortPaths(paths)
	bestCombination := [][]int{sortedPaths[0]}

	// Try different combinations of paths.
	for i := 1; i < len(sortedPaths); i++ {
		candidatePath := sortedPaths[i]
		isCompatible := true

		// Check compatibility with all current best paths.
		for _, existingPath := range bestCombination {
			if !isGoodCombination([][]int{existingPath}, candidatePath) {
				isCompatible = false
				break
			}
		}

		// Add path only if it's compatible with all existing paths.
		if isCompatible {
			if len(bestCombination) > 1 {
				// If we already have multiple paths, only keep the better combination.
				if len(candidatePath) < len(bestCombination[len(bestCombination)-1]) {
					bestCombination[len(bestCombination)-1] = candidatePath
				}
			} else {
				bestCombination = append(bestCombination, candidatePath)
			}
		}
	}

	// Sort the final combination to match expected order.
	for i := 0; i < len(bestCombination)-1; i++ {
		for j := i + 1; j < len(bestCombination); j++ {
			if len(bestCombination[i]) > len(bestCombination[j]) {
				bestCombination[i], bestCombination[j] = bestCombination[j], bestCombination[i]
			}
		}
	}
	return bestCombination
}

func isGoodCombination(existingPaths [][]int, newPath []int) bool {
	// Check if new path shares any intermediate nodes with existing paths.
	for _, existingPath := range existingPaths {
		sharedNodes := 0
		for i := 1; i < len(newPath)-1; i++ {
			for j := 1; j < len(existingPath)-1; j++ {
				if newPath[i] == existingPath[j] {
					sharedNodes++
				}
			}
			if sharedNodes != 0 {
				return false
			}
		}
	}
	return true
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

	// Sort paths primarily by length.
	for i := 0; i < len(filtered)-1; i++ {
		for j := i + 1; j < len(filtered); j++ {
			if shouldSwapPaths(filtered[i], filtered[j]) {
				filtered[i], filtered[j] = filtered[j], filtered[i]
			}
		}
	}
	return filtered
}

func shouldSwapPaths(path1, path2 []int) bool {
	// Primary sort by length.
	if len(path1) != len(path2) {
		return len(path1) > len(path2)
	}

	// Secondary sort by path values for consistent ordering.
	for i := 0; i < len(path1) && i < len(path2); i++ {
		if path1[i] != path2[i] {
			return path1[i] > path2[i]
		}
	}
	return false
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
