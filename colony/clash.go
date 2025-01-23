package colony

import (
	"fmt"
	"reflect"
	"sort"
)

func Clash(paths [][]any) [][]any {
	if len(paths) <= 1 {
		return paths
	}

	// Filter and sort paths at the start.
	sPaths := filterAndSortPaths(paths)
	sortedPaths := Unique(sPaths)
	bestCombination := [][]any{sortedPaths[0]}

	// Try different combinations of paths.
	for i := 1; i < len(sortedPaths); i++ {
		candidatePath := sortedPaths[i]
		isCompatible := true

		// Check compatibility with all current best paths.
		for _, existingPath := range bestCombination {
			if !isGoodCombination([][]any{existingPath}, candidatePath) {
				isCompatible = false
				break
			}
		}

		// Add path only if it's compatible with all existing paths.
		if isCompatible {
			if len(bestCombination) > 1 && len(candidatePath) < len(bestCombination[len(bestCombination)-1]) {
				// If we already have multiple paths, only keep the better combination.
				bestCombination[len(bestCombination)-1] = candidatePath
			} else {
				bestCombination = append(bestCombination, candidatePath)
			}
		}
	}

	// Sort the final combination by path length.
	sort.Slice(bestCombination, func(i, j int) bool {
		return len(bestCombination[i]) < len(bestCombination[j])
	})

	return bestCombination
}

func isGoodCombination(existingPaths [][]any, newPath []any) bool {
	// Check if new path shares any intermediate nodes with existing paths.
	for _, existingPath := range existingPaths {
		for i := 1; i < len(newPath)-1; i++ {
			for j := 1; j < len(existingPath)-1; j++ {
				if newPath[i] == existingPath[j] {
					return false
				}
			}
		}
	}
	return true
}

func filterAndSortPaths(paths [][]any) [][]any {
	// Filter valid paths
	filtered := make([][]any, 0, len(paths))
	for _, path := range paths {
		if isValidPath(path) {
			filtered = append(filtered, path)
		}
	}

	// Sort paths by length and lexicographically if equal length.
	sort.Slice(filtered, func(i, j int) bool {
		if len(filtered[i]) != len(filtered[j]) {
			return len(filtered[i]) < len(filtered[j])
		}
		for k := 0; k < len(filtered[i]) && k < len(filtered[j]); k++ {
			if filtered[i][k] != filtered[j][k] {
				return filtered[i][k].(string) < filtered[j][k].(string)
			}
		}
		return false
	})

	return filtered
}

func isValidPath(path []any) bool {
	// Path should have at least 2 elements and no duplicates.
	if len(path) < 2 {
		return false
	}

	visited := make(map[any]bool)
	for _, node := range path {
		if visited[node] {
			return false
		}
		visited[node] = true
	}
	return true
}

func Unique(sortedPaths [][]any) [][]any {
	unique := [][]any{}
	seen := make(map[string]bool)

	for _, path := range sortedPaths {
		// Serialize each path as a string to ensure uniqueness.
		pathKey := serializePath(path)
		if !seen[pathKey] {
			seen[pathKey] = true
			unique = append(unique, path)
		}
	}

	// Call UniqueElements to remove any further unwanted duplicates if necessary
	return UniqueElements(unique)
}

func UniqueElements(holder [][]any) [][]any {
	uniquePaths := make([][]any, 0, len(holder))

	for _, path := range holder {
		if !containsPath(uniquePaths, path) {
			uniquePaths = append(uniquePaths, path)
		}
	}

	// If no unique paths found, return the shortest paths
	if len(uniquePaths) == 0 {
		minLength := shortest(holder)
		for _, path := range holder {
			if len(path) == minLength {
				uniquePaths = append(uniquePaths, path)
			}
		}
	}

	return uniquePaths
}

func containsPath(paths [][]any, path []any) bool {
	// Check if the path is already in the list of paths (deep comparison).
	for _, p := range paths {
		if reflect.DeepEqual(p, path) {
			return true
		}
	}
	return false
}

func serializePath(path []any) string {
	// Convert a path to a string for easy comparison.
	// You can adjust this for more complex structures, like JSON or another serialization format.
	return fmt.Sprintf("%v", path)
}

func shortest(holder [][]any) int {
	// Find the shortest path in terms of length.
	minLength := len(holder[0])
	for _, path := range holder {
		if len(path) < minLength {
			minLength = len(path)
		}
	}
	return minLength
}
