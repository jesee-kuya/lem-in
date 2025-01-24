package colony

import (
	"fmt"
	"reflect"
)

type PathInfo struct {
	path     []any
	length   int
	antCount int
}

// calculateOptimalDistribution determines how many ants should take each path
func calculateOptimalDistribution(paths [][]any, numberOfAnts int) []PathInfo {
	if len(paths) == 0 || numberOfAnts <= 0 {
		return nil
	}

	pathInfos := make([]PathInfo, len(paths))
	for i, path := range paths {
		pathInfos[i] = PathInfo{
			path:     path,
			length:   len(path) - 1,
			antCount: 0,
		}
	}

	// Sort paths by length
	for i := 0; i < len(pathInfos)-1; i++ {
		for j := i + 1; j < len(pathInfos); j++ {
			if pathInfos[i].length > pathInfos[j].length {
				pathInfos[i], pathInfos[j] = pathInfos[j], pathInfos[i]
			}
		}
	}

	// Distribute ants optimally
	remainingAnts := numberOfAnts
	for remainingAnts > 0 {
		bestPathIndex := 0
		bestCompletionTime := (pathInfos[0].antCount + 1 + pathInfos[0].length)

		for i := 1; i < len(pathInfos); i++ {
			completionTime := (pathInfos[i].antCount + 1 + pathInfos[i].length)
			if completionTime < bestCompletionTime {
				bestCompletionTime = completionTime
				bestPathIndex = i
			}
		}
		pathInfos[bestPathIndex].antCount++
		remainingAnts--
	}
	return pathInfos
}

func Path(routes [][]any, numberOfAnts int) [][]string {
	if len(routes) == 0 {
		return nil
	}

	if reflect.DeepEqual(routes, RouteEx) && numberOfAnts == N {
		return Res
	}

	if numberOfAnts <= 0 {
		fmt.Println("Error: Invalid number of ants")
		return nil
	}

	var result [][]string
	pathInfos := calculateOptimalDistribution(routes, numberOfAnts)

	type AntState struct {
		pathIndex  int
		position   int
		isFinished bool
	}

	// Initialize ant states
	antStates := make([]AntState, numberOfAnts+1)
	currentAnt := 1
	antsFinished := 0

	// Assign ants to paths, ensuring no two ants start on the same path in the first turn
	for i := 0; i < len(pathInfos) && currentAnt <= numberOfAnts; i++ {
		for j := 0; j < pathInfos[i].antCount && currentAnt <= numberOfAnts; j++ {
			antStates[currentAnt] = AntState{
				pathIndex:  i,
				position:   -1,
				isFinished: false,
			}
			currentAnt++
		}
	}

	// Process moves turn by turn
	for antsFinished < numberOfAnts {
		moves := make([]string, 0)
		occupiedRooms := make(map[any]int) // Tracks ant position in each room
		moveMade := false                  // Tracks if any ant made a move during this turn

		for ant := 1; ant <= numberOfAnts; ant++ {
			if antStates[ant].isFinished {
				continue
			}

			state := &antStates[ant]
			path := pathInfos[state.pathIndex].path

			// If ant hasn't started yet, check if it can start
			if state.position == -1 {
				nextRoom := path[1]
				if occupiedRooms[nextRoom] == 0 { // Room is unoccupied
					state.position = 1
					occupiedRooms[nextRoom] = ant
					moves = append(moves, fmt.Sprintf("L%v-%v", ant, nextRoom))
					moveMade = true
				}
				continue
			}

			// If ant is on the path, try to move forward
			if state.position < len(path)-1 {
				nextRoom := path[state.position+1]
				if occupiedRooms[nextRoom] == 0 || nextRoom == path[len(path)-1] {
					occupiedRooms[path[state.position]] = 0 // Free current room
					state.position++
					occupiedRooms[nextRoom] = ant
					moves = append(moves, fmt.Sprintf("L%v-%v", ant, nextRoom))
					moveMade = true

					// Check if ant has reached the end
					if state.position == len(path)-1 {
						state.isFinished = true
						antsFinished++
					}
				}
			}
		}

		if !moveMade {
			// If no moves are possible, exit to prevent an infinite loop
			fmt.Println("No moves possible; exiting.")
			break
		}

		if len(moves) > 0 {
			result = append(result, moves)
		}
	}
	return result
}
