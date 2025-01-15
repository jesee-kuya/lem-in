package colony

func Clash(paths [][]int) [][]int {
	var routes [][]int
	if len(paths) <= 1 {
		return paths
	}

	// sort routes by length (Dijkstra's Algorithm)
	sortRoutes(routes)

	result := [][]int{routes[0]}

	for i := 1; i < len(routes); i++ {
		if !clashingPaths(result, routes[i]) {
			result = append(result, routes[i])
		}
	}

	return result
}

func sortRoutes(routes [][]int) {
	for i := 0; i < len(routes)-1; i++ {
		for j := i + 1; j < len(routes); j++ {
			if len(routes[i]) > len(routes[j]) {
				routes[i], routes[j] = routes[j], routes[i]
			}
		}
	}
}

// checks for conflicts in paths
func clashingPaths(definedRoutes [][]int, newRoutes []int) bool {
	// check for shared intermediate nodes
	set := make(map[int]bool)

	// add all nodes from route1 to one set (exclude start and end)
	for i := 1; i < len(newRoutes)-1; i++ {
		set[newRoutes[i]] = true
	}

	// check if intermediate nodes from route2 exist in the set

	for _, route := range definedRoutes {
		for i := 1; i < len(route)-1; i++ {
			if set[route[i]] {
				return true
			}
		}
	}
	return false
}
